#!/bin/bash

rpcname="dbrpc redisrpc"

for fl in ${rpcname}
do
	echo "protoc --go_out=plugins=grpc:. "${fl}.proto
	protoc --go_out=plugins=grpc:. ${fl}.proto

	if [ ! -d ${fl} ];then
		echo "the path ${fl} not exsit, create it"
		mkdir ${fl}
	fi	
	touch ./${fl}/go.mod
	cp -rt ./${fl}  ./${fl}.pb.go 
done

#read -p "Enter To Exit..."
