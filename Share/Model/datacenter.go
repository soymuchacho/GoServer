package dbmodel

type Amc_info struct {
	Time    int    `orm:"time" json:"time"`
	Amcinfo string `orm:"amcinfo" json:"amcinfo"`
}

func (*Amc_info) TableName() string {
	return "amc_info"
}

type Craft_hv struct {
	Craftid int     `orm:"craftid" json:"craftid"`
	Time    int     `orm:"time" json:"time"`
	Hvalue  float64 `orm:"hvalue" json:"hvalue"`
	Tvalue  float64 `orm:"tvalue" json:"tvalue"`
}

func (*Craft_hv) TableName() string {
	return "craft_hv"
}

type Craft_particle struct {
	Craftid      int    `orm:"craftid" json:"craftid"`
	Time         int    `orm:"time" json:"time"`
	ParticleInfo string `orm:"particle_info" json:"particle_info"`
}

func (*Craft_particle) TableName() string {
	return "craft_particle"
}

type Emission_Info struct {
	Time   int    `orm:"time" json:"time"`
	EmInfo string `orm:"EmInfo" json:"EmInfo"`
}

func (*Emission_Info) TableName() string {
	return "Emission_Info"
}

type Pure_info struct {
	Time     int    `orm:"time" json:"time"`
	Pureinfo string `orm:"pureinfo" json:"pureinfo"`
}

func (*Pure_info) TableName() string {
	return "pure_info"
}

type User_info struct {
	UserId        int    `orm:"user_id" json:"user_id"`
	UserPhone     string `orm:"user_phone" json:"user_phone"`
	UserMail      string `orm:"user_mail" json:"user_mail"`
	UserName      string `orm:"user_name" json:"user_name"`
	UserPrivilege int    `orm:"user_privilege" json:"user_privilege"`
	UserSex       int    `orm:"user_sex" json:"user_sex"`
	UserAge       int    `orm:"user_age" json:"user_age"`
}

func (*User_info) TableName() string {
	return "user_info"
}
