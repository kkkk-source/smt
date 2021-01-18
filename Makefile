build:
	go build -o bin/main main.go
	gcc -o bin/time time.c -Wall -Werror

run:
	go run main.go

fmt: 
	gofmt -w -s .
