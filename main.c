#include <stdio.h>

#ifdef gcc
#include "Print.h"
int cmain(int argc, char *argv[]);
int main(int argc, char *argv[]){
	printf("Desde C ");
	return cmain(argc,argv);
}
#else
extern void Print(char* argumento);
#endif


int cmain(int argc, char *argv[]){
	int i;
	printf("Argumentos: %d\n",argc);
	for(i=0; i<argc; i++){
		Print(argv[i]);
	}
	return 0;	
}
