ARG BINARY=alcolator

FROM golang:latest AS builder
ARG BINARY
COPY . ${GOPATH}/src/github.com/oddlid/alcolator/
WORKDIR ${GOPATH}/src/github.com/oddlid/alcolator/
RUN make BINARY=${BINARY}


FROM alpine:latest
LABEL maintainer="Odd E. Ebbesen <oddebb@gmail.com>"
ARG BINARY
ARG BINPATH=/usr/local/bin/
RUN apk add --no-cache --update ca-certificates \
    && rm -rf /var/cache/apk/*
RUN adduser -D -u 1000 alcsrv
COPY --from=builder /go/src/github.com/oddlid/alcolator/${BINARY} ${BINPATH}${BINARY}
RUN chown alcsrv ${BINPATH}${BINARY} && chmod 555 ${BINPATH}${BINARY}
EXPOSE 9600
USER alcsrv
# ARG does not expand in the CMD instruction, so we need to copy the ARG to an ENV var
ENV ALC_BIN=${BINARY}
CMD ${ALC_BIN} serve -l :9600
