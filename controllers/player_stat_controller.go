package controllers

import (
	"net/http"
	"speed_quest_server/dao"
	"speed_quest_server/requests"
	"speed_quest_server/responses"
	"speed_quest_server/utils"

	"github.com/gin-gonic/gin"
)

type PlayerStatContoller struct {
	FactoryDAO     *dao.FactoryDAO
	PlayerStatsDAO *dao.PlayerStatsDAO
}

// get a players stats
func (psc *PlayerStatContoller) GetPlayerStat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// user_id := ctx.Param("id")
		// utils.Logger.Debug().Msg(user_id)
		// get player stats
		// get user id from context
		user_id:=ctx.GetString("user_id")
		player_stats, err := psc.PlayerStatsDAO.GetByUserId(user_id)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Data:    nil,
				Message: "Error getting player stats",
			})
			return
		}
		ctx.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Data:    map[string]interface{}{"data": player_stats},
			Message: "ok",
		})
	}
}

// update player stats
func (psc *PlayerStatContoller) UpdatePlayerStats() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateStatsReq requests.UpdatePlayerStatsReq
		// user_id := ctx.Param("id")
		// get user id 
		user_id :=ctx.GetString("user_id")
		if user_id == "" {
			ctx.JSON(http.StatusBadRequest, responses.GenericResponse{
				Status:  http.StatusBadRequest,
				Message: "request id error",
				Data:    nil,
			})
		}
		err := ctx.BindJSON(&updateStatsReq)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			ctx.JSON(http.StatusBadRequest, responses.GenericResponse{
				Status:  http.StatusBadRequest,
				Message: "request body error",
				Data:    nil,
			})
			return
		}
		// validate data

		// get player stats
		player_stats, err := psc.PlayerStatsDAO.GetByUserId(user_id)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Data:    nil,
				Message: "Error getting player stats",
			})
			return
		}

		// if updateStatsReq.HighScore != 0 {
		// 	player_stats.HighScore = updateStatsReq.HighScore
		// }
		if updateStatsReq.LastScore != 0 {
			// utils.Logger.Debug().Int64("last score",updateStatsReq.LastScore).Int64("highscore",player_stats.HighScore).Msg("")
			if (updateStatsReq.LastScore > player_stats.HighScore) {
				player_stats.HighScore = updateStatsReq.LastScore
			}else {
				// utils.Logger.Debug().Msg("No need to set high score")
			}
			
			player_stats.LastScore = updateStatsReq.LastScore
		}

		err = psc.PlayerStatsDAO.Update(player_stats.Id, player_stats)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Data:   nil ,
				Message: "Error updating stats",
			})
			return 
		}
		
		ctx.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Data:    map[string]interface{}{"data": player_stats},
			Message: "ok",
		})
	}
}


