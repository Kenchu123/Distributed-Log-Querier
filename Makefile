build: build-client build-server

build-client:
	go build -o bin/client cmd/client/main.go

build-server:
	go build -o bin/server cmd/server/main.go

run-client:
	go run cmd/client/main.go

run-server:
	go run cmd/server/main.go

clean:
	rm -rf bin
