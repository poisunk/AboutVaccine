package entity

type User struct {
	Id       int64  `xorm:"notnull pk autoincr INT(11) id" json:"id"`
	Nickname string `xorm:"null VARCHAR(255) nickname" json:"nickname"`
	Password string `xorm:"null VARCHAR(255) password" json:"password"`
}

func (u *User) TableName() string {
	return "user"
}
