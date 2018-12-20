FROM alpine:latest
MAINTAINER Odd E. Ebbesen <oddebb@gmail.com>

RUN apk add --no-cache --update tini ca-certificates \
                && \
                rm -rf /var/cache/apk/*
RUN adduser -D -u 1000 alcsrv

COPY alcolatorsrv.bin /usr/local/bin/alcolatorsrv
RUN chown alcsrv /usr/local/bin/alcolatorsrv && chmod 555 /usr/local/bin/alcolatorsrv
EXPOSE 9600
ENV PATH="/usr/local/bin:${PATH}"
USER alcsrv
ENTRYPOINT ["tini", "-g", "--", "alcolatorsrv"]
CMD ["-l", ":9600"]

