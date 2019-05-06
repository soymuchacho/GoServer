#!/bin/bash

rpcname="dbrpc redisrpc"

for fl in ${rpcname}
do
	echo "protoc --go_out=plugins=grpc:. "${fl}.proto
	protoc --go_out=plugins=grpc:. ${fl}.proto
	
	mkdir ${fl} 
	cp -rt ./${fl}  ./${fl}.pb.go 
done

read -p "Enter To Exit..."
