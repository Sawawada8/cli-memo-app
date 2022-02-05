#!/bin/bash
docker run -v $(PWD):/go golang
GOOS=darwin GOARCH=amd64 go build -o cmemo
