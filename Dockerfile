FROM golang:latest

RUN mkdir -p /go/src/product-search_go_solr

WORKDIR /go/src/product-search_go_solr

COPY . /go/src/product-search_go_solr/

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/product-search_go_solr