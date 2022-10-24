#!/bin/sh

echo ">>Building go binary"
go build main.go
echo ">>Success Build Binary"

echo ">>Build image"
docker build -t test/test .
echo ">>Success Build Image"

# fetch athentication from 3rd party
curl https://blablabla
# doing other process