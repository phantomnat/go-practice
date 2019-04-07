FROM golang:1.7.4-alpine as builder

WORKDIR /go/src

ADD . /go/src

RUN go build -o main main.go

FROM python:2.7-alpine
WORKDIR /app
COPY --from=builder /go/src/main .
ADD . .
RUN ls -al \
    && python interactive_runner.py python testing_tool.py 0 -- ./main \
    && python interactive_runner.py python testing_tool.py 1 -- ./main