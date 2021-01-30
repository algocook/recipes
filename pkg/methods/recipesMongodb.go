package methods

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IngredientJSON struct ...
type IngredientJSON struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// RecipeInfoJSON struct ...
type RecipeInfoJSON struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Ingredients []IngredientJSON `json:"ingredients"`
}

// RecipeJSON struct
type RecipeJSON struct {
	//ID         int64          `json:"_id"`
	recipeInfo RecipeInfoJSON `json:"recipe_info"`
}

// MongoDBClient ...
type MongoDBClient struct {
	client *mongo.Client
}

// NewMongoDBClient ...
func NewMongoDBClient(uri string) (*MongoDBClient, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	return &MongoDBClient{client}, nil
}

// GetRecipeByID ...
func (client *MongoDBClient) GetRecipeByID(id string) (RecipeJSON, error) {
	collection := client.client.Database("test").Collection("recipes")

	var result RecipeJSON

	var resultTmp RecipeJSON
	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&resultTmp)
	log.Output(1, resultTmp.recipeInfo.Title)

	filter := bson.M{}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return RecipeJSON{}, err
	}

	log.Output(1, fmt.Sprint(collection.CountDocuments(context.TODO(), bson.M{"_id": id})))
	//log.Output(1, fmt.Sprint(resultTmp.ID))
	log.Output(1, resultTmp.recipeInfo.Title)
	//log.Output(1, fmt.Sprint(result.ID))
	log.Output(1, result.recipeInfo.Title)
	log.Output(1, result.recipeInfo.Description)
	log.Output(1, fmt.Sprint(result.recipeInfo.Ingredients))

	return result, nil
}
