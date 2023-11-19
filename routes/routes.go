package routes

import (
	"speed_quest_server/controllers"
	"speed_quest_server/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	userController      *controllers.UserController
	playerStatContoller *controllers.PlayerStatContoller
	walletController    *controllers.WalletController
	serverController    *controllers.ServerController
)

func initializeController() {
	factoryDAO, _ := GetDAO()

	userController = &controllers.UserController{
		FactoryDAO: factoryDAO,
		UserDAO:    UserDAO,
		AuthDAO:    AuthDAO,
	}
	playerStatContoller = &controllers.PlayerStatContoller{
		FactoryDAO:     factoryDAO,
		PlayerStatsDAO: PlayerStatsDAO,
	}
	walletController = &controllers.WalletController{
		FactoryDAO: factoryDAO,
		WalletDAO:  WalletDAO,
	}
	serverController = &controllers.ServerController{}
}
func Routes(router *gin.Engine) {
	// initialize controllers
	initializeController()
	router.GET("/leader_board", playerStatContoller.GetLeaderBoard())
	router.GET("/test_email", playerStatContoller.TestEmail())
	userRouter := router.Group("/user")
	userRouter.POST("/create", userController.CreateUser())
	userRouter.POST("/get_code", userController.GetCode())
	userRouter.POST("/login", userController.Login())

	userAuthRouter := router.Group("/auth/user")
	userAuthRouter.Use(middlewares.PlayerAuth(*AuthDAO))
	userAuthRouter.GET("/get", userController.GetUser())

	playerStatsRouter := router.Group("/stats")
	playerStatsRouter.Use(middlewares.PlayerAuth(*AuthDAO)).POST("/update", playerStatContoller.UpdatePlayerStats())
	playerStatsRouter.Use(middlewares.PlayerAuth(*AuthDAO)).GET("/get", playerStatContoller.GetPlayerStat())

	walletRouter := router.Group("/wallet")
	walletRouter.Use(middlewares.PlayerAuth(*AuthDAO))
	walletRouter.GET("/get", walletController.GetWallet())
	walletRouter.POST("/update", walletController.UpdateWallet())

	serverRouter := router.Group("/server")
	serverRouter.GET("/latest_android_version", serverController.GetLatestAndroidAppVersion())
}
