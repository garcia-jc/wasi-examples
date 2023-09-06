package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)
	ack := make(chan struct{})
	go func() {
		v := <-ch
		fmt.Println("valor recibido", v)
		ack <- struct{}{}
	}()
	fmt.Println("enviando valor aleatorio")
	ch <- 1 + rand.Intn(10)
	close(ch)
	<-ack
}
