FROM golang:latest

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o server

EXPOSE 80

CMD ["/app/server"]
