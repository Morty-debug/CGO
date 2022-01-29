package main

import "C"
import "fmt"

//export Print
func Print(argumento *C.char){
	fmt.Println(C.GoString(argumento))
}

func main(){
	fmt.Print("Desde Go ")
	cmain()
}
