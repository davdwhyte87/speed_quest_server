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
}
func Routes(router *gin.Engine) {
	// initialize controllers
	initializeController()
	userRouter := router.Group("/user")
	userRouter.POST("/create", userController.CreateUser())
	userRouter.POST("/get_code", userController.GetCode())
	userRouter.POST("/login", userController.Login())

	userAuthRouter := router.Group("/auth/user")
	userAuthRouter.Use(middlewares.PlayerAuth(*AuthDAO))
	userAuthRouter.GET("/get", userController.GetUser())

	playerStatsRouter := router.Group("/stats")
	playerStatsRouter.POST("/update/:id", playerStatContoller.UpdatePlayerStats())
	playerStatsRouter.GET("/get/:id", playerStatContoller.GetPlayerStat())

	walletRouter := router.Group("/wallet")
	walletRouter.Use(middlewares.PlayerAuth(*AuthDAO))
	walletRouter.GET("/get", walletController.GetWallet())
	walletRouter.POST("/update", walletController.UpdateWallet())

}
