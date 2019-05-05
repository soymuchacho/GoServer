#!/bin/bash

case "$1" in
	agent)
		./build_agent.sh
		;;
	convert)
		./build_convert.sh
		;;
	test)
		./build_test.sh
		;;
	all)
		./build_agent.sh
		./build_convert.sh
		./build_test.sh
		;;
	*)
		echo "Usage: build.sh {agent|convert|test|all}"
		exit 2
esac


