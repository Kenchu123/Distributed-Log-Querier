# first stage
FROM golang:1.20.7-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o bin/ds-grep cmd/ds-grep/main.go
RUN go build -o bin/ds-grep-server cmd/ds-grep-server/main.go
EXPOSE 7122
CMD ["bin/ds-grep-server"]
