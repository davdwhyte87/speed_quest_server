package routes

// import (
// 	"speed_quest_server/controllers"

// 	"github.com/gin-gonic/gin"
// )

// var (
// 	userController *controllers.UserController
// )

// func initializeController() {
// 	factoryDAO, _ := GetDAO()

// 	userController = &controllers.UserController{
// 		FactoryDAO: factoryDAO,
// 		UserDAO: UserDAO,
// 	}
// }
// func UserRoute(router *gin.Engine) {
// 	// initialize controllers
// 	initializeController()
// 	userRouter := router.Group("/user")
// 	userRouter.POST("/create", userController.CreateUser())
// 	userRouter.POST("/get_code", userController.GetCode())
// 	userRouter.POST("/login", userController.Login())
// }
