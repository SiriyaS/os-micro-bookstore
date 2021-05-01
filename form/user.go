package form

type UserInfo struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type UserInfoRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
