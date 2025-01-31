FROM golang:1.23.5-alpine3.21

EXPOSE 8080

ADD . /app
WORKDIR /app

RUN apk add --no-cache git
RUN go build -o app ./cmd

CMD ["./app"]