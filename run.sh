#!/bin/bash
export GOPATH=$(pwd)

go build hello.go
./hello.go &
