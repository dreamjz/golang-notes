#
# Build
#
FROM golang:alpine AS build-env
# set golang proxy
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . /app
RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o /go-jwt-demo
#
# Deploy
#
FROM gcr.io/distroless/static
COPY --from=build-env /go-jwt-demo /
EXPOSE 9090
ENTRYPOINT ["/go-jwt-demo"]
