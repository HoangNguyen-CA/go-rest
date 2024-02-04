package main

import (
	"log"
	"net/http"
	"regexp"
)

func main() {

	http.Handle("/", &HomeHandler{})
	http.Handle("/recipes", &RecipesHandler{})
	http.Handle("/recipes/", &RecipesHandler{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type RecipesHandler struct {
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
