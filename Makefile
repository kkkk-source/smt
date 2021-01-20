build:
	go build -o bin/smt main.go
	gcc -o bin/time time.c -Wall -Werror

run:
	go run main.go

fmt: 
	gofmt -w -s .
