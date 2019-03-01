package main

import (
	"errors"
	"fmt"
	"log"
)

//begin show OMIT
func main() {
	var a = 10
	var b = 3

	if res, err := summ(a, b); err != nil {
		err = fmt.Errorf("%v, summ result = %d", err, res)
		log.Printf("FAIL, %v", err)
		return
	}

	log.Println("OK")
}

// in case summ of a and b is lower than 100 will return error
// if summ is greater than 100 - summ and nil error would be returned
func summ(a, b int) (int, error) {
	c := a + b
	if c < 100 {
		return 0, errors.New("some error")
	}

	return c, nil
}

//end show OMIT
