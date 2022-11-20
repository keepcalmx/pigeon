package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keepcalmx/go-pigeon/cache"
	C "github.com/keepcalmx/go-pigeon/common/constant"
)

func ReadPrivateMsg(uuid string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// pass
	}
}

func ReadGroupMsg(uuid string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// pass
	}
}

func QueryGroupMsg() gin.HandlerFunc {
	return func(c *gin.Context) {
		to := c.Query("to")
		if to == "" {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "缺少群组UUID。",
				"data": nil,
			})
			return
		}
		keywords := c.Query("keywords")
		if keywords == "" {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "缺少关键词。",
				"data": nil,
			})
			return
		}

		offsetStr := c.DefaultQuery("offset", "0")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "分页偏移不合法。",
				"data": nil,
			})
			return
		}
		limitStr := c.DefaultQuery("limit", "")
		limit := C.DEFAULT_MSG_LIMIT
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "单页数量不合法。",
				"data": nil,
			})
			return
		}

		// do redis search
		msgs, total, err := cache.QueryGroupMsg(to, keywords, offset, limit)
		if err != nil {
			log.Println(err)
			// TODO query from mysql
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功。",
			"data": gin.H{
				"total": total,
				"msgs":  msgs,
			},
		})
	}
}

func QueryPrivateMsg() gin.HandlerFunc {
	return func(c *gin.Context) {
		from := c.Query("from")
		if from == "" {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "缺少UUID from。",
				"data": nil,
			})
			return
		}

		to := c.Query("to")
		if to == "" {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "缺少UUID to。",
				"data": nil,
			})
			return
		}

		keywords := c.Query("keywords")
		if keywords == "" {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "缺少关键词。",
				"data": nil,
			})
			return
		}

		offsetStr := c.DefaultQuery("offset", "0")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "分页偏移不合法。",
				"data": nil,
			})
			return
		}

		limitStr := c.DefaultQuery("limit", "")
		limit := C.DEFAULT_MSG_LIMIT
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "单页数量不合法。",
				"data": nil,
			})
			return
		}

		// do redis search
		msgs, total, err := cache.QueryPrivateMsg(from, to, keywords, offset, limit)
		if err != nil {
			log.Println(err)
			// TODO query from mysql
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功。",
			"data": gin.H{
				"total": total,
				"msgs":  msgs,
			},
		})
	}
}
