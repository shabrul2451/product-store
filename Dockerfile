FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8085

CMD ["./main"]
