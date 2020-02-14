package main

import (
	"fmt"

	"github.com/johnksft/gocurso/maps"
)

type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s \n", name, p.Name)
}

//Herencia de Structs
type PlatziCarrer struct {
	PlatziCourse
}

const helloWorld string = " Hola %s %s bienvenido \n"

func main() {
	mapa := maps.GetMap("Juan")
	fmt.Println(mapa)
	platziCourse := PlatziCourse{
		Name: "Go Lang", Slug: "/go-lang", Skills: []string{"Backend", "2"},
	}
	platziCourse1 := new(PlatziCourse)
	platziCourse1.Name = "Go1"
	platziCourse1.Slug = "/go"
	platziCourse1.Skills = []string{"backend"}
	fmt.Println(platziCourse)
	platziCourse.Subscribe("Johnk")
	fmt.Println(platziCourse1)
}
