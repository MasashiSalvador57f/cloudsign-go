package cloudsign

import "time"

// AccessToken is represents a access token of cloudsign api.
type AccessToken struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   uint      `json:"expires_in"`
	CreatedAt   time.Time `json:"created_at"`
}

// IsExpired represnets access token expiration.
func (a *AccessToken) IsExpired() bool {
	d := time.Duration(a.ExpiresIn) * time.Second
	expireAt := a.CreatedAt.Add(d)
	return time.Now().Before(expireAt)
}
