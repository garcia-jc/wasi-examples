package main

import "fmt"

func main() {
	ch := make(chan int)
	ack := make(chan struct{})
	go func() {
		v := <-ch
		fmt.Println("valor recibido", v)
		ack <- struct{}{}
	}()
	fmt.Println("enviando valor")
	ch <- 7
	close(ch)
	<-ack
}
