FROM golang:1.15.3-alpine3.12
RUN apk add git
COPY . /go/src/github.com/hexcraft-biz/email-manager/db-builder
WORKDIR /go/src/github.com/hexcraft-biz/email-manager/db-builder
RUN go get github.com/hexcraft-biz/email-manager/db-builder
RUN go install ./
CMD /go/bin/db-builder
