FROM oddlid/arch-cli:latest as builder
MAINTAINER Odd E. Ebbesen <oddebb@gmail.com>
ENV GOPATH=/go

RUN go get -d -u -v github.com/gorilla/mux
RUN go get -d -u -v github.com/shurcooL/httpfs/filter
RUN go get -d -u -v github.com/shurcooL/httpfs/html/vfstemplate
RUN go get -d -u -v github.com/shurcooL/httpfs/union
RUN go get -d -u -v github.com/Sirupsen/logrus
RUN go get -d -u -v github.com/urfave/cli

COPY . ${GOPATH}/src/github.com/oddlid/alcolator
WORKDIR ${GOPATH}/src/github.com/oddlid/alcolator/srv
RUN make


FROM alpine:latest
RUN apk add --no-cache --update tini ca-certificates \
                && \
                rm -rf /var/cache/apk/*
RUN adduser -D -u 1000 alcsrv

ARG BINARY=alcolatorsrv
ARG BINPATH=/usr/local/bin/

COPY --from=builder /go/src/github.com/oddlid/alcolator/srv/${BINARY}.bin ${BINPATH}${BINARY}
RUN chown alcsrv ${BINPATH}${BINARY} && chmod 555 ${BINPATH}${BINARY}
EXPOSE 9600
USER alcsrv
ENTRYPOINT ["tini", "-g", "--", "${BINARY}"]
CMD ["-l", ":9600"]

