package spectra

import (
	"encoding/json"
	"testing"
)

func Test_BinaryExpression_Evaluate(t *testing.T) {
	data := Data{
		"foo": "foo",
	}

	b1 := &BinaryExpression{
		left:      "foo",
		operation: eq,
		right:     "foo",
	}

	if !b1.Evaluate(data) {
		t.Error("Expected true, got false")
	}

	b2 := &BinaryExpression{
		left:      "foo",
		operation: eq,
		right:     "bar",
	}

	if b2.Evaluate(data) {
		t.Error("Expected false, got true")
	}

	b3 := &BinaryExpression{
		left:      "foo",
		operation: neq,
		right:     "bar",
	}

	if !b3.Evaluate(data) {
		t.Error("Expected true, got false")
	}

	b4 := &BinaryExpression{
		left:      "foo",
		operation: neq,
		right:     "foo",
	}

	if b4.Evaluate(data) {
		t.Error("Expected false, got true")
	}
}

func Test_BinaryExpression_EvaluateWithRef(t *testing.T) {
	data := Data{
		"foo": "foo",
		"bar": "foo",
	}

	b1 := &BinaryExpression{
		left:      "foo",
		operation: eq,
		right:     &RefValue{ref: "bar"},
	}

	if !b1.Evaluate(data) {
		t.Error("Expected true, got false")
	}

	b2 := &BinaryExpression{
		left:      "foo",
		operation: eq,
		right:     "bar",
	}

	if b2.Evaluate(data) {
		t.Error("Expected false, got true")
	}
}

func Test_BinaryExpression_GetFields(t *testing.T) {
	b1 := &BinaryExpression{
		left:      "foo",
		operation: eq,
		right:     "foo",
	}

	fields := b1.GetFields()

	if fields[0] != "foo" {
		t.Errorf("Expected foo, got %v", fields[0])
	}

	b2 := &BinaryExpression{
		left:      "foo",
		operation: eq,
		right:     &RefValue{ref: "bar"},
	}

	fields = b2.GetFields()

	if fields[0] != "foo" {
		t.Errorf("Expected foo, got %v", fields[0])
	}

	if fields[1] != "bar" {
		t.Errorf("Expected bar, got %v", fields[1])
	}
}

func Test_BinaryExpression_MarshalJSON(t *testing.T) {
	b1 := &BinaryExpression{
		left:      "foo",
		operation: eq,
		right:     "foo",
	}

	expected := `{"left":"foo","operation":"=","right":"foo"}`
	jsonBytes, _ := json.Marshal(b1)
	if string(jsonBytes) != expected {
		t.Errorf("Expected %v, got %v", expected, string(jsonBytes))
	}
}
