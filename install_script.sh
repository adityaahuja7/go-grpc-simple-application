#!/bin/bash
sudo apt-get update
sudo apt install -y tar 
sudo apt install -y git
sudo apt install -y curl
curl -O https://dl.google.com/go/go1.22.0.linux-amd64.tar.gz
tar -xvf go1.22.0.linux-amd64.tar.gz
sudo mv go1.22.0.linux-amd64 /usr/local
tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
gsutil cp -r gs://dscd-files .
echo "DONE!"


