FROM golang:1.14 as build

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

WORKDIR /go/cache
ADD ./App/go.mod .
ADD ./App/go.sum .
RUN go mod download

WORKDIR /go/release

ADD ./App/ .
RUN go build -o app coffee.go


FROM ubuntu:18.04 as prod
COPY config /config
VOLUME ["/UserData", "/Thumb"]
COPY --from=build /go/release/app /
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/
RUN ls && chmod +x /app
CMD ["/app"]
EXPOSE 30070
