package commands

import (
	"encoding/json"
	"slices"
	"testing"

	"github.com/deuce109/mercurydb/v2/common/collections"
)

func TestParsingDatalessCommand(t *testing.T) {
	testString := "CONTAINS Point(50.123, 23.123) GET population, demographics, income"

	command, args, _ := ParseCommand(testString)

	t.Log(command)
	t.Log(args)

	argCollection := collections.NewCollection(args)

	if command != "CONTAINS" {
		t.Fatalf("Expected CONTAINS, got %v", command)
	}

	if argCollection.Count() != 5 {
		t.Fatalf("Expected args to have length 5 got %v", len(args))
	} else if !argCollection.ContainsAll("Point(50.123, 23.123)", "GET", "population,", "demographics,", "income") {
		t.Fatalf("Args did not contain expected values, got %v", args)
	}

}

func TestParsingCommandWithData(t *testing.T) {
	testString := "INSERT Polygon(1,2,3,4) test {\"json\": \"data\"} test2 {\"json\": \"data2\"}"
	command, args, _ := ParseCommand(testString)

	t.Log(args)

	if command != "INSERT" {
		t.Fatal("Expected INSERT, got ", command)
	}

	jsonData := args[2]
	t.Log(jsonData)
	var obj map[string]string
	err := json.Unmarshal([]byte(jsonData), &obj)

	if err != nil {
		t.Fatalf("Expected valid JSON, got error: %v", err)
	}

	if obj["json"] != "data" {
		t.Fatalf("Expected json key to have value 'data', got %s", obj["json"])
	}

	obj = nil

	jsonData = args[4]
	t.Log(jsonData)
	err = json.Unmarshal([]byte(jsonData), &obj)

	if err != nil {
		t.Fatalf("Expected valid JSON, got error: %v", err)
	}

	if obj["json"] != "data2" {
		t.Fatalf("Expected json key to have value 'data2', got %s", obj["json"])
	}
}

func TestParsingABadCommand(t *testing.T) {
	testString := "NOCOMMAND test"
	command, args, err := ParseCommand(testString)

	if command != "" || args != nil {
		t.Fatalf("Expected command and args to be unset got %s %v", command, args)
	}

	if err == nil {
		t.Fatal("Expected error to not be nil")
	}

	if err.Error() != "no command found in beginning of string" {
		t.Fatalf("Got incorrect error %s", err.Error())
	}
}

func TestNoArgCommand(t *testing.T) {
	testString := "PING"
	command, args, err := ParseCommand(testString)

	if err != nil {
		t.Fatalf("Expected error to be nil got %v", err)
	}

	if command != "PING" || !slices.Equal(args, []string{}) {
		t.Fatalf("Expected command to be PING and args to be empty got %s, %v", command, args)
	}

}
