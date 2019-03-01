package main

import (
	"errors"
	"log"
)

func main() {
	f, err := Foo()
	if err != nil {
		log.Printf("FAIL, %v", err)
		return
	}

	log.Printf("OK, %t", f)
}

//begin show OMIT

// Foo  - API function wrapping
// flag - some indicator
// err  - some error
func Foo() (flag bool, err error) {
	return bar()
}

func bar() (bool, error) {
	var i int

	if i == 0 {
		return false, errors.New("i should not be 0")
	}

	return true, nil
}

//end show OMIT
