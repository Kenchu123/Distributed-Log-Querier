# first stage
FROM golang:1.20.7-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o bin/ds-grep cmd/ds-grep/main.go
RUN go build -o bin/ds-grep-server cmd/ds-grep-server/main.go

# second stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/bin/ds-grep .
COPY --from=builder /app/bin/ds-grep-server .
COPY --from=builder /app/.dsgrep/config.yml .dsgrep/config.yml
EXPOSE 7122
CMD ["/app/ds-grep-server"]
