package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const PATH_FILE = "./files/hello.txt"

func main() {
	readFile()
}

func readFile(){
	dat, err := ioutil.ReadFile(PATH_FILE)
	check(err)
	text:=string(dat)
	var content []string
	content = strings.Split(text,"\n")
	for i:=0;i < len(content); i++{
		fmt.Printf(" Linea %d contiene: %s \n",i,content[i])
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}