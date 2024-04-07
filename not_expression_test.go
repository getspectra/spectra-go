package spectra

import "testing"

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
