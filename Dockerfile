FROM golang:1.9 as build

RUN go get -u github.com/onsi/ginkgo/ginkgo
RUN go get -u github.com/onsi/gomega/...

RUN mkdir -p /go/src/github.com/stevenceuppens/go-rest-query-filter
WORKDIR /go/src/github.com/stevenceuppens/go-rest-query-filter
COPY . .
RUN go get -d -v ./...

CMD ["ginkgo", "-r", "-v"]