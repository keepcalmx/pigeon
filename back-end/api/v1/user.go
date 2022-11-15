package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	const_ "github.com/keepcalmx/go-pigeon/common/constant"
	"github.com/keepcalmx/go-pigeon/ent"
	"github.com/keepcalmx/go-pigeon/model/rest"
	"github.com/keepcalmx/go-pigeon/storage"
	"github.com/keepcalmx/go-pigeon/websocket/msg"
	"golang.org/x/crypto/bcrypt"
)

const (
	USER  = "user"
	GROUP = "group"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userForm rest.CreateUserForm
		err := c.BindJSON(&userForm)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "参数错误。",
				"data": nil,
			})
			return
		}

		existUser, _ := storage.GetUserByUsername(userForm.Username)
		if existUser != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "已存在该用户名。",
				"data": nil,
			})
			return
		}

		hashedPwd, err := bcrypt.GenerateFromPassword(
			[]byte(userForm.Password), bcrypt.DefaultCost)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "密码加密错误。",
				"data": nil,
			})
			return
		}
		userForm.Password = string(hashedPwd)
		user, err := storage.CreateUser(userForm)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "保存用户信息失败。",
				"data": nil,
			})
			return
		}

		storage.AddGroupUser(const_.SYSTEM_GROUP_UUID, []string{user.UUID})

		c.IndentedJSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "创建用户成功。",
			"data": map[string]any{
				"uuid":     user.UUID,
				"username": user.Username,
			},
		})
	}
}

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		UUID, exist := c.Params.Get("uuid")
		if !exist {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "用户UUID不存在。",
				"data": nil,
			})
		}

		user, err := storage.GetUserByUUID(UUID)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "获取用户信息失败。",
				"data": nil,
			})
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}

func AddGroups() gin.HandlerFunc {
	return func(c *gin.Context) {
		// pass
	}
}

func GetDisplayName(user *ent.User) string {
	if user.Nickname != "" {
		return user.Nickname
	}
	return user.Username
}

func GetUserContactInfos() gin.HandlerFunc {
	return func(c *gin.Context) {
		UUID, exist := c.Params.Get("uuid")
		if !exist {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "用户UUID不存在。",
				"data": nil,
			})
		}
		// TODO 暂时显示所有用户
		users, err := storage.GetAllUsers()
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "获取好友列表失败。",
				"data": nil,
			})
		}
		allGroups, err := storage.GetUserGroups(UUID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "获取群聊列表失败",
				"data": nil,
			})
		}

		var data []rest.ContactInfo

		for _, user := range users {
			if user.UUID == UUID {
				continue
			}
			msgs, _ := storage.ListPrivateMsgs(UUID, user.UUID)
			data = append(data, rest.ContactInfo{
				UUID:        user.UUID,
				Type:        USER,
				DisplayName: GetDisplayName(user),
				Online:      msg.GetHub().IsOnline(user.UUID),
				Avatar:      user.Avatar,
				RecentMsgs:  msgs,
			})
		}

		for _, group := range allGroups {
			msgs, _ := storage.GetGroupMsgs(group.UUID)
			data = append(data, rest.ContactInfo{
				UUID:        group.UUID,
				Type:        GROUP,
				DisplayName: group.Name,
				Online:      true,
				Avatar:      group.Avatar,
				RecentMsgs:  msgs,
			})
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": data,
		})
	}
}
