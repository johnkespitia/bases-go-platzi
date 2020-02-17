package main

import (
	"io"
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hola Go!")
}

func main(){
	http.HandleFunc("/", handler)
	fmt.Println("Runing Server on port 8000")
	http.ListenAndServe(":8000",nil)
	
}