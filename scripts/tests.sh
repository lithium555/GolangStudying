#!/bin/bash

export GO111MODULE=on
go fmt ./...
go test -v -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out ./...