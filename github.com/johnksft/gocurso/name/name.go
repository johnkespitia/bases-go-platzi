package name

import "fmt"

// GetName obtiene y retorna el nombre de la persona
func GetName() string {
	var name string = "nombre test"
	fmt.Print("Ingresa tu nombre:")
	fmt.Scanf("%s", &name)
	return name
}
