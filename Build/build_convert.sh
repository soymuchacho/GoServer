#!/bin/bash

projectName="DBConvert"
testpath="../"$projectName

cd $testpath

echo -en "\e[1;32mstart build $projectName : \e[0m"

go build -o $projectName main.go

ret=$?
if [ $ret -eq 0 ];then
	echo -e "\e[1;32m 【ok】\e[0m"
else
	echo -e "\e[1;32m 【failed】\e[0m"
fi

