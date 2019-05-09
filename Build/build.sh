#!/bin/bash

PROJECT_NAME="GoServer"

Archive(){
	dir=../${PROJECT_NAME}-${BUILD_VERSION}
	mkdir -p $dir
	mkdir -p $dir/AgentServer
	mkdir -p $dir/DBServer
	mkdir -p $dir/AgentServer/conf
	mkdir -p $dir/DBServer/conf

	cp -rt $dir ../goserver.service
	cp -rt $dir ../install.sh
	cp -rt $dir ../runserv.sh
	cp -rt $dir/AgentServer ../AgentServer/AgentServer
	cp -rt $dir/AgentServer ../AgentServer/conf
	cp -rt $dir/DBServer ../DBServer/DBServer
	cp -rt $dir/DBServer ../DBServer/conf
	tar -czf ${PROJECT_NAME}-${BUILD_VERSION}.tar $dir

	mv ${PROJECT_NAME}-${BUILD_VERSION}.tar ../
}


case "$1" in
	proto)
		./build_proto.sh
		;;
	agent)
		./build_agent.sh
		;;
	convert)
		./build_convert.sh
		;;
	db)
		./build_db.sh
		;;
	test)
		./build_test.sh
		;;
	all)
		./build_proto.sh
		./build_agent.sh
		./build_convert.sh
		./build_db.sh
		./build_test.sh
		Archive
		;;
	*)
		echo "Usage: build.sh {proto|agent|convert|db|test|all}"
		exit 2
esac


