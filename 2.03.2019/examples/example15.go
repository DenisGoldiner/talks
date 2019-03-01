package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	//begin show OMIT
	var err error
	var a = 10

	defer func() {
		if err != nil {
			log.Printf("FAIL, %v", err)
			return
		}
		log.Println("OK")
	}()

	execute := func(b int) (err error) {
		c := b + a
		if c < 100 {
			err = errors.New("some error")
			return
		}
		a = c
		return
	}

	if err := execute(3); err != nil {
		err = fmt.Errorf("problem with a = %d", a)
	}
	//end show OMIT
}
