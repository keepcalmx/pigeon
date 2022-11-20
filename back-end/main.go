package main

import (
	"flag"
	"time"

	"github.com/gin-contrib/cors"
	apiV1 "github.com/keepcalmx/go-pigeon/api/v1"
	"github.com/keepcalmx/go-pigeon/cache"
	"github.com/keepcalmx/go-pigeon/storage"
	"github.com/keepcalmx/go-pigeon/websocket/msg"

	"github.com/gin-gonic/gin"
)

var RunMode = flag.String("mode", "debug", "server run mode")

func init() {
	flag.Parse()

	if (*RunMode) == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	storage.Migrate()
	cache.CacheGroupMsg()
	cache.CachePrivateMsg()

	// websocket client hub
	go msg.GetHub().Run()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	v1 := r.Group("/api/v1")
	apiV1.RegisterRouter(v1)

	r.Run(":8000")
}
