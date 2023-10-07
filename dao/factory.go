package dao

import (
	"context"
	"errors"
	"speed_quest_server/configs"
	"speed_quest_server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FactoryDAO struct {
	db          *mongo.Client
	ctx         context.Context
	Collections map[string]*mongo.Collection
}

// creates and setups a new dao factory
func InitializeFactory(db *mongo.Client, ctx context.Context) *FactoryDAO {
	collectionList := []string{models.UserCollection, models.ApplicationAnswerCollection,
		models.ApplicationQuestionCollection,
		models.RoleCollection,
		models.VisaApplicationCollection,
		models.AuthCollection,
		models.PlayerStatsCollection,
		models.WalletCollection,
	}
	collections := make(map[string]*mongo.Collection)
	for _, key := range collectionList {
		col := configs.GetCollection(db, key)
		collections[key] = col
		// collections = append(collections[], map[string]*mongo.Collection{"":col})
	}
	return &FactoryDAO{
		db:          db,
		ctx:         context.TODO(),
		Collections: collections,
	}
}

func (dao *FactoryDAO) Insert(key string, data interface{}) error {
	collection, ok := dao.Collections[key]
	if !ok {
		return errors.New("invalid collection")
	}

	c, _ := bson.Marshal(data)

	_, err := collection.InsertOne(dao.ctx, c)
	return err
}

// insert many data into a single collection
func (dao *FactoryDAO) InsertMany(key string, data []interface{}) error {
	collection, ok := dao.Collections[key]
	if !ok {
		return errors.New("invalid collection")
	}

	// c, _ := bson.Marshal(data)

	_, err := collection.InsertMany(dao.ctx, data)
	return err
}

// get resrouce by id from any collection
func (dao *FactoryDAO) Get(key string, id string) (*mongo.SingleResult, error) {
	collection, ok := dao.Collections[key]
	if !ok {
		return nil, errors.New("invalid collection")
	}

	// c, _ := bson.Marshal(data)

	docID, _ := primitive.ObjectIDFromHex(id)
	// var data bson.M

	result := collection.FindOne(dao.ctx, bson.M{"_id": docID})
	return result, nil
}

// update any collection based on id
func (dao *FactoryDAO) Update(key string, id string, obj interface{}) error {
	collection, ok := dao.Collections[key]
	if !ok {
		return errors.New("invalid collection")
	}
	docID, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.UpdateOne(dao.ctx, bson.M{"_id": docID}, bson.M{"$set": obj})
	return err
}

// get

// get by id

// soft  delete
