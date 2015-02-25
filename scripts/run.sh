#!/bin/bash
wd="/var/www"
export GOPATH=$wd
export GO_ENV=production
export PORT=80

go build -o $wd/hello $wd/hello.go
(cd $wd; exec $wd/hello >> log.txt 2>&1 &)