module github.com/soymuchacho/GoServer

go 1.12

require (
	config v0.0.0-00010101000000-000000000000 // indirect
	db v0.0.0-00010101000000-000000000000 // indirect
	dbrpc v0.0.0-00010101000000-000000000000 // indirect
	github.com/Unknwon/goconfig v0.0.0-20190425194916-3dba17dd7b9e // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/gin-gonic/gin v1.4.0 // indirect
	github.com/gohouse/converter v0.0.3 // indirect
	github.com/jinzhu/gorm v1.9.7 // indirect
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/klauspost/reedsolomon v1.9.1 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20181023030647-4e92f724b73b // indirect
	github.com/tjfoc/gmsm v1.0.1 // indirect
	github.com/xtaci/kcp-go v5.2.8+incompatible // indirect
	google.golang.org/grpc v1.20.1 // indirect
	network v0.0.0-00010101000000-000000000000 // indirect
	public v0.0.0-00010101000000-000000000000 // indirect
	srpc v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	component => ./Common/component
	config => ./Common/config
	db => ./Common/db
	dbrpc => ./Share/Proto/dbrpc
	handle => ./AgentServer/handle
	network => ./Common/network
	public => ./Common/public
	redisrpc => ./Share/Proto/redisrpc
	srpc => ./Common/srpc
)
