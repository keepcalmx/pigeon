package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keepcalmx/go-pigeon/model/rest"
	"github.com/keepcalmx/go-pigeon/storage"
)

func CreateGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var groupForm rest.CreateGroupForm
		err := c.BindJSON(&groupForm)
		if err != nil {
			log.Fatalln(err)
		}

		group, err := storage.CreateGroup(groupForm.Name, groupForm.Avatar)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "保存群组信息失败。",
				"data":    nil,
			})
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    group,
		})
	}
}

func AddGroupUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		groupUUID := c.Param("uuid")

		var form rest.AddGroupUserForm
		err := c.BindJSON(&form)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "参数错误。",
				"data":    nil,
			})
		}
		group, err := storage.AddGroupUser(groupUUID, form.UserUUIDs)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "保存群成员失败。",
				"data":    nil,
			})
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    group,
		})
	}
}

func ListGroupUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")

		users, err := storage.ListGroupUsers(uuid)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "查询群成员失败。",
				"data":    nil,
			})
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    users,
		})
	}
}

func GetGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		// pass
	}
}

func RemoveGroupUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// pass
	}
}

func DeleteGroupByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// pass
	}
}

func UpdateGroupByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// pass
	}
}
