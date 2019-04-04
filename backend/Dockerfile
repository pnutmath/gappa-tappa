FROM golang:1.12.1-alpine3.9
RUN mkdir -p /go/src/github.com/pnutdev/gappa-tappa/backend
RUN apk update && apk add curl && apk add git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ADD . /go/src/github.com/pnutdev/gappa-tappa/backend/
WORKDIR /go/src/github.com/pnutdev/gappa-tappa/backend
RUN dep ensure -vendor-only
RUN go build -o /main
CMD ["/main"]