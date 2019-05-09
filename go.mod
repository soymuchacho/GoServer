module github.com/soymuchacho/GoServer

go 1.12

require (
	config v0.0.0-00010101000000-000000000000
	db v0.0.0-00010101000000-000000000000
	dbrpc v0.0.0-00010101000000-000000000000
	github.com/Unknwon/goconfig v0.0.0-20190425194916-3dba17dd7b9e
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/gin-gonic/gin v1.4.0
	github.com/gohouse/converter v0.0.3
	github.com/golang/protobuf v1.3.1
	github.com/jinzhu/gorm v1.9.7
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/klauspost/reedsolomon v1.9.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20181023030647-4e92f724b73b // indirect
	github.com/tjfoc/gmsm v1.0.1 // indirect
	github.com/xtaci/kcp-go v5.2.8+incompatible // indirect
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
	google.golang.org/grpc v1.20.1
	network v0.0.0-00010101000000-000000000000
	public v0.0.0-00010101000000-000000000000
	srpc v0.0.0-00010101000000-000000000000
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.38.0
	component => ./Common/component
	config => ./Common/config
	db => ./Common/db
	dbrpc => ./Share/Proto/dbrpc
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190506204251-e1dfcc566284
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190509005604-ec0fef3eb65a
	golang.org/x/image => github.com/golang/image v0.0.0-20190507092727-e4e5bf290fec
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190415191353-3e0bab5405d6
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190508220229-2d0786266e9c
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190509014725-d996b19ee77c
	google.golang.org/api => github.com/googleapis/googleapis v0.0.0-20190508164559-2f6e293d9a00
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190508193815-b515fa19cec8
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
	handle => ./AgentServer/handle
	network => ./Common/network
	public => ./Common/public
	redisrpc => ./Share/Proto/redisrpc
	srpc => ./Common/srpc
)
