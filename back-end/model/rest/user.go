package rest

type CreateUserForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type ContactInfo struct {
	UUID        string `json:"uuid"`
	Type        string `json:"type"`
	DisplayName string `json:"displayName"`
	Online      bool   `json:"online"`
	Avatar      string `json:"avatar"`
	RecentMsgs  any    `json:"recentMsg"`
}
