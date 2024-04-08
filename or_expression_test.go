package spectra

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_OrExpression_Evaluate(t *testing.T) {
	true1 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	true2 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	false1 := &BinaryExpression{
		left:      "user.id",
		operation: neq,
		right:     1,
	}

	false2 := &BinaryExpression{
		left:      "user.id",
		operation: neq,
		right:     1,
	}

	data := Data{
		"user.id": 1,
	}

	or1 := &OrExpression{
		expressions: []Expression{true1, true2},
	}

	if !or1.Evaluate(data) {
		t.Error("OrExpression should be true")
	}

	or2 := &OrExpression{
		expressions: []Expression{true1, false1},
	}

	if !or2.Evaluate(data) {
		t.Error("OrExpression should be true")
	}

	or3 := &OrExpression{
		expressions: []Expression{false1, false2},
	}

	if or3.Evaluate(data) {
		t.Error("OrExpression should be false")
	}

	or4 := &OrExpression{
		expressions: []Expression{true1, true2, false1},
	}

	if !or4.Evaluate(data) {
		t.Error("OrExpression should be true")
	}
}

func Test_OrExpression_GetFields(t *testing.T) {
	true1 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	true2 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	or1 := &OrExpression{
		expressions: []Expression{true1, true2},
	}

	fields := or1.GetFields()

	fmt.Println(fields)

	if len(fields) != 1 {
		t.Errorf("Expected 1 field, got %d", len(fields))
	}

	if fields[0] != "user.id" {
		t.Errorf("Expected field 'user.id', got %s", fields[0])
	}
}

func Test_OrExpression_MarshalJSON(t *testing.T) {
	true1 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	true2 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	or1 := &OrExpression{
		expressions: []Expression{true1, true2},
	}

	expected := `{"or":[{"left":"user.id","operation":"=","right":1},{"left":"user.id","operation":"=","right":1}]}`
	jsonBytes, _ := json.Marshal(or1)
	if string(jsonBytes) != expected {
		t.Errorf("Expected %s, got %s", expected, string(jsonBytes))
	}
}
