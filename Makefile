BINARY := alcolator.bin
VERSION := 2023-02-02
SOURCES := $(wildcard cmd/*.go)
COMMIT_ID := $(shell git describe --tags --always)
BUILD_TIME := $(shell go run tool/rfc3339date.go)
UNAME := $(shell uname -s)
LDFLAGS = -ldflags "-X main.version=${VERSION} -X main.compiled=${BUILD_TIME} -X main.commitID=${COMMIT_ID} -s -w"

ifeq ($(UNAME), Linux)
	DFLAG := -d
endif

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	env CGO_ENABLED=0 go build ${LDFLAGS} -o $@ ${SOURCES}

# .PHONY: vfsgen
# vfsgen:
# 	cd srv/assets && go generate

# .PHONY: run
# run:
# 	go run -tags=dev . -d srv -l :9696

# .PHONY: install
# install:
# 	env CGO_ENABLED=0 go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ]; then rm -f ${BINARY}; fi

