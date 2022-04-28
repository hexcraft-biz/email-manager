FROM golang:1.15.3-alpine3.12 as golang-builder
RUN apk add git
COPY . /go/src/github.com/hexcraft-biz/email-manager
WORKDIR /go/src/github.com/hexcraft-biz/email-manager
RUN go get github.com/hexcraft-biz/email-manager
RUN go install ./

FROM alpine:3.12
COPY --from=golang-builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=golang-builder /go/bin/email-manager /var/www/app/
WORKDIR /var/www/app
EXPOSE 80
ENTRYPOINT /var/www/app/email-manager
