package main

import (
	"log"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			log.Println(i)
		}()
	}

	time.Sleep(1 * time.Second)
}
