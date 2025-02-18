# build
FROM golang:1.23

WORKDIR /go-app-dir

COPY . .

RUN apt-get update && apt-get install -y bash

RUN go mod tidy

CMD ["go", "run", "main.go"]
# CMD ["sh", "run.sh"]