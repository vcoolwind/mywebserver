FROM alpine:3.11

RUN apk update  \
    && apk add --no-cache libc6-compat \
    && rm -rf /var/cache/apk/* \
    && mkdir -p /go/webserver

ENV SERVER_NAME webserver-default

WORKDIR /go/webserver
COPY  mywebserver /go/webserver/mywebserver
EXPOSE 12345
RUN chmod +x /go/webserver/mywebserver
CMD ["/go/webserver/mywebserver"]

