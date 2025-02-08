FROM golang:1.23-alpine

WORKDIR /go-app-dir

COPY . .

RUN go mod tidy
RUN go test -v ./...

CMD ["go", "run", "main.go"]