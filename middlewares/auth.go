package middlewares

import (

	"net/http"
	"speed_quest_server/dao"
	"speed_quest_server/responses"
	"speed_quest_server/utils"

	"github.com/gin-gonic/gin"
)


func useAuth(nextHandler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// utils.RespondWithError(w, http.StatusUnauthorized, "An authorized error occurred")
	})
}

func PlayerAuth(dao dao.AuthDAO) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		authToken :=ctx.Request.Header.Get("Authorization")
		if authToken == ""{
			ctx.JSON(http.StatusUnauthorized, responses.GenericResponse{
				Status: http.StatusUnauthorized,
				Data: nil,
				Message: "Authorization error",
			})
			ctx.Abort()
			return
		}

		// checking the user token
		auth, err :=dao.GetByToken(authToken)
		if err!= nil {
			utils.Logger.Error().Msg(err.Error())
			ctx.JSON(http.StatusUnauthorized, responses.GenericResponse{
				Status: http.StatusUnauthorized,
				Data: nil,
				Message: "Authorization error",
			})
			ctx.Abort()
			return
		}

		// insert data into request context
		ctx.Set("user_id", auth.UserId)
		ctx.Set("user_name", auth.UserName)
		
	}
}