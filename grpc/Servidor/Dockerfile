FROM golang:1.17 as build-env

WORKDIR /go/src/app
COPY *.go .

RUN go mod init
RUN go get -d -v ./...
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static
ENV MONGO_ADDRESS="mongodb://sopes1:managerl@34.125.118.239:27017"
ENV REDIS_ADDRESS="redis://34.125.118.239:6379"
COPY --from=build-env /go/bin/app /
CMD ["/app"]