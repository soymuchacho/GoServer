#!/bin/bash

agentpath="../AgentServer"

cd $agentpath

echo -e "\e[1;32mbegin build agent server\e[0m"

go build -o AgentServer main.go
ret=$?
if [ $ret -eq 0 ];then
	echo -e "\e[1;32mbuild agent server successful\e[0m"
else
	echo -e "\e[1;32mbuild agent server failed\e[0m"
fi


