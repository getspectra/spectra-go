package spectra

import "testing"

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
