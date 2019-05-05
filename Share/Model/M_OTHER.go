package dbmodel

import "time"

type User_auto_id struct {
	AutoID     int       `orm:"autoID" json:"autoID"`
	CreateTime time.Time `orm:"create_time" json:"create_time"`
	CreateIp   string    `orm:"create_ip" json:"create_ip"`
	Account    string    `orm:"account" json:"account"`
	Password   string    `orm:"password" json:"password"`
}

func (*User_auto_id) TableName() string {
	return "user_auto_id"
}
