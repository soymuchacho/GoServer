package session

type Session struct {
	Ip           string // session Ip address
	Port         int32  // session Port
	ConnTime     int64  // session connect time
	PackTime     int64  // session recv pack time
	LastPackTime int64  // session last recv pack time
}
