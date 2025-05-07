package parsers

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParsingDatalessCommand(t *testing.T) {
	testString := "CONTAINING Point(50.123, 23.123) GET population, demographics, income"

	tokens, _ := NewTokenize(testString)

	if tokens[0] != "CONTAINING" {
		t.Fatal("Expected CONTAINING, got ", tokens[0])
	}

}

func TestParsingCommandWithData(t *testing.T) {
	testString := "INSERT Polygon(1,2,3,4) {\"json\": \"data\"}"
	tokens, _ := NewTokenize(testString)

	fmt.Printf("Tokens: %v\n", tokens)

	if tokens[0] != "INSERT" {
		t.Fatal("Expected INSERT, got ", tokens[0])
	}

	jsonData := tokens[3]
	var obj any
	err := json.Unmarshal([]byte(jsonData), &obj)

	if err != nil {
		t.Fatalf("Expected valid JSON, got error: %v", err)
	}

}
