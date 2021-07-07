package main

func main() {
	go func() {
		for {}
	}()
	c := make(chan bool) // Unbuffred Channel
	// c := make(chan bool, 1) // Buffered Channel, buffer size of 1
	c <- true
	<-c
}
