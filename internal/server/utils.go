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

func findJSONLD(htmlFile io.Reader) ([]string, error) {
	var scriptContents []string
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
				inScriptTag = true
			}
		case html.TextToken:
			if inScriptTag {
				scriptContents = append(scriptContents, token.Data)
			}
		case html.EndTagToken:
			if token.Data == "script" {
				inScriptTag = false // existed the script tag
			}
		}
	}
}
