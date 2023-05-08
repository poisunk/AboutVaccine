package schema

type UserInfo struct {
	Uid      int64  `json:"uid"`
	Nickname string `json:"nickname"`
}

type UserClaim struct {
	Uid      int64  `json:"uid"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

type UserAdd struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
