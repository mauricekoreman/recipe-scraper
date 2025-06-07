package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", s.recipeUrlHandler)

	// Wrap mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-ALlow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-ALlow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-ALlow-Credentials", "false") // Set to "true" if credentials are required

		// Handlre preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

func (s *Server) helloWorldHanlder(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello world"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) recipeUrlHandler(w http.ResponseWriter, r *http.Request) {
	type RecipeRequest struct {
		RecipeURL string `json:"recipeURL"`
	}

	decoder := json.NewDecoder(r.Body)
	req := RecipeRequest{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON", err)
		return
	}

	res, err := http.Get(req.RecipeURL)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong with fetching the recipe URL you provided", err)
		return
	}

	fmt.Println("Fetched recipe...")

	rawJSON, err := findJSONLD(res.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong parsing the data", err)
		return
	}

	recipeJSON, err := getRecipeJSON(rawJSON)

	// TODO: find a way to get the ingredient counts as seperate values instead of part of the string.

	fmt.Println("Done")
}
