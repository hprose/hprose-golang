#!/bin/bash -e
# Requires installation of: `github.com/wadey/gocovmerge`

cd $GOPATH/src/github.com/hprose/hprose-golang
go test -v -covermode=atomic -coverprofile=full_cov.out -coverpkg=./... ./...
