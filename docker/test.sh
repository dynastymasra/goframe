#!/bin/bash

#Create Test Results Folder
cd /go/src/github.com/dynastymasra/goframe
mkdir -p test-results

#Test Apis Bulk
cd /go/src/github.com/dynastymasra/goframe/app/controller
go test -v -cover -covermode set -coverprofile=/go/src/github.com/dynastymasra/goframe/test-results/app_controller_ping.cover.txt | tee /go/src/github.com/dynastymasra/goframe/test-results/app_controller_ping.cover.out

#Compile Test Results
cd /go/src/github.com/dynastymasra/goframe/test-results
cat *.cover.txt | grep -v mode: | sort -r | \
cat *.cover.out >> coverage.out
rm *.cover.txt
rm *.cover.out

#Collect artifacts
mkdir -p /artifacts/
cp /go/src/github.com/dynastymasra/goframe/test-results/*.out /artifacts/