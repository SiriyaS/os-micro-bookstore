package form

type User struct {
	ID        uint64 `json:"id"`
	UserSubID string `json:"user_sub_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserRequest struct {
	UserSubID string `json:"user_sub_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

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

type TokenClaim struct {
	Sub           string `json:"sub"` // user ID
	Azp           string `json:"azp"`
	Aud           string `json:"aud"` // app's client ID
	Exp           string `json:"exp"` // expiry time
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
}

type UserClaim struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}
