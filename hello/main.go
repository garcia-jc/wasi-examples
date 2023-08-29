package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("duracion", time.Since(start))
	}()
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Println(err)
		return
	}
	for _, entry := range entries {
		log.Println(entry.Type(), entry.Name())
	}
}
