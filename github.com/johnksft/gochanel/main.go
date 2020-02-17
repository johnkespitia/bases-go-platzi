package main

import "fmt"

func main() {
	fmt.Println("ASIGNACIÓN BÁSICA")
	ch := make(chan string)
	go func() {
		defer close(ch)
		ch <- "Hola Channel"
	}()
	fmt.Println(<-ch)
	fmt.Println("ASIGNACIÓN ITERATIVA")
	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3000
		ch1 <- 4
		ch1 <- 5
		ch1 <- 6
	}()
	for n := range ch1 {
		fmt.Printf("%d\n", n)
	}
	fmt.Println("ASIGNACIÓN LIMITADA")
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	ch2 <- 3
	fmt.Println(<-ch2)
}
