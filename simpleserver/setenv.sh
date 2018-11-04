#!/bin/bash

# Project: Simple Server Go version
# File: setenv.sh
# Description: Sets GO_HOME, and sets PATH.
# NOTE: Nothing.
# Copyright (c) 2018 Kari Marttila
# Author: Kari Marttila
# Version history:
# - 2018-11-02: First version.


export GOPATH=$(pwd)
echo "GOPATH="$GOPATH
export GOROOT=/mnt/local/go-1.11
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
echo "PATH="$PATH

