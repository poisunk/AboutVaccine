package service

type UserService interface {
	// Login 登录
	Login(User) (string, error)
	// Register 注册
	Register(User) (string, error)
	// Logout 注销
	Logout(string) error
	// Status 刷新登录状态
	Status(string) (string, error)
	GetUserList(int, int) ([]*User, int64, error)
	GetUser(int64) (*User, error)
}

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password,omitempty"`
}
