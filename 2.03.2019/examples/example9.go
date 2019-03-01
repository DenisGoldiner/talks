package main

import "fmt"

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

// end custom error OMIT
