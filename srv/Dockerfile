#FROM oddlid/arch-cli:latest as builder
#MAINTAINER Odd E. Ebbesen <oddebb@gmail.com>
#ENV GOPATH=/go
#ARG BINARY=alcolatorsrv
#
#RUN go get -d -u -v github.com/gorilla/mux
#RUN go get -d -u -v github.com/oddlid/alcolator
#RUN go get -d -u -v github.com/urfave/cli
#RUN go get -d -u -v github.com/Sirupsen/logrus
#RUN go get -d -u -v github.com/GeertJohan/go.rice
#RUN go get -u -v github.com/GeertJohan/go.rice/rice

FROM alpine:latest
RUN apk add --no-cache --update tini ca-certificates \
                && \
                rm -rf /var/cache/apk/*
RUN adduser -D -u 1000 alcsrv

ARG BINPATH=/usr/local/bin/

COPY ${BINARY}.bin ${BINPATH}${BINARY}
RUN chown alcsrv ${BINPATH}${BINARY} && chmod 555 ${BINPATH}${BINARY}
EXPOSE 9600
USER alcsrv
ENTRYPOINT ["tini", "-g", "--", "${BINARY}"]
CMD ["-l", ":9600"]

