TARGET := ./bin/blog

all:
	go build -o ${TARGET} ./main.go
