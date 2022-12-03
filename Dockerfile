FROM golang:1.10.5-alpine3.8

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]
