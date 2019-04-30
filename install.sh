#!/bin/bash

BasePath="/opt/goserver"
Version="goserver0.0.1-build1"
lnfile="goserver-build-centos0"

rm -rf ${BasePath} 

mkdir ${BasePath} 
mkdir ${BasePath}/${Version}
mkdir ${BasePath}/${Version}/TestServer
mkdir ${BasePath}/${Version}/AgentServer

cp -rt ${BasePath}/${Version} goserver.service 
cp -rt ${BasePath}/${Version} runserv.sh

cp -rt ${BasePath}/${Version}/AgentServer ./AgentServer/conf
cp -rt ${BasePath}/${Version}/AgentServer ./AgentServer/start.sh
cp -rt ${BasePath}/${Version}/AgentServer ./AgentServer/AgentServer

cp -rt ${BasePath}/${Version}/TestServer ./TestServer/conf
cp -rt ${BasePath}/${Version}/TestServer ./TestServer/start.sh
cp -rt ${BasePath}/${Version}/TestServer ./TestServer/TestServer


cp -rt /usr/lib/systemd/system/ goserver.service 

echo "ln -s ${BasePath}/${Version} ${BasePath}/${lnfile}"
ln -s ${BasePath}/${Version} ${BasePath}/${lnfile}

systemctl stop goserver
systemctl disable goserver.service
systemctl enable goserver.service
systemctl start goserver


