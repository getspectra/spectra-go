package spectra

import "testing"

func Test_Factory_And(t *testing.T) {
	and1 := And([]Expression{
		Eq("user.id", 1),
		Eq("team.id", 1),
	})

	data := Data{
		"user.id": 1,
		"team.id": 1,
	}

	if !and1.Evaluate(data) {
		t.Error("and1 should be true")
	}
}

func Test_Factory_Or(t *testing.T) {
	or1 := Or([]Expression{
		Eq("user.id", 1),
		Eq("team.id", 1),
	})

	data := Data{
		"user.id": 1,
		"team.id": 2,
	}

	if !or1.Evaluate(data) {
		t.Error("or1 should be true")
	}
}

func Test_Factory_Not(t *testing.T) {
	not1 := Not(Eq("user.id", 1))

	data := Data{
		"user.id": 2,
	}

	if !not1.Evaluate(data) {
		t.Error("not1 should be true")
	}
}

func Test_Factory_Eq(t *testing.T) {
	eq1 := Eq("user.id", 1)

	data := Data{
		"user.id": 1,
	}

	if !eq1.Evaluate(data) {
		t.Error("eq1 should be true")
	}
}

func Test_Factory_Neq(t *testing.T) {
	neq1 := Neq("user.id", 1)

	data := Data{
		"user.id": 2,
	}

	if !neq1.Evaluate(data) {
		t.Error("ne1 should be true")
	}
}

func Test_Factory_Neq2(t *testing.T) {
	neq21 := Neq2("user.id", 1)

	data := Data{
		"user.id": 2,
	}

	if !neq21.Evaluate(data) {
		t.Error("neq21 should be true")
	}
}

func Test_Factory_Gt(t *testing.T) {
	gt1 := Gt("user.id", 1)

	data := Data{
		"user.id": 2,
	}

	if !gt1.Evaluate(data) {
		t.Error("gt1 should be true")
	}
}

func Test_Factory_Gte(t *testing.T) {
	gte1 := Gte("user.id", 1)

	data := Data{
		"user.id": 1,
	}

	if !gte1.Evaluate(data) {
		t.Error("gte1 should be true")
	}
}

func Test_Factory_Lt(t *testing.T) {
	lt1 := Lt("user.id", 2)

	data := Data{
		"user.id": 1,
	}

	if !lt1.Evaluate(data) {
		t.Error("lt1 should be true")
	}
}

func Test_Factory_Lte(t *testing.T) {
	lte1 := Lte("user.id", 1)

	data := Data{
		"user.id": 1,
	}

	if !lte1.Evaluate(data) {
		t.Error("lte1 should be true")
	}
}

func Test_Factory_In(t *testing.T) {
	in1 := In("user.id", []int{1, 2})

	data := Data{
		"user.id": 1,
	}

	if !in1.Evaluate(data) {
		t.Error("in1 should be true")
	}
}

func Test_Factory_Nin(t *testing.T) {
	nin1 := Nin("user.id", []int{1, 2})

	data := Data{
		"user.id": 3,
	}

	if !nin1.Evaluate(data) {
		t.Error("nin1 should be true")
	}
}

func Test_Factory_NotIn(t *testing.T) {
	notIn1 := NotIn("user.id", []int{1, 2})

	data := Data{
		"user.id": 3,
	}

	if !notIn1.Evaluate(data) {
		t.Error("notIn1 should be true")
	}
}

func Test_Factory_NormalizeExpression(t *testing.T) {
	data := Data{
		"user.id": 1,
		"team.id": 1,
	}

	and1 := NormalizeExpression(map[string]any{
		"and": []any{
			[]any{"user.id", "=", 1},
			[]any{"team.id", "=", 1},
		},
	})

	if !and1.Evaluate(data) {
		t.Error("and1 should be true")
	}

	or1 := NormalizeExpression(map[string]any{
		"or": []any{
			[]any{"user.id", "=", 1},
			[]any{"team.id", "=", 10},
		},
	})

	if !or1.Evaluate(data) {
		t.Error("or1 should be true")
	}

	not1 := NormalizeExpression(map[string]any{
		"not": []any{"user.id", "=", 2},
	})

	if !not1.Evaluate(data) {
		t.Error("not1 should be true")
	}
}
