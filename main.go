package main

//extern int cmain(int argc, char *argv[]);
import "C"
import(
	"os"
	"fmt"
)

//export Print
func Print(argumento *C.char){
	fmt.Println(C.GoString(argumento))
}

func main() {
	argc := len(os.Args)
	var argv = make([]*C.char, argc)
	for i:=0 ; i<argc; i++ {
		argv[i] =  C.CString(os.Args[i])
	}
	C.cmain(C.int(argc), &argv[0])
}
