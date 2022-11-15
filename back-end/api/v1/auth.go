package v1

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/keepcalmx/go-pigeon/common/utils"
	"github.com/keepcalmx/go-pigeon/model/rest"
	"github.com/keepcalmx/go-pigeon/storage"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginForm rest.LoginForm
		err := c.BindJSON(&loginForm)
		if err != nil {
			panic(err)
		}

		userInfo, err := storage.GetUserByUsername(loginForm.Username)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 404,
				"msg":  "没有找到拥有此用户名的用户。",
				"data": nil,
			})
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(userInfo.Password),
			[]byte(loginForm.Password),
		)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "账号或用户名错误，验证失败。",
				"data": nil,
			})
		}

		token, err := utils.GenToken(userInfo.UUID)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "生成Token失败。",
				"data": nil,
			})
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证通过。",
			"data": map[string]any{
				"userUUID": userInfo.UUID,
				"displayName": func() string {
					if userInfo.Nickname != "" {
						return userInfo.Nickname
					}
					return userInfo.Username
				}(),
				"displayAvatar": userInfo.Avatar,
				"token":         token,
			},
		})
	}
}
