package form

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
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

type GoogleTokenClaim struct {
	Sub           string `json:"sub"` // user ID
	Azp           string `json:"azp"`
	Aud           string `json:"aud"` // app's client ID
	Exp           string `json:"exp"` // expiry time
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
}

type GoogleUserClaim struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

type GitHubClaim struct {
	Url   string     `json:"url"`
	Token string     `json:"token"`
	User  GitHubUser `json:"user"`
}

type GitHubUser struct {
	Login     string `json:"login"`
	ID        uint64 `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	HtmlUrl   string `json:"html_url"`
	Type      string `json:"type"`
}
