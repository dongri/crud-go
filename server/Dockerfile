FROM golang:1.19.3 as builder
LABEL  maintainer "Dongri Jin <dongrify@gmail.com>"

COPY . /go/src/github.com/dongri/crud-go/server
WORKDIR /go/src/github.com/dongri/crud-go/server

RUN go mod download
RUN CGO_ENABLED=0 go build

FROM alpine:latest

WORKDIR /server
COPY --from=builder /go/src/github.com/dongri/crud-go/server/server .

EXPOSE 3001
ENTRYPOINT ["./server"]
