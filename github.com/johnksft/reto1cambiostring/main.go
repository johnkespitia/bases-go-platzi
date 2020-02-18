package main

import (
	"fmt"
	"strings"
)

func main() {
	text := getTexto()
	textTransform:=replaceString(text)
	fmt.Println(textTransform)
}

func getTexto() string{
	var texto string 
	fmt.Println("Ingrese la palabra:")
	fmt.Scanf("%s",&texto)
	return texto
}

func replaceString(texto string) string{
	old := string([]rune(texto)[0])
	return strings.Replace(texto, old, strings.ToUpper(old), 1)
}