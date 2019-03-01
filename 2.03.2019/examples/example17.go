package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var err error

	defer func() {
		if err != nil {
			log.Printf("FAIL, %v", err)
			return
		}

		log.Println("OK")
	}()

	if err := errors.New("some error"); err != nil {
		err = fmt.Errorf("hmm..., %v", err)
		log.Printf("ooops, %v", err)
		return
	}
}

type myHandler struct{}

//begin show OMIT
func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var err error                                        // err 1

    defer func() {
        if err != nil {                                  // err 1
            log.Printf("FAIL, %v", err)
            http.Error(w, err, http.StatusBadRequest)    // err 1
            return
        }
        log.Println("OK")
    }()

    if true {
        a, err := foo()                                   // err 2
        if err != nil {                                   // err 2
            return
        }
        ...
        b, err := bar()                                   // err 3
        if err != nil {                                   // err 3
            return
        }
    }
}

//end show OMIT
