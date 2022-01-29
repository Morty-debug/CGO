
### compilar codigo en Go
```batch
go build -o main.exe
```


### compilar codigo en C
```batch
go build -o Print.dll -tags gcc -buildmode=c-shared
gcc -o main.exe main.c -D gcc ./Print.dll
```
