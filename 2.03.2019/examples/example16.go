package main

import (
	"errors"
	"fmt"
)

func main() {
	a, err := foo(1)
	fmt.Println(a, err)
}

func foo(i int) (bool, error) {
	switch i {
	case 0:
		return true, nil
	case 1:
		return true, nil
	default:
		return false, errors.New("nor found")
	}
}

func bar(i int) (flag bool, err error) {
	switch i {
	case 0:
		flag = true
		return
	case 1:
		flag = true
		return
	default:
		err = errors.New("nor found")
		return
	}
}
