FROM golang:1.17-alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git
RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service *.go

FROM alpine:3.14
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/consignment-service .
CMD ['./consignment-service']