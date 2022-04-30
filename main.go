package main

import (
	"os"

	"bamboo-api/app/clients/database"
	"bamboo-api/app/routers"

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
	routers.InitRouters(r)
	database.InitMysql()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
