package controllers

import (
	"net/http"
	"os"
	"speed_quest_server/responses"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
}

// chekc if there is a new version
func (sc *ServerController) GetLatestAndroidAppVersion() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get android app version
		latestAndroidAppVersion := os.Getenv("LATEST_ANDROID_APP_VERSION")
		ctx.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "Ok",
			Data:    map[string]interface{}{"data": latestAndroidAppVersion},
		})
	}
}
