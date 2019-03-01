package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

//begin main OMIT
func main() {
	if err := foo(); err != nil {
		if err, ok := err.(StackError); ok {
			log.Println(err.WithStack())
			return
		}

		log.Println(err.Error())
	}
}

func foo() error {
	return &myErr{
		code:  400,
		msg:   "Bad Request",
		stack: debug.Stack(),
	}
}

//end main OMIT

//begin err OMIT
type myErr struct {
	code  int
	msg   string
	stack []byte
}

type StackError interface {
	WithStack() string
}

func (me *myErr) WithStack() string {
	return fmt.Sprintf("%s, stack:\n%s", me.Error(), string(me.stack))
}

//end err OMIT

func (me *myErr) Error() string {
	switch me.code {
	case 400:
		return fmt.Sprintf("error: Not Found; \ncause: %s", me.msg)
	default:
		return me.msg
	}
}
