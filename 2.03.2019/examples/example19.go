package main

func main() {

	//begin show OMIT
	switch verb {
	case 'v', 's', 'x', 'X', 'q':
		// Is it an error or Stringer?
		// The duplication in the bodies is necessary:
		// setting handled and deferring catchPanic
		// must happen before calling the method.
		switch v := p.arg.(type) {
		case error:
			handled = true
			defer p.catchPanic(p.arg, verb)
			p.fmtString(v.Error(), verb) // HL
			return

		case Stringer:
			handled = true
			defer p.catchPanic(p.arg, verb)
			p.fmtString(v.String(), verb) // HL
			return
		}
	}
	//end show OMIT
}

