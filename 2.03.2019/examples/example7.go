package main

import (
	"log"
	"time"
)

//begin show OMIT
func main() {
	for i := 0; i < 10; i++ {
		go execute(i) // HL
	}

	time.Sleep(1 * time.Second)
}

// Prints an int value
func execute(i int) {
	log.Println(i)
}

//end show OMIT
