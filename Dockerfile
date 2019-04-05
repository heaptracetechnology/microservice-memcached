FROM golang

RUN go get github.com/bradfitz/gomemcache/memcache

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/microservice-memcached

ADD . /go/src/github.com/heaptracetechnology/microservice-memcached

RUN go install github.com/heaptracetechnology/microservice-memcached

ENTRYPOINT microservice-memcached

EXPOSE 8000