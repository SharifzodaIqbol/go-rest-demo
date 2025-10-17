package main

import (
	"my-gin-app/pkg/recipes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecipesHandler struct {
	store recipeStore
}
type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	List() (map[string]recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	Remove(name string) error
}

func (h RecipesHandler) CreateRecipe(c *gin.Context) {}
func (h RecipesHandler) ListRecipes(c *gin.Context)  {}
func (h RecipesHandler) GetRecipe(c *gin.Context)    {}
func (h RecipesHandler) UpdateRecipe(c *gin.Context) {}
func (h RecipesHandler) DeleteRecipe(c *gin.Context) {}
func NewRecipeHandler(s recipeStore) *RecipesHandler {
	return &RecipesHandler{
		store: s,
	}
}
func main() {
	router := gin.Default()
	router.GET("/", homePage)
	router.Run(":8082")
}
func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is my home page")
}
