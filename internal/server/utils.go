package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}

func findJSONLD(htmlFile io.Reader) (string, error) {
	var scriptContents string
	tokenizer := html.NewTokenizer(htmlFile)

	inScriptTag := false

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()

		switch tokenType {
		case html.ErrorToken:
			if tokenizer.Err() == io.EOF {
				return scriptContents, nil // end of the document
			}
			fmt.Printf("Error during tokenization: %v\n", tokenizer.Err())
			return scriptContents, tokenizer.Err() // return what we got so far.
		case html.StartTagToken:
			if token.Data == "script" {
				for _, attr := range token.Attr {
					if attr.Key == "type" && attr.Val == "application/ld+json" {
						inScriptTag = true
					}
				}
			}
		case html.TextToken:
			if inScriptTag {
				scriptContents = token.Data
			}
		case html.EndTagToken:
			if token.Data == "script" && inScriptTag {
				inScriptTag = false // existed the script tag
				return scriptContents, nil
			}
		}
	}
}

type SchemaGraph struct {
	Graph []interface{} `json:"@graph"`
}

type RecipeNutrition struct {
	Type                  string `json:"@type"`
	Calories              string `json:"calories"`
	CarbohydrateContent   string `json:"carbohydrateContent"`
	ProteinContent        string `json:"proteinContent"`
	FatContent            string `json:"fatContent"`
	SaturatedFatContent   string `json:"saturatedFatContent"`
	SodiumContent         string `json:"sodiumContent"`
	FiberContent          string `json:"fiberContent"`
	SugarContent          string `json:"sugarContent"`
	UnsaturatedFatContent string `json:"unsaturatedFatContent"`
	ServingSize           string `json:"servingSize"`
}

type Recipe struct {
	Type               string            `json:"@type"`
	Name               string            `json:"name"`
	Image              []string          `json:"image"`
	RecipeYield        []string          `json:"recipeYield"`
	PrepTime           string            `json:"prepTime"`
	CookTime           string            `json:"cookTime"`
	TotalTime          string            `json:"totalTime"`
	RecipeIngredient   []string          `json:"recipeIngredient"`
	RecipeNutrition    []RecipeNutrition `json:"recipeNutrition"`
	RecipeInstructions []struct {
		Text string `json:"text"`
	} `json:"recipeInstructions"`
}

func getRecipeJSON(payload string) (*Recipe, error) {
	var schemaGraph SchemaGraph
	err := json.Unmarshal([]byte(payload), &schemaGraph)
	if err != nil {
		return nil, fmt.Errorf("something went wrong with unmarshalling recipe SchemaGraph")
	}

	var foundRecipe *Recipe // use pointer so its nil if not found

	for _, item := range schemaGraph.Graph {
		// Convert item to map to check its type
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// Check if this item is a Recipe
		if itemType, exists := itemMap["@type"]; exists && itemType == "Recipe" {
			// Convert the item back to JSON bytes
			itemJSON, err := json.Marshal(item)
			if err != nil {
				return nil, fmt.Errorf("error marshaling recipe item: %w", err)
			}

			if err := json.Unmarshal(itemJSON, &foundRecipe); err != nil {
				return nil, fmt.Errorf("error unmarshaling recipe: %w", err)
			}

			break
		}
	}

	if foundRecipe == nil {
		return nil, fmt.Errorf("no recipe found in schema graph")
	}

	return foundRecipe, nil
}
