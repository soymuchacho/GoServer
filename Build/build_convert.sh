#!/bin/bash

projectName="DBConvert"
testpath="../"$projectName

cd $testpath

echo -e "\e[1;32mbegin build $projectName\e[0m"

go build -o $projectName main.go

ret=$?
if [ $ret -eq 0 ];then
	echo -e "\e[1;32mbuild $projectName successful\e[0m"
else
	echo -e "\e[1;32mbuild $projectName failed\e[0m"
fi

