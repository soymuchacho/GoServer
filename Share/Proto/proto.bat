
rpcprotos="mysqlrpc.proto redisrpc.proto"

for fl in ${rpcprotos}
do
	echo "protoc --go_out=plugins=grpc:. "$fl
	
	protoc --go_out=plugins=grpc:. $fl
done
