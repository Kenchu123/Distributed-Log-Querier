build: build-client build-server

build-client:
	go build -o bin/ds-grep cmd/ds-grep/main.go

build-server:
	go build -o bin/ds-grep-server cmd/ds-grep-server/main.go

run-client:
	go run cmd/ds-grep/main.go $(ARGS)

run-server:
	go run cmd/ds-grep-server/main.go $(ARGS)

clean:
	rm -rf bin
