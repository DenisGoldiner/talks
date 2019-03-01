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
	var err error

	defer func() {
		if err != nil {
			log.Printf("FAIL, %v", err)
			http.Error(w, err, http.StatusBadRequest)
			return
		}

		log.Println("OK")
	}()

	a, err := foo()
	if err != nil {
		return
	}
	...
	b, err := bar()
	if err != nil {
		return
	}
	...
}

//end show OMIT
