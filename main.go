package main

import (
	// "context"
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

	router.Run("localhost:6000")
}
