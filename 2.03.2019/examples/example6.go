package main

import (
	"log"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) { // HL
			log.Println(i)
		}(i) // HL
	}

	time.Sleep(1 * time.Second)
}
