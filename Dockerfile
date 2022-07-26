FROM golang:1.18
MAINTAINER Nakashima
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /ashi
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY *.go ./
RUN go build -o ashi .
EXPOSE 10490
CMD ["./ashi"]