package main

import (
	"fmt"
	"github.com/johnksft/reto2cuadrado/figura"
)

func main() {
	fmt.Println("***************************************************************")
	fmt.Println("**    <(°.°)> Area Calculator                                ")
	fmt.Println("**      | |   First app on GO Lang                           ")
	fmt.Println("**      |||   JohnK Espitia                                  ")
	fmt.Println("***************************************************************")
	cuadrado := new(figura.Cuadrado)
	crearCuadrado(cuadrado)
	cuadrado.CalcPerimeter()
	cuadrado.CalcArea()
	fmt.Println("***************************************************************")
	fmt.Println("**     (°.°)/ Area Calculator                                ")
	fmt.Println("**     /| |   Al área de la figura es: ",cuadrado.Area,"                    ")
	fmt.Println("**      |||   Al Perímetro de la figura es: ",cuadrado.Perimetro,"             ")
	fmt.Println("**            JohnK Espitia                                  ")
	fmt.Println("***************************************************************")
}

func crearCuadrado(cuadrado *figura.Cuadrado){ 
	fmt.Println("Ingrese el Alto")
	fmt.Scanf("%f", &cuadrado.Alto)
	fmt.Println("Ingrese el Largo")
	fmt.Scanf("%f", &cuadrado.Largo)
}