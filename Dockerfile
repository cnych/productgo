FROM golang:1.14.0-alpine3.11 AS build-env
ADD . /go/src/app
WORKDIR /go/src/app
RUN GOPROXY=https://goproxy.cn && \
    GO111MODULE=on && \
    CGO_ENABLED=0 && \
    go build -v -o /go/src/app/app-server

FROM alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add -U tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
WORKDIR /var/www/html
COPY --from=build-env /go/src/app/app-server /var/www/html/app-server
COPY --from=build-env /go/src/app/static /var/www/html/static
COPY --from=build-env /go/src/app/views /var/www/html/views
COPY --from=build-env /go/src/app/conf /var/www/html/conf
EXPOSE 8080
CMD [ "./app-server" ]
