package spectra

import (
	"encoding/json"
	"testing"
)

func Test_Policy_Apply(t *testing.T) {
	expression := &AndExpression{
		expressions: []Expression{
			&BinaryExpression{
				left:      "user.id",
				operation: eq,
				right:     1,
			},
			&BinaryExpression{
				left:      "team.id",
				operation: eq,
				right:     1,
			},
		},
	}

	data1 := Data{
		"user.id": 1,
		"team.id": 1,
	}

	data2 := Data{
		"user.id": 1,
		"team.id": 2,
	}

	data3 := Data{
		"user.id": 2,
		"team.id": 1,
	}

	p := &Policy{
		expression:  expression,
		effect:      Allow,
		permissions: []string{"EDIT_FILE"},
		description: "test policy",
	}

	if !p.Apply(data1) {
		t.Error("Expected true, got false")
	}

	if p.Apply(data2) {
		t.Error("Expected false, got true")
	}

	if p.Apply(data3) {
		t.Error("Expected false, got true")
	}
}

func Test_Policy_JsonSerialize(t *testing.T) {
	expression := &AndExpression{
		expressions: []Expression{
			&BinaryExpression{
				left:      "user.id",
				operation: eq,
				right:     1,
			},
			&BinaryExpression{
				left:      "team.id",
				operation: eq,
				right:     1,
			},
		},
	}

	p := &Policy{
		expression:  expression,
		effect:      Allow,
		permissions: []string{"EDIT_FILE"},
		description: "test policy",
	}

	expected := `{"description":"test policy","effect":"allow","filter":{"and":[{"left":"user.id","operation":"=","right":1},{"left":"team.id","operation":"=","right":1}]},"permissions":["EDIT_FILE"]}`
	jsonBytes, _ := json.Marshal(p)
	if string(jsonBytes) != expected {
		t.Errorf("Expected %s, got %s", expected, string(jsonBytes))
	}
}
