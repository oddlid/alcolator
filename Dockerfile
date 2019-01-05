FROM golang:stretch as builder
COPY . ${GOPATH}/src/github.com/oddlid/alcolator
WORKDIR ${GOPATH}/src/github.com/oddlid/alcolator/srv
RUN go get -d -v ./...
RUN make


FROM alpine:latest
LABEL maintainer="Odd E. Ebbesen <oddebb@gmail.com>"
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
ENTRYPOINT ["tini", "-g", "--", "alcolatorsrv"]
CMD ["-l", ":9600"]

