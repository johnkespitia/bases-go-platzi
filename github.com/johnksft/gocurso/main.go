package main

import (
	"errors"
	"fmt"

	"github.com/johnksft/gocurso/structs"
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
	pointerTest()
	structs.InterfaceTest()
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
