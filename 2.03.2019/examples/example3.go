package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var err error                                       // error 1 // HL

	defer func() {
		if err != nil {                                 // error 1
			log.Printf("FAIL, %v", err)
			return
		}

		log.Println("OK")
	}()

	if err := errors.New("some error"); err != nil {    // error 2 // HL
		err = fmt.Errorf("hmm..., %v", err)             // error 2
		log.Printf("ooops, %v", err)
		return
	}
}
