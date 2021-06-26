FROM golang:1.16-buster

RUN mkdir -p /app
WORKDIR /app

ADD . /app

ENV GO111MODULE=on
#development only
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

#production
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./api /app/cmd/api/main.go
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./dist/cron /app/cmd/cron/main.go
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./dist/migrations /app/cmd/migrations/main.go

#development only
RUN apt-get install gcc make

EXPOSE 8080

CMD ["/dist/api"]
