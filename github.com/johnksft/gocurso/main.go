package main

import (
	"errors"
	"fmt"
	"time"
	//"github.com/johnksft/gocurso/structs"
)

const helloWorld string = " Hola %s %s bienvenido \n"

func main() {
	/*
		structs.InterfaceTest()
		res, err := sumaRegular(3, 3)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	*/
	//pointerTest()
	//structs.InterfaceTest()
	go forGo(300)
	//go forGo(20)
	time.Sleep(10000 * time.Millisecond)
}

func sumaRegular(num1 interface{}, num2 interface{}) (int, error) {
	switch num1.(type) {
	case string:
		return 0, errors.New("El primer valor es un String")
	}
	switch num2.(type) {
	case string:
		return 0, errors.New("El segundo valor es un String")
	}

	return num1.(int) + num2.(int), nil
}

// pointerTest
func pointerTest() {
	var a int = 10
	var b *int
	b = &a
	a = a * 2
	*b = *b * 2
	fmt.Println(a, *b)
	fmt.Println(&a, b)
}

func helloGo(index int) {
	fmt.Println("Hola Soy un print rn la Go Rutine #", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}
