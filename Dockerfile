FROM dhi.io/golang:1-debian13-dev AS builder
WORKDIR /build
COPY . .
RUN make

FROM dhi.io/alpine-base:3.23
LABEL maintainer="Odd E. Ebbesen <oddebb@gmail.com>"
COPY --from=builder --chmod=555 /build/alcolator /usr/local/bin/alcolator
EXPOSE 9600
CMD ["/usr/local/bin/alcolator", "serve", "-l", ":9600"]
