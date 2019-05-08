#!/bin/bash

projectName="Proto"
testpath="../Share/"$projectName

cd $testpath

echo -e "\e[1;32mstart build $projectName : \e[0m"

./proto.sh

echo -en "\e[1;32mbuild result : \e[0m"
ret=$?
if [ $ret -eq 0 ];then
	echo -e "\e[1;32m 【ok】\e[0m"
else
	echo -e "\e[1;32m 【failed】\e[0m"
fi

