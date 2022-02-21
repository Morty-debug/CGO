#include <stdio.h>

#ifdef gcc
#include "Print.h"
int cmain(int argc, char *argv[]); //metodo cmain que esta fuera de la condicion
int main(int argc, char *argv[]){
	printf("Desde C ");
	return cmain(argc,argv);
}
#else
extern void Print(char* argumento); //metodo Print que esta en main.go
#endif

//funcion principal para el compilador Go o por el compilador C
int cmain(int argc, char *argv[]){
	int i;
	printf("Argumentos: %d\n",argc);
	for(i=0; i<argc; i++){
		Print(argv[i]);
	}
	return 0;	
}
