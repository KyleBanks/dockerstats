#!/bin/bash
set -e

go get -u github.com/golang/lint/golint
go get -u github.com/robertkrimen/godocdown/godocdown

echo "-------------  TEST  -------------"
go test -cover .

echo "-------------  VET  -------------"
go vet ./...

echo "-------------  LINT  -------------"
golint ./...

echo "-------------  FMT  -------------"
go fmt ./...

echo "-------------  DOCS  -------------"
godocdown -template=".godocdown.md" . > README.md