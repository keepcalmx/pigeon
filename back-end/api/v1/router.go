package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	g.POST("/users", CreateUser())
	g.GET("/users/:uuid", GetUserInfo())
	g.POST("/users/:uuid/groups", AddGroups())
	g.GET("/users/:uuid/contacts", GetUserContactInfos())

	g.POST("/groups", CreateGroup())
	g.POST("/groups/:uuid/users", AddGroupUser())
	g.GET("/groups/:uuid/users", ListGroupUsers())
}
