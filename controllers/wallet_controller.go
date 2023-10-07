package controllers

import (
	"net/http"
	"speed_quest_server/dao"
	"speed_quest_server/utils"

	"speed_quest_server/requests"
	"speed_quest_server/responses"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	FactoryDAO *dao.FactoryDAO
	WalletDAO *dao.WalletDAO
}


// get a players wallet
func (wc *WalletController) GetWallet()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// get user id from contect
		userId:=ctx.GetString("user_id")
		// userId := ctx.Param("id")

		// get wallet from db
		wallet, err := wc.WalletDAO.GetByUserId(userId)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status: http.StatusInternalServerError,
				Message: "Error getting user wallet",
				Data: nil,
			})
			return
		}

		ctx.JSON(http.StatusOK, responses.GenericResponse{
			Status: http.StatusOK,
			Message: "Ok",
			Data: map[string]interface{}{"data":wallet},
		})

	}
}

// wallet update balance 
func (wc *WalletController) UpdateWallet() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// get user id 
		userId:=ctx.GetString("user_id")
		// get request data 
		var updateWalletRequest requests.UpdateWalletReq
		err := ctx.BindJSON(&updateWalletRequest)
		if err!=nil {
			utils.Logger.Error().Msg(err.Error())
			ctx.JSON(http.StatusBadRequest, responses.GenericResponse{
				Status: http.StatusBadRequest,
				Message: "Error with request",
				Data: nil,
			})
			return
		}

		// validate data
		validationError := validate.Struct(&updateWalletRequest)
		if validationError != nil {
			ctx.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation Error",
				Data:    map[string]interface{}{"data": validationError.Error()},
			})
			return
		}

		// get wallet from db
		wallet, err := wc.WalletDAO.GetByUserId(userId)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status: http.StatusInternalServerError,
				Message: "Error getting user wallet",
				Data: nil,
			})
			return
		}

		wallet.Balance = updateWalletRequest.Amount

		// update wallet in database 
		err = wc.WalletDAO.Update(wallet.Id, wallet)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status: http.StatusInternalServerError,
				Message: "Error getting user wallet",
				Data: nil,
			})
			return
		}


		ctx.JSON(http.StatusOK, responses.GenericResponse{
			Status: http.StatusOK,
			Message: "Ok",
			Data: nil,
		})
	}
}