#!/bin/bash

InstallPath="/opt/goserver/goserver-build-centos0"
servers="Agent"

start() {
	for serv in $servers;
	do
		echo $"starting goserver $serv :"
		cd ${InstallPath}/${serv}Server
		./start.sh
		ret=$?
		if [ $ret -eq 0 ];then
			echo -e "\e[1;32mOK\e[0m"
		else
			echo -e "\e[1;32mFailed\e[0m"
			retval=$ret
		fi
	done
	return $retval
}

stop() {
	for serv in $servers;
	do
		echo $"stopping goserver $serv :"
		kill -9 $(ps -ax | grep "$serv" | grep "Server" | awk '{print $1}') > /dev/null
		ret=$?
		if [ $ret -eq 0 ];then
			echo -e "\e[1;32mOK\e[0m"
		else
			echo -e "\e[1;32mFailed\e[0m"
			retval=$ret
		fi
	done
	return $retval
}

restart() {
	#stop
	#start
	echo "restart"
}

case "$1" in
	start)
		$1
		;;
	stop)
		$1
		;;
	restart)
		$1
		;;
	*)
		echo $"Usage: $0 {start|stop|restart}"
		exit 2
esac


