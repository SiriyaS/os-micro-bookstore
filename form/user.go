package form

type User struct {
	ID        uint64 `json:"id"`
	UserSubID string `json:"user_sub_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   string `json:"picture"`
	Email     string `json:"email"`
	ExpTime   string `json:"exp_time"`
	IsLogin   bool   `json:"is_login"`
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

type VerifyTokenRequest struct {
	ClientID string `json:"client_id"`
	IDToken  string `json:"id_token"`
}

type GoogleClaim struct {
	Iss           string `json:"iss"` // Issuer: accounts.google.com or https://accounts.google.com.
	Sub           string `json:"sub"` // user ID
	Azp           string `json:"azp"`
	Aud           string `json:"aud"` // app's client ID
	Iat           string `json:"iat"`
	Exp           string `json:"exp"` // expiry time
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

type SignOutRequest struct {
	Sub string `json:"sub"`
}
