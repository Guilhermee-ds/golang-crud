FROM golang:latest

workdir /app

COPY ..

RUN go get -d -v ./...

RUN go build -o go-crud .

EXPOSE 8000

CMD ["./go-crud"]


