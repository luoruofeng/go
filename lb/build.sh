#!/bin/bash

rm -rf ../bin/*

#Linux
cd /Users/luoruofeng/Desktop/go/lb/config
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/config
chmod 774 ../bin/config

cd /Users/luoruofeng/Desktop/go/lb/proxy
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/proxy
chmod 774 ../bin/proxy

cd /Users/luoruofeng/Desktop/go/lb/rs
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/rs
chmod 774 ../bin/rs

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/rs2
chmod 774 ../bin/rs2

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/rs3
chmod 774 ../bin/rs3

#Mac
# cd /Users/luoruofeng/Desktop/go/lb/config
# go build -o ../bin/config
# chmod 774 ../bin/config

# cd /Users/luoruofeng/Desktop/go/lb/proxy
# go build -o ../bin/proxy
# chmod 774 ../bin/proxy

# cd /Users/luoruofeng/Desktop/go/lb/rs
# go build -o ../bin/rs
# chmod 774 ../bin/rs

# go build -o ../bin/rs2
# chmod 774 ../bin/rs2

# go build -o ../bin/rs3
# chmod 774 ../bin/rs3
