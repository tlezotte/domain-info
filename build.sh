#!/bin/bash
archs=(linux darwin windows)

for arch in ${archs[@]}
do
    env GOOS=${arch} GOARCH=amd64 go build -o hello_${arch}
done
