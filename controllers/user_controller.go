package controllers

import (
	// "context"

	"os"
	"speed_quest_server/configs"
	"speed_quest_server/dao"
	"speed_quest_server/models"
	"speed_quest_server/requests"
	"speed_quest_server/responses"
	"speed_quest_server/utils"
	"time"

	// "time"

	"net/http"

	// "firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New(validator.WithRequiredStructEnabled())
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "Users")

type UserController struct {
	FactoryDAO *dao.FactoryDAO
	UserDAO    *dao.UserDAO
	AuthDAO    *dao.AuthDAO
}

func (userController *UserController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUserRequest requests.CreateUserReq
		err := c.BindJSON(&createUserRequest)

		// validate request body
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		// validate data
		validationError := validate.Struct(&createUserRequest)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation Error",
				Data:    map[string]interface{}{"data": validationError.Error()},
			})
			return
		}

		// check if user already exists
		userExist, err := userController.UserDAO.DoesUserExist(createUserRequest.Email, createUserRequest.UserName)
		if userExist {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "User exists",
				Data:    nil,
			})
			return
		}
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Error checking user state",
				Data:    nil,
			})
			return
		}

		// create timeout context for db
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()

		// prep user data
		newUser := models.User{
			Id:         uuid.NewString(),
			UserName:   createUserRequest.UserName,
			Email:      createUserRequest.Email,
			UserRoleId: 0,
		}

		// create players stats
		playerStats := models.PlayerStats{}
		playerStats.Id = uuid.NewString()
		playerStats.CreateAt = time.Now().String()
		playerStats.UpdatedAt = time.Now().String()
		playerStats.HighScore = 0
		playerStats.LastScore = 0
		playerStats.UserID = newUser.Id
		playerStats.UserName = newUser.UserName
		err = userController.FactoryDAO.Insert(models.PlayerStatsCollection, playerStats)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error creating user ",
				Data:    nil,
			})
			return
		}
		// create player wallet
		wallet := models.Wallet{}
		wallet.Id = uuid.NewString()
		wallet.Balance = 0
		wallet.CreatedAt = time.Now().String()
		wallet.UpdatedAt = time.Now().String()
		wallet.UserId = newUser.Id
		wallet.Address = newUser.UserName
		wallet.Email = newUser.Email
		wallet.Version = 0.1
		// create new block
		var block models.Block 
		block.Amount = 0
		block.Date = time.Now().String()
		block.Id = uuid.NewString()
		block.PreviousHash = "000000000000"
		block.ReceiverAddress = wallet.Address
		block.SenderAddress = "000000000000"

		// add block to wallet blocks 
		wallet.Blocks = append(wallet.Blocks, block)
		
		err = userController.FactoryDAO.Insert(models.WalletCollection, wallet)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error creating user ",
				Data:    nil,
			})
			return
		}
		// insert user data into db
		// result, err := userCollection.InsertOne(ctx,newUser )
		insertErr := userController.FactoryDAO.Insert(models.UserCollection, newUser)
		if insertErr != nil {
			utils.Logger.Error().Err(insertErr).Msg(insertErr.Error())
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    nil})
			return
		}

		// send user welcome email
		if !userController.sendWelcomeEmail(newUser) {
			c.JSON(http.StatusOK, responses.UserResponse{
				Status:  http.StatusOK,
				Data:    map[string]interface{}{"data": "OK"},
				Message: "User created, email send error",
			})
		}

		// return data
		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Data:    map[string]interface{}{"data": "OK"},
			Message: "User created",
		})

	}
}


func (usc *UserController) sendWelcomeEmail(user models.User) bool{
	data := utils.EmailData{}
	data.ContentData = map[string]interface{}{"Name": user.UserName, "Code": user.Code, "URL":os.Getenv("API_URL")}
	data.EmailTo = user.Email
	data.Template ="welcome_email/welcome_email.html"
	data.Title ="Welcome To SpeedQuest"

	err :=utils.SendEmail(data)
	if err != nil {
		return false
	}

	return true
}

// this function helps a user get his authentication code via mail
func (userController *UserController) GetCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getCodeRequest requests.GetCodeReq
		err := c.BindJSON(&getCodeRequest)
		if err != nil {
			utils.Logger.Error().Err(err).Msg(err.Error())
			c.JSON(http.StatusBadRequest, responses.GenericResponse{
				Status:  http.StatusBadRequest,
				Message: "Bad request",
				Data:    nil,
			})
			return
		}
		// validate data
		validationError := validate.Struct(&getCodeRequest)

		if validationError != nil {
			utils.Logger.Error().Msg(validationError.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation Error",
				Data:    map[string]interface{}{"data": validationError.Error()},
			})
			return
		}
		// get user dataa from database
		user, err := userController.UserDAO.GetByEmail(getCodeRequest.Email)
		if err != nil {
			utils.Logger.Error().Err(err).Msg(err.Error())
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "error getting code",
				Data:    nil,
			})
			return
		}
		if user.Email == "" {
			c.JSON(http.StatusNotFound, responses.GenericResponse{
				Status:  http.StatusNotFound,
				Message: "User not found",
				Data:    nil,
			})
			return
		}
		code := "00000"
		println(code)
		user.Code = code
		// update the user with code in database
		err = userController.UserDAO.UpdateByEmail(user.Email, user)
		if err != nil {
			utils.Logger.Error().Err(err)
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Could not get code",
				Data:    nil,
			})
		}

		// send the code to users email
		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "Code sent to email",
			Data:    nil,
		})
	}
}

// helps users login with email and code
func (userController *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest requests.LoginReq
		// perse request data
		err := c.BindJSON(&loginRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GenericResponse{
				Status:  http.StatusBadRequest,
				Message: "Error with request data",
				Data:    nil,
			})
			return
		}

		// validate data
		validationError := validate.Struct(&loginRequest)

		if validationError != nil {
			utils.Logger.Error().Msg(validationError.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation Error",
				Data:    map[string]interface{}{"data": validationError.Error()},
			})
			return
		}

		// get user data
		user, err := userController.UserDAO.GetByEmail(loginRequest.Email)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "error getting code",
				Data:    nil,
			})
			return
		}
		// check if code matches
		if loginRequest.Code != user.Code {
			c.JSON(http.StatusUnauthorized, responses.GenericResponse{
				Status:  http.StatusUnauthorized,
				Data:    nil,
				Message: "Unauthorized",
			})
			return
		}

		// generate a user token
		token := uuid.New()
		// insert token into database
		var auth models.Auth
		auth.Token = token.String()
		auth.UserName = user.UserName
		auth.UserId = user.Id

		// delete token that exists
		userController.AuthDAO.DeleteByUserName(user.UserName)
		inautherr := userController.FactoryDAO.Insert(models.AuthCollection, auth)
		if inautherr != nil {
			utils.Logger.Error().Msg(inautherr.Error())
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Data:    nil,
				Message: "Error logging in",
			})
			return
		}

		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Data:    map[string]interface{}{"token": token, "user":user.Safe()},
			Message: "Logged in!",
		})
	}
}


// get user data for a logged in user 
func (uc *UserController) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get user id
		userId := ctx.GetString("user_id")

		//  get user data from db
		user, err := uc.UserDAO.GetByUserId(userId)
		if err != nil {
			utils.Logger.Error().Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Data:    nil,
				Message: "Error logging in",
			})
			return
		}

		ctx.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Data:    map[string]interface{}{"data": user.Safe()},
			Message: "Ok",
		})
	}
}
