#!/bin/bash

# This script is for when we need GO in WSL Ubuntu which means we can run generate.sh on Windows

goVersion="1.18.2"

curl -OL https://golang.org/dl/go$goVersion.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go$goVersion.linux-amd64.tar.gz
rm -f go$goVersion.linux-amd64.tar.gz

echo "export PATH=$PATH:/usr/local/go/bin" > ~/.bash_profile
export PATH=$PATH:/usr/local/go/bin
go version