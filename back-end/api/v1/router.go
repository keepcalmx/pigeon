package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/keepcalmx/go-pigeon/common/utils"
	"github.com/keepcalmx/go-pigeon/websocket/msg"
)

func RegisterRouter(g *gin.RouterGroup) {
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g.GET("/msg", func(c *gin.Context) {
		msg.HandleRequest(c.Writer, c.Request)
	})

	g.POST("/login", Login())

	g.POST("/register", Register())
	g.GET("/users/:uuid", JWTAuthRequired(), GetUserInfo())
	g.POST("/users/:uuid/groups", JWTAuthRequired(), AddGroups())
	g.GET("/users/:uuid/contacts", JWTAuthRequired(), GetUserContactInfos())

	g.POST("/groups", JWTAuthRequired(), CreateGroup())
	g.POST("/groups/:uuid/users", JWTAuthRequired(), AddGroupUser())
	g.GET("/groups/:uuid/users", JWTAuthRequired(), ListGroupUsers())

	g.GET("/group/msg", JWTAuthRequired(), QueryGroupMsg())
	g.GET("/user/msg", JWTAuthRequired(), QueryPrivateMsg())
}

func JWTAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "需要token。",
			})
			return
		}

		segments := strings.SplitN(auth, " ", 2)
		if len(segments) != 2 || segments[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token需以Beare开头，空格隔开。",
			})
			return
		}

		myClaims, err := utils.ParseToken(segments[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "无效的token。",
			})
			return
		}

		c.Set("uuid", myClaims.UUID)
		c.Next()
	}
}
