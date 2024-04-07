package spectra

import "testing"

func Test_AndExpression_Evaluate(t *testing.T) {
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

	and1 := &AndExpression{
		expressions: []Expression{true1, true2},
	}

	if !and1.Evaluate(data) {
		t.Error("AndExpression should return true")
	}

	and2 := &AndExpression{
		expressions: []Expression{true1, false1},
	}

	if and2.Evaluate(data) {
		t.Error("AndExpression should return false")
	}

	and3 := &AndExpression{
		expressions: []Expression{false1, true1},
	}

	if and3.Evaluate(data) {
		t.Error("AndExpression should return false")
	}

	and4 := &AndExpression{
		expressions: []Expression{false1, false2},
	}

	if and4.Evaluate(data) {
		t.Error("AndExpression should return false")
	}

	and5 := &AndExpression{
		expressions: []Expression{true1, true2, false1},
	}

	if and5.Evaluate(data) {
		t.Error("AndExpression should return false")
	}
}

func Test_AndExpression_GetFields(t *testing.T) {
	expression1 := &BinaryExpression{
		left:      "user.id",
		operation: eq,
		right:     1,
	}

	expression2 := &BinaryExpression{
		left:      "user.name",
		operation: eq,
		right:     "ramzeng",
	}

	expression3 := &BinaryExpression{
		left:      "user.age",
		operation: eq,
		right:     28,
	}

	and1 := &AndExpression{
		expressions: []Expression{expression1},
	}

	if len(and1.GetFields()) != 1 && and1.GetFields()[0] != "user.id" {
		t.Error("AndExpression should return user.id")
	}

	and2 := &AndExpression{
		expressions: []Expression{expression1, expression2},
	}

	if len(and2.GetFields()) != 2 && and2.GetFields()[0] != "user.id" && and2.GetFields()[1] != "user.name" {
		t.Error("AndExpression should return user.id and user.name")
	}

	and3 := &AndExpression{
		expressions: []Expression{expression1, expression2, expression3},
	}

	if len(and3.GetFields()) != 3 && and3.GetFields()[0] != "user.id" && and3.GetFields()[1] != "user.name" && and3.GetFields()[2] != "user.age" {
		t.Error("AndExpression should return user.id, user.name and user.age")
	}
}
