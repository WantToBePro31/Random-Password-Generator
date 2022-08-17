package main

import (
	"github.com/WantToBePro31/rand-pass/config"
	"github.com/WantToBePro31/rand-pass/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	defer config.DisconnectDB()
	config.ConnectDB()
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	routes.SetUpRoutes(server)
	server.Run()
}
