package main

import "fmt"

// begin interface OMIT
type Wrapper interface {
	Wrap(string)
}

// end interface OMIT

// begin custom error OMIT
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

// end custom error OMIT
