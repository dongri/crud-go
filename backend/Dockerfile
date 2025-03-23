FROM golang:1.24 as builder
LABEL  maintainer "Dongri Jin <dongrium@gmail.com>"

COPY . /go/src/github.com/dongri/crud-go/backend
WORKDIR /go/src/github.com/dongri/crud-go/backend

RUN go mod download
RUN CGO_ENABLED=0 go build

FROM alpine:latest

WORKDIR /backend
COPY --from=builder /go/src/github.com/dongri/crud-go/backend/backend .

EXPOSE 3001
ENTRYPOINT ["./backend"]
