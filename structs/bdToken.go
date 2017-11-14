package structs

//百度Token
type BdToken struct {
	Access_token string
	Refresh_token string
	Session_key string
	Scope string
	Session_secret string
	Expires_in int64
}
