#!/usr/bin/env bash

set -xeEou pipefail

if [ "$(basename "$PWD")" != std ]; then
  echo "Only run from the std repo"
fi

if [[ ! $(go version) =~ ([0-9]+\.[0-9]+(\.[0-9]+)?) ]]; then
  echo "Could not extract Go version from \"$(go version)\""
  exit 1
fi
GO_VERSION=${BASH_REMATCH[1]}
GO_MINOR_N=$(echo "$GO_VERSION" | cut -d\. -f2)
GIT_TAG="v0.2.$GO_VERSION"

if [[ $GO_MINOR_N -lt 17 ]]; then
  go mod tidy
else
  go mod tidy -compat=1.17
fi
if [[ -n $(git status --short) ]]; then
  echo "Dirty working directory after go mod tidy"
  exit 1
fi

RET_CODE=0
git rev-list "$GIT_TAG.." > /dev/null 2>&1 || RET_CODE=$?
if [[ "$RET_CODE" -eq "0" ]]; then
  echo "Tag $GIT_TAG already exists"
  exit 1
fi

find . -path ./moqueries -prune -type f -o -name \*.go -exec rm {} \;

./moqueries/moqueries package std \
  --destination-dir . \
  --exclude-pkg-path-regex syscall

go build ./...
go test ./...
