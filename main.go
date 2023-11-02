package main

import (
	// "context"
	"os"
	"speed_quest_server/configs"
	"speed_quest_server/utils"

	// "speed_quest_server/dao"
	"speed_quest_server/routes"

	"github.com/gin-gonic/gin"
)

// var (
// 	FactoryDAO  *dao.FactoryDAO

// )
func main() {
	// setup logging
	utils.InitLogger()

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "Armorgens API running ..."})
	})

	//run database
	configs.ConnectDB()

	//
	// FactoryDAO = dao.InitializeFactory(configs.DB, context.TODO())

	// setup data access objects
	routes.SetupDAO()
	// routes setup
	routes.Routes(router)
	routes.VisaRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run("localhost:" + port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
