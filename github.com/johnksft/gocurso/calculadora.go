package main

import "fmt"

func main2() {
	greeting()
	num1 := getNumber()
	menu()
	option := getOption()
	num2 := getNumber()
	if option == 1 {
		num1 = sum(num1, num2)
	} else if option == 2 {
		num1 = substract(num1, num2)
	}
	fmt.Printf("Operación %d \n", option)
	fmt.Printf("El valor actual es: %d \n", num1)
	fmt.Println(getArray())
	fmt.Println(getSlice())
}

func getNumber() int {
	var number int
	fmt.Print("Ingrese un número: ")
	fmt.Scanf("%d", &number)
	return number
}

func getOption() int {
	var option int
	fmt.Print("Seleccione la opción: ")
	fmt.Scanf("%d", &option)
	return option
}

func sum(val1 int, val2 int) int {
	return val1 + val2
}

func substract(val1 int, val2 int) int {
	return val1 - val2
}

func greeting() {
	fmt.Println("***************************************************************")
	fmt.Println("**    <(°.°)> CALCULATOR ACTION                              **")
	fmt.Println("**      | |   First app on GO Lang                           **")
	fmt.Println("**      |||   JohnK Espitia                                  **")
	fmt.Println("***************************************************************")
}

func menu() {
	fmt.Println("***************************************************************")
	fmt.Println("**    OPTIONS...                                             **")
	fmt.Println("**    1. Sum                                                 **")
	fmt.Println("**    2. subtract                                            **")
	fmt.Println("**    3. show result                                         **")
	fmt.Println("***************************************************************")
}

func getArray() [3]string {
	var array1 [3]string
	array1[0] = "arreglo-demo-1"
	array1[1] = "arreglo-demo-2"
	array1[2] = "arreglo-demo-3"
	return array1
}

func getSlice() []string {
	var array1 []string
	array1 = append(array1, "Hola", "mundo")
	array1 = append(array1, "Adios", "Gente")
	return array1
}
