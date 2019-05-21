#!/usr/bin/env make

.PHONY: all test

all: vendor Gopkg.lock cf/client.go interface.go foreach

test:
	ginkgo
	ginkgo cf
	ginkgo foreach

vendor:
	dep ensure

Gopkg.lock:
	dep ensure -update

cf/client.go:
	cd cf && go generate

interface.go:
	go generate definition.go

foreach:
	cd foreach && go generate
