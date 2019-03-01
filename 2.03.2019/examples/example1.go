package main

import "log"

func main() {
	var v = 10                               // v 1

	if true { // HL
		v := 20                              // v 2
		log.Printf("inside if v = %d", v)    // v 2
	} // HL

	log.Printf("outside if v = %d", v)       // v 1
}
