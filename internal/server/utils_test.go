package server

import (
	"fmt"
	"os"
	"testing"
)

func TestFindJSONLD(t *testing.T) {
	file, err := os.Open("mocks/mockdata.html")
	if err != nil {
		t.Fatalf("Failed to open mock HTML file: %v", err)
	}
	defer file.Close()

	contents, err := findJSONLD(file)
	if err != nil {
		t.Fatal()
	}

	fmt.Println(contents)
}
