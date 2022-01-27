#include <stdio.h>
extern void Print(char* argumento);

int cmain(int argc, char *argv[]){
	int i;
	printf("Argumentos: %d\n",argc);
	for(i=0; i<argc; i++){
		Print(argv[i]);
	}	
}
