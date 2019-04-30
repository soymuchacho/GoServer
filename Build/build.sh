#!/bin/bash

case "$1" in
	agent)
		./build_agent.sh
		;;
	test)
		./build_test.sh
		;;
	all)
		./build_agent.sh
		./build_test.sh
		;;
	*)
		echo "Usage: build.sh {agent|test|all}"
		exit 2
esac


