BINARY := alcolator.bin
VERSION := 2021-08-20
SOURCES := $(wildcard *.go)
COMMIT_ID := $(shell git describe --tags --always)
BUILD_TIME := $(shell date +%FT%T%:z)
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.BUILD_DATE=${BUILD_TIME} -X main.COMMIT_ID=${COMMIT_ID} -d -s -w"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	env CGO_ENABLED=0 go build ${LDFLAGS} -o $@ ${SOURCES}

.PHONY: vfsgen
vfsgen:
	cd srv/assets && go generate

.PHONY: run
run:
	go run -tags=dev . -d srv -l :9696

.PHONY: install
install:
	env CGO_ENABLED=0 go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ]; then rm -f ${BINARY}; fi

