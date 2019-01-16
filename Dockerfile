FROM golang:1.11.1-alpine3.8
WORKDIR /go/src/github.com/minhkhiemm/example-go
ADD . /go/src/github.com/minhkhiemm/example-go
RUN cd cmd/server/ && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o server .

FROM alpine:3.8
RUN apk add --no-cache ca-certificates
COPY --from=0 /go/src/github.com/minhkhiemm/example-go/cmd/server/server  /server 
WORKDIR /
ENV PORT 9988
CMD ["/server"]
