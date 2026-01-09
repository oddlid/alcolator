#!/usr/bin/env bash

# This requires the following to have been run at some point before running this script:
# docker buildx create --name builder --driver docker-container --use

readonly _version=$(gawk '/^VERSION/{print $NF}' Makefile)
docker buildx build \
  -t "oddlid/alcolatorsrv:$_version" \
  --build-arg VERSION="$_version" \
  --build-arg BUILD_DATE="$(date --rfc-3339=ns)" \
  --build-arg VCS_REF="$(git rev-parse --short HEAD)" \
  --platform=linux/amd64,linux/arm64 \
  --push .
