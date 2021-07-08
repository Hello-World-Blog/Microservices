FROM golang:1.16.5 as builder

WORKDIR /go-microservice/

COPY . .

RUN CGO_ENABLED=0 go build -o microservice /go-microservice/main.go

FROM alpine:latest

WORKDIR /go-microservice

COPY --from=builder /go-microservice/ /go-microservice/

EXPOSE 9090

CMD ./microservice