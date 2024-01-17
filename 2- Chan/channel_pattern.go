package main

import (
	"fmt"
	"time"
)

func boring(str string, c chan string) {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("%s %d", str, i)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	c := make(chan string)
	go boring("Test", c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Ended")
}
