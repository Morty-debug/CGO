// file nogcc.go

// +build !gcc

package main

//extern int cmain(int argc, char *argv[]);
import "C"
import "os"

func cmain() {
	argc := len(os.Args)
	var argv = make([]*C.char, argc)
	for i:=0 ; i<argc; i++ {
		argv[i] =  C.CString(os.Args[i])
	}
	C.cmain(C.int(argc), &argv[0])
}

