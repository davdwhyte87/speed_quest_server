package routes

import (
	"context"
	"speed_quest_server/configs"
	"speed_quest_server/dao"
	"speed_quest_server/models"
)

// this contains all the setups for data acess for all application routes


var factoryDAO *dao.FactoryDAO
var UserDAO *dao.UserDAO
var FactoryDAO *dao.FactoryDAO
var RoleDAO *dao.RoleDAO
var VisaApplicationDAO *dao.VisaApplicationDAO
var PlayerStatsDAO *dao.PlayerStatsDAO
var WalletDAO *dao.WalletDAO
var AuthDAO *dao.AuthDAO

func GetDAO () (*dao.FactoryDAO, *dao.UserDAO){
	return factoryDAO, UserDAO
}

func SetupDAO (){
	factoryDAO = dao.InitializeFactory(configs.DB, context.TODO())
	UserDAO = &dao.UserDAO{
		 Collection: configs.GetCollection(configs.DB, models.UserCollection),
		 Context: context.TODO(),
	}

	RoleDAO = &dao.RoleDAO{
		Collection: configs.GetCollection(configs.DB, models.RoleCollection),
		Context: context.TODO(),	
	}
	VisaApplicationDAO = &dao.VisaApplicationDAO{
		Collection: configs.GetCollection(configs.DB, models.VisaApplicationCollection),
		Context: context.TODO(),	
	}

	PlayerStatsDAO = &dao.PlayerStatsDAO{
		Collection: configs.GetCollection(configs.DB, models.PlayerStatsCollection),
		Context: context.TODO(),	
	}

	WalletDAO = &dao.WalletDAO{
		Collection: configs.GetCollection(configs.DB, models.WalletCollection),
		Context: context.TODO(),	
	}

	AuthDAO = &dao.AuthDAO{
		Collection: configs.GetCollection(configs.DB, models.AuthCollection),
		Context: context.TODO(),	
	}

	
}