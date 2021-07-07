package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	alice := 10000
	bob := 10000
	total := alice + bob

	var mu sync.Mutex

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			alice -= 1
			bob += 1
			mu.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			bob -= 1
			alice += 1
			mu.Unlock()
		}
	}()


	start := time.Now()
	for time.Since(start) <= 1*time.Second {
		mu.Lock()
		if alice +bob != total {
			fmt.Println("violation, alice=%v, bob=%v, total=%v", alice, bob, alice+bob)
		}
		mu.Unlock()
	}


}
