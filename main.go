package main

import (
	"os"
	"time"

	"bamboo-api/app/clients"

	"bamboo-api/app/routers"

	"bamboo-api/app/clients/database"
	"bamboo-api/app/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
}

func main() {
	r := gin.Default()
	env := "dev"
	if os.Getenv("ENV") == "prod" {
		env = "prod"
	}

	config.Init(env)
	database.InitMysql()
	clients.InitCache()

	if env == "dev" {
		store, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("redis_secret"))
		if nil != err {
			log.Errorf("session init redis stroe failed: err=%+v", err)
		}
		r.Use(sessions.Sessions("session", store))
	} else {
		store, err := sessions.NewRedisStore(10, "tcp", "bamboo-redis.e1juwh.ng.0001.apne1.cache.amazonaws.com:6379", "", []byte("redis_secret"))
		if nil != err {
			log.Errorf("session init redis stroe failed: err=%+v", err)
		}
		r.Use(sessions.Sessions("session", store))
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://aaa.bamboownft.com", "https://www.bamboownft.com"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "WalletId", "Accept", "User-Agent", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           12 * time.Hour,
	}))
	r.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	routers.InitRouters(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
