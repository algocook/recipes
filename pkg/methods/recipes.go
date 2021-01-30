package methods

import (
	"context"

	"github.com/algocook/proto/recipes"
)

// RecipesMainServer comment
type RecipesMainServer struct{}

// GetRecipe gRPC function
func (s *RecipesMainServer) GetRecipe(c context.Context, request *recipes.GetRecipeRequest) (response *recipes.RecipeResponse, err error) {
	return
	mongoClient, err := NewMongoDBClient("mongodb://localhost:27017/test")
	if err != nil {
		return nil, err
	}
	defer mongoClient.client.Disconnect(c)

	result, err := mongoClient.GetRecipeByID(request.RecipeId.RecipeId)
	if err != nil {
		return nil, err
	}

	response.Recipe = &recipes.Recipe{
		Id: request.RecipeId,
		Info: &recipes.Recipe_Info{
			Desc: &recipes.Recipe_Info_Desc{
				Title:       result.recipeInfo.Title,
				Description: result.recipeInfo.Description,
			},
		},
	}

	return
}

// PostRecipe gRPC function
func (s *RecipesMainServer) PostRecipe(c context.Context, request *recipes.PostRecipeRequest) (response *recipes.RecipeResponse, err error) {

	response.Recipe.Id = &recipes.Recipe_ID{
		RecipeId: "string",
		OwnerId:  1,
	}

	return
}

// DeleteRecipe method
func (s *RecipesMainServer) DeleteRecipe(c context.Context, request *recipes.DeleteRecipeRequest) (response *recipes.Error, err error) {
	return
}
