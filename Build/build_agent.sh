#!/bin/bash

agentpath="../AgentServer"

cd $agentpath

go build -o AgentServer main.go

echo "build agent server end"


