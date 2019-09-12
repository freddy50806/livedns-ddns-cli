FROM golang:1.13
MAINTAINER GTB

ADD . /opt/livedns-ddns-cli/
WORKDIR /opt/livedns-ddns-cli/
RUN go mod tidy && go build
ENTRYPOINT [ "./livedns-ddns-cli" ] 