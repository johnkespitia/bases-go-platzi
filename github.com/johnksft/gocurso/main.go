package main

import (
	"fmt"

	"github.com/johnksft/gocurso/maps"
)

const helloWorld string = " Hola %s %s bienvenido \n"

func main() {
	mapa := maps.GetMap("Juan")
	fmt.Println(mapa)
}
