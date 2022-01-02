FROM golang:alpine
WORKDIR /go/src/app
COPY *.go .
RUN go mod init
RUN go get -d -v ./...

CMD ["go", "run", "subscriber.go"]