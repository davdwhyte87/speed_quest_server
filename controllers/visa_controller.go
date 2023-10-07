package controllers

import (
	// "context"
	"log"
	"speed_quest_server/configs"
	"speed_quest_server/dao"
	"speed_quest_server/models"
	"speed_quest_server/requests"
	"speed_quest_server/responses"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var visaApplicationAnswersCollection *mongo.Collection = configs.GetCollection(configs.DB, "VisaApplicationAnswers")
var visaApplicationCollection *mongo.Collection = configs.GetCollection(configs.DB, "VisaApplications")

type VisaController struct {
	FactoryDAO         *dao.FactoryDAO
	UserDAO            *dao.UserDAO
	RoleDAO            *dao.RoleDAO
	VisaApplicationDAO *dao.VisaApplicationDAO
}

func (visaController VisaController) CreateVisaApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var visaApplicationRequest requests.VisaApplicationRequest
		err := c.BindJSON(&visaApplicationRequest)

		// validate request body
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// validate data
		validationError := validate.Struct(&visaApplicationRequest)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation Error",
				Data:    map[string]interface{}{"data": validationError.Error()},
			})
			return
		}
		// create application id
		applicationId := primitive.NewObjectID()

		// assign id to all questions answers
		applicationAnswers := make([]interface{}, 0)
		for _, ans := range visaApplicationRequest.ApplicationAnswers {
			var applicationAnswer models.ApplicationAnswer

			qid, err := primitive.ObjectIDFromHex(ans.QuestionId)
			if err != nil {
				qid = primitive.NewObjectID()
			}
			applicationAnswer.ApplicationId = applicationId
			applicationAnswer.Id = primitive.NewObjectID()
			applicationAnswer.QuestionId = qid
			applicationAnswer.TextAnswer = ans.TextAnswer
			applicationAnswer.YesNoAnswer = ans.YesNoAnswer
			applicationAnswers = append(applicationAnswers, applicationAnswer)

		}

		// create database context
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()
		log.Printf("ans", len(visaApplicationRequest.ApplicationAnswers))
		// insert all the answers to the questions
		// visaApplicationAnswersCollection.InsertMany(ctx, applicationAnswers)
		insertAnsErr := visaController.FactoryDAO.InsertMany(models.ApplicationAnswerCollection, applicationAnswers)
		if insertAnsErr != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    map[string]interface{}{"data": insertAnsErr.Error()},
			})
			return
		}

		// insert new application
		var visaApplication models.VisaApplication
		visaApplication.Id = applicationId
		visaApplication.Email = visaApplicationRequest.Email
		visaApplication.Phone = visaApplicationRequest.Phone
		visaApplication.Name = visaApplicationRequest.Name
		visaApplication.Location = visaApplicationRequest.Location
		visaApplication.Profession = visaApplicationRequest.Profession

		// result, _ := visaApplicationCollection.InsertOne(ctx, visaApplication)

		insertErr := visaController.FactoryDAO.Insert(models.VisaApplicationCollection, visaApplication)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    map[string]interface{}{"data": insertErr.Error()},
			})
			return
		}
		// send response
		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "Created Application",
			Data:    map[string]interface{}{"data": "OK"},
		})
	}
}

func (visaController VisaController) ApproveVisaApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get id from request
		applicationId := c.Param("id")
		var visaApplication models.VisaApplication
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()
		// err := visaApplicationCollection.FindOne(ctx, bson.M{"_id": applicationId}).Decode(&visaApplication)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		//ctxUpdate, cancelUpdate := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancelUpdate()
		// update the visa application
		//updateResult, updateErr := visaApplicationCollection.UpdateOne(ctxUpdate, bson.M{"_id"}, &visaApplication)

		// get the role data from the jwt
		//TODO JWT

		// get visa application from database
		va, _ := visaController.FactoryDAO.Get(models.VisaApplicationCollection, applicationId)
		e := va.Decode(&visaApplication)
		if e != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Server Error",
				Data:    map[string]interface{}{"Data": "Error Casting"},
			})
			return
		}

		roleID := 1

		// get the particular user role from the database
		role := visaController.RoleDAO.GetUserRole(roleID)
		// check if the user can perform that specific action
		if !role.CanApproveVisaApplications {
			c.JSON(http.StatusUnauthorized, responses.GenericResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    map[string]interface{}{"Data": "No permission"},
			})
			return
		}

		// update visa application
		visaApplication.Status = models.ApplicationStatus(2)
		err := visaController.FactoryDAO.Update(models.VisaApplicationCollection, applicationId, visaApplication)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Server Error",
				Data:    map[string]interface{}{"Data": "Error updating"},
			})
			return
		}
		// done...
		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "Approved",
			Data:    map[string]interface{}{"Data": applicationId},
		})
	}
}

// get all the visa applications
func (visacontroller *VisaController) GetAllVisaApplications() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := visacontroller.VisaApplicationDAO.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Server Error",
				Data:    map[string]interface{}{"Data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "Ok",
			Data:    map[string]interface{}{"Data": result},
		})
		return
	}
}

// get a single visa application
func (visacontroller *VisaController) GetSingleApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		applicationID := c.Param("id")

		applicationIDOBJ, objectIDErr := primitive.ObjectIDFromHex(applicationID)
		if objectIDErr != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    map[string]interface{}{"Data": objectIDErr.Error()},
			})
			return
		}

		// get single from db
		result, err := visacontroller.VisaApplicationDAO.GetSingle(applicationIDOBJ)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Server Error",
				Data:    map[string]interface{}{"Data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "Ok",
			Data:    map[string]interface{}{"Data": result},
		})
		return
	}
}
