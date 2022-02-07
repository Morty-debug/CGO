package main

import "C"
import "fmt"

//export Print
func Print(argumento *C.char){
	fmt.Println(C.GoString(argumento))
}

func main(){
	fmt.Print("Desde Go ")
	cmain() //metodo que esta en gcc.go y en nogcc.go 
		//se invoca segun sea el argumento del compilador, si usara C no lo usara
}
