#!/bin/bash
docker run --rm -v $(PWD):/go golang
GOOS=darwin GOARCH=amd64 go build -o cmemo
