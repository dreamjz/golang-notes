FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

WORKDIR $GOPATH/src/go-gin-example
COPY . $GOPATH/src/go-gin-example
RUN go build .

EXPOSE 9090
ENTRYPOINT ["./go-gin-example"]