FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o ./main

EXPOSE 3000

ENTRYPOINT [ "./main" ]