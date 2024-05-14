FROM golang:1.15.7-alpine3.13

workdir /app

COPY ..

RUN go get -d -v ./...

RUN go build -o go-crud .

EXPOSE 8000

CMD ["./go-crud"]


