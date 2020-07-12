#!/usr/bin/env bash

mkdir -p $GOPATH/src/GrantZheng/kit
git clone https://github.com/GrantZheng/kit.git $GOPATH/src/GrantZheng/kit
cd $GOPATH/src/GrantZheng/kit
go install
