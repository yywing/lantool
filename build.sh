#! /bin/bash
set -ex

# GITHASH=$(git rev-parse --short=8 HEAD 2>/dev/null || echo "__unknown__")
# GOGITHASH="-X 'main.gitHash=${GITHASH}'"

VERSION="-X 'main.version=$1'"
PRODUCT="lantool"


CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags community -ldflags "-w -s -extldflags '-static' ${VERSION}" -o release/${PRODUCT}_windows_amd64.exe main.go
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -tags community -ldflags "-w -s -extldflags '-static' ${VERSION}" -o release/${PRODUCT}_windows_386.exe main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags community -ldflags "-w -s -extldflags '-static' ${VERSION}" -o release/${PRODUCT}_linux_amd64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -tags community -ldflags "-w -s -extldflags '-static' ${VERSION}" -o release/${PRODUCT}_linux_386 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags community -ldflags "-w -s -extldflags '-static' ${VERSION}" -o release/${PRODUCT}_linux_arm64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -tags community -ldflags "-w -s -extldflags '-static' ${VERSION}" -o release/${PRODUCT}_linux_arm main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags community -ldflags "-w -s -extldflags '-static' ${VERSION}" -o release/${PRODUCT}_darwin_amd64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -tags community -ldflags "-w -s -extldflags '-static' ${VERSION}" -o release/${PRODUCT}_darwin_arm64 main.go

# 生成 sha256
find release/ ! -name 'sha256.txt' -type f -exec shasum -a 256 {} \; > release/sha256.txt