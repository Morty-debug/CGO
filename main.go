package main

import "C"
import "fmt"

//export Print
func Print(argumento *C.char){
	fmt.Println(C.GoString(argumento))
}

func main(){
	fmt.Print("Desde Go ")
	cmain() //metodo qu esta en gcc.go y en nogcc.go invoca dependiendo del compilado si usara C no lo usara
}
