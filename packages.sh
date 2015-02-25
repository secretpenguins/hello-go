#!/bin/bash
yum install golang
export GOPATH=$(pwd)

go get github.com/bradfitz/gomemcache/memcache
go get github.com/martini-contrib/sessions
go get github.com/go-sql-driver/mysql
go get github.com/coopernurse/gorp
go get github.com/go-martini/martini
go get github.com/codegangsta/martini-contrib/render
go get golang.org/x/crypto/bcrypt
