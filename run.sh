#!/bin/bash
wd="/var/www"
export GOPATH=$wd
export GO_ENV=production
export PORT=80

go build $wd/hello.go
$wd/hello.go &
