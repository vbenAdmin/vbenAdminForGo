package Models

type Users struct {
	Common
	//UserType int 	`gorm:"column:user_type;" json:"user_type"` // 用户类型： 1管理员 2店铺 3用户
	Account  string `gorm:"column:account;type:varchar(50)" form:"account" json:"account,omitempty" binding:"required,max=30"`
	Password string `gorm:"column:password;type:varchar(50)" form:"password" json:"password,omitempty" binding:"required,max=50"`
	Avatar   string `gorm:"column:avatar;type:varchar(255)" json:"avatar,omitempty"`
	Token    string `json:"token,omitempty"'`
}

func (t *Users) TableName() string {
	return "users"
}
