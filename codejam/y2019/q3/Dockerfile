FROM golang:1.7.4-alpine

WORKDIR /go/src

ADD . /go/src

RUN go build -o main main.go \
    \
    && echo "---\n" \
    && ./main < input.txt \
    && echo "\n---"
