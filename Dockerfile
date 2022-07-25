FROM golang:alpine

COPY ./src/ /var/work/src
COPY ./go.mod /var/work/go.mod
WORKDIR /var/work

ENTRYPOINT ["/var/work/src/server/server"]