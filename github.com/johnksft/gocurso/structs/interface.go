package structs

import "fmt"

type Platzi interface {
	Subscribe(name string)
}

func deferTest() {
	fmt.Println("Ha terminado la ejecución")
}

func InterfaceTest() {
	defer deferTest()
	platziCourse1 := new(PlatziCourse)
	platziCourse1.Name = "Go1"
	platziCourse1.Slug = "/go"
	platziCourse1.Skills = []string{"backend"}

	platziCarrier1 := new(PlatziCarrer)
	platziCarrier1.Name = "GoCarrier"
	platziCarrier1.Slug = "/go"
	platziCarrier1.Skills = []string{"backend"}

	callSubscribe(platziCourse1)
	callSubscribe(platziCarrier1)

}

func callSubscribe(p Platzi) {
	p.Subscribe("Johnk")
}
