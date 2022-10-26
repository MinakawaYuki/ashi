FROM golang:1.18
MAINTAINER Nakashima
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /ashi
ADD . /ashi
RUN go build -o ashi .
EXPOSE 10490
CMD ["./ashi"]
