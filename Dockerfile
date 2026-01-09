FROM dhi.io/golang:1-debian13-dev AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN \
  --mount=type=cache,target=/go/pkg/mod \
  go mod download

COPY . .
RUN \
  --mount=type=cache,target=/root/.cache/go-build \
  make

FROM scratch
COPY --from=builder --chmod=555 /build/alcolator /alcolator
ENTRYPOINT [ "/alcolator" ]
# CMD ["/alcolator", "serve", "-l", ":9600"]
CMD ["--help"]

ARG AUTHORS="Odd E. Ebbesen <oddebb@gmail.com>"
ARG VERSION="unknown"
ARG BUILD_DATE="unknown"
ARG VCS_REF="unknown"
LABEL \
  org.opencontainers.image.version="${VERSION}" \
  org.opencontainers.image.created="${BUILD_DATE}" \
  org.opencontainers.image.revision="${VCS_REF}" \
  org.opencontainers.image.authors="${AUTHORS}}"

