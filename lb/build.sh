#!/bin/bash

cd /Users/luoruofeng/Desktop/go/lb/proxy
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/proxy_mod

cd /Users/luoruofeng/Desktop/go/lb/rs
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/rs_mod
