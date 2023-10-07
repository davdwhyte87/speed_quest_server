package routes

import (
	"speed_quest_server/controllers"

	"github.com/gin-gonic/gin"
)

var visaController *controllers.VisaController

func initControllers(){
	factoryDAO, _ := GetDAO()
	visaController = &controllers.VisaController{
		FactoryDAO: factoryDAO,
		UserDAO: UserDAO ,
		RoleDAO: RoleDAO,
		VisaApplicationDAO: VisaApplicationDAO,
	}
}
func VisaRoutes(router *gin.Engine){
	initControllers()
	router.POST("/visa_application", visaController.CreateVisaApplication())
	router.GET("/approve_visa_application/:id", visaController.ApproveVisaApplication())
	router.GET("visa_applications", visaController.GetAllVisaApplications())	
	router.GET("/visa_application/:id", visaController.GetSingleApplication())
}