package spectra

import (
	"encoding/json"
	"testing"
)

func Test_NotExpression_Evaluate(t *testing.T) {
	true1 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	false1 := &BinaryExpression{
		left:      "user.id",
		operation: neq,
		right:     1,
	}

	data := Data{
		"user.id": 1,
	}

	not1 := &NotExpression{
		expression: true1,
	}

	if not1.Evaluate(data) {
		t.Errorf("Expected false, got true")
	}

	not2 := &NotExpression{
		expression: false1,
	}

	if !not2.Evaluate(data) {
		t.Errorf("Expected true, got false")
	}
}

func Test_NotExpression_GetFields(t *testing.T) {
	true1 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	not1 := &NotExpression{
		expression: true1,
	}

	fields := not1.GetFields()

	if len(fields) != 1 {
		t.Errorf("Expected 1 field, got %d", len(fields))
	}

	if fields[0] != "user.id" {
		t.Errorf("Expected field 'user.id', got %s", fields[0])
	}
}

func Test_NotExpression_MarshalJSON(t *testing.T) {
	true1 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	not1 := &NotExpression{
		expression: true1,
	}

	expected := `{"not":{"left":"user.id","operation":"=","right":1}}`
	jsonBytes, _ := json.Marshal(not1)
	if string(jsonBytes) != expected {
		t.Errorf("Expected %s, got %s", expected, string(jsonBytes))
	}
}
