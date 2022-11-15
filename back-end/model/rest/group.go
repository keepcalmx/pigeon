package rest

type CreateGroupForm struct {
	Name   string `json:"name" binding:"required"`
	Avatar string `json:"avatar"`
}

type AddGroupUserForm struct {
	UserUUIDs []string `json:"userUUIDs" binding:"required"`
}
