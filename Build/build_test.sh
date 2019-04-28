#!/bin/bash

projectName="TestServer"
testpath="../"$projectName

cd $testpath

go build -o $projectName main.go

echo "build "$projectName" end"


