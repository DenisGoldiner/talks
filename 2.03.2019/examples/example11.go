package main

import (
	"fmt"
	"log"
)

//begin show OMIT
func main() {
	if err := foo(); err != nil {
		err := fmt.Errorf("failed to execute foo, %v", err)
		log.Println(err)
	}
}

func foo() error {
	return &myErr{
		code: 400,
		msg:  "Bad Request",
	}
}

//end show OMIT

type myErr struct {
	code int
	msg  string
}

func (me *myErr) Error() string {
	switch me.code {
	case 400:
		return fmt.Sprintf("error: Not Found; \ncause: %s", me.msg)
	default:
		return me.msg
	}
}
