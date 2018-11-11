#!/bin/bash

#go build -a -v -race -o ./bin/simpleserver github.com/karimarttila/go/simpleserver/app/...

go build -o output/simpleserver github.com/karimarttila/go/simpleserver/app/main

