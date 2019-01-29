CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/aaronland/go-aws-tools
	cp -r auth src/github.com/aaronland/go-aws-tools/
	cp -r config src/github.com/aaronland/go-aws-tools/
	cp -r utils src/github.com/aaronland/go-aws-tools/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/go-ini/ini"
	@GOPATH=$(GOPATH) go get -u "github.com/jmespath/go-jmespath"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/iso8601duration"
	@GOPATH=$(GOPATH) go get -u "github.com/aws/aws-sdk-go-v2"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt auth/*.go
	go fmt cmd/*.go
	go fmt config/*.go
	go fmt utils/*.go


bin: 	self
	rm -rf bin/*
	@GOPATH=$(GOPATH) go build -o bin/aws-cloudfront-invalidate cmd/aws-cloudfront-invalidate.go
	@GOPATH=$(GOPATH) go build -o bin/aws-mfa-session cmd/aws-mfa-session.go
	@GOPATH=$(GOPATH) go build -o bin/aws-set-env cmd/aws-set-env.go
	@GOPATH=$(GOPATH) go build -o bin/aws-get-credential cmd/aws-get-credential.go
