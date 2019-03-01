package main

import (
	"fmt"
	"log"
)

type Wrapper interface {
	Wrap(string)
}

//begin show OMIT
func main() {
	if err := foo(); err != nil {
		wrap(err, "failed to execute foo")
		log.Println(err)
	}
}

func foo() error {
	return &myErr{
		code: 400,
		msg:  "Bad Request",
	}
}

func wrap(err error, msg string) {
	if err, ok := err.(Wrapper); ok {
		err.Wrap(msg)
		return
	}

	err = fmt.Errorf("%v, %s", err, msg)
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

func (me *myErr) Wrap(msg string) {
	me.msg = fmt.Sprintf("%s; %s", me.msg, msg)
}
