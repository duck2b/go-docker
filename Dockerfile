FROM golang:latest

WORKDIR	/app

COPY . .

EXPOSE  8090

CMD ["/app/main"]