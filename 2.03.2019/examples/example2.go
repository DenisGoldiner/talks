package main

import (
	"errors"
	"log"
)

//begin shadowing OMIT
func main() {
	if err := foo(); err != nil {
		log.Printf("FAIL, %v", err)
		return
	}

	log.Print("OK")
}

func foo() (err error) {                                // err 1 // HL
	if err := errors.New("some error"); err != nil {    // err 2
		log.Printf("ooops, %v", err)                    // err 2
	}

	return                                              // err 1 // HL
}

//end shadowing OMIT

func bar() (err error) {
	if confusionErr := errors.New("some error"); err != nil {
		log.Printf("ooops, %v", err)

		return confusionErr
	}

	return
}
