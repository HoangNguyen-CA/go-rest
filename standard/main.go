package main

import (
	"log"
	"net/http"
	"regexp"

	"example.com/standard/pkg/recipes"
)

func main() {

	store := recipes.NewMemStore()
	recipesHandler := NewRecipesHandler(store)

	http.Handle("/", &HomeHandler{})
	http.Handle("/recipes", recipesHandler)
	http.Handle("/recipes/", recipesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type RecipesHandler struct {
	store recipeStore
}

func NewRecipesHandler(store recipeStore) *RecipesHandler {
	return &RecipesHandler{store: store}
}

func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var RecipeRe = regexp.MustCompile(`^/recipes/*$`)
	var RecipeReWithId = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)

	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipes(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithId.MatchString(r.URL.Path):
		h.GetRecipe(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithId.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithId.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	default:
		return
	}

}
func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {

}

func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {

}

type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	List() (map[string]recipes.Recipe, error)
	Remove(name string) error
}
