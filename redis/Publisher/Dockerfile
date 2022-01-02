FROM golang:alpine as build
WORKDIR /go/src/app
ADD *.go /go/src/app
RUN go mod init
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM  gcr.io/distroless/static
COPY --from=build /go/bin/app /
CMD ["/app"]