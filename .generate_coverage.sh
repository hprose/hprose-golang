#!/bin/bash -e
# Requires installation of: `github.com/wadey/gocovmerge`

cd $GOPATH/src/github.com/hprose/hprose-golang

rm -rf ./cov
mkdir cov

i=0
for dir in $(find . -maxdepth 10 -not -path './examples/*' -not -path './.git*' -not -path '*/_test.go' -type d);
do
    if ls ${dir}/*.go &> /dev/null; then
        go test -v -covermode=atomic -coverprofile=./cov/$i.out ./${dir}
        i=$((i+1))
    fi
done

gocovmerge ./cov/*.out > full_cov.out
rm -rf ./cov
