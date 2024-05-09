
all: autocannon testserver

.PHONY: autocannon
autocannon:
	go build -o bin/autocannon cmd/autocannon/main.go

.PHONY: testserver
testserver:
	go build -o bin/testserver cmd/testserver/main.go


