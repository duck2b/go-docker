FROM golang:latest

WORKDIR	/app

COPY . .

EXPOSE  8081

CMD ["/app/main"]