package spectra

import (
	"testing"
)

func Test_ToOperation(t *testing.T) {
	t.Run("toOperation", func(t *testing.T) {
		if op, err := toOperation("="); err != nil || op != eq {
			t.Errorf("toOperation(\"=\") = %v, %v; want eq, nil", op, err)
		}

		if op, err := toOperation("!="); err != nil || op != neq {
			t.Errorf("toOperation(\"!=\") = %v, %v; want neq, nil", op, err)
		}

		if op, err := toOperation("!"); err == nil {
			t.Errorf("toOperation(\"!\") = %v, %v; want nil, error", op, err)
		}
	})
}

func Test_isValidOperation(t *testing.T) {
	t.Run("isValidOperation", func(t *testing.T) {
		if isValidOperation("=") == false {
			t.Errorf("isValidOperation(\"=\") = false; want true")
		}

		if isValidOperation("!=") == false {
			t.Errorf("isValidOperation(\"!=\") = false; want true")
		}

		if isValidOperation("!") == true {
			t.Errorf("isValidOperation(\"!\") = true; want false")
		}
	})
}

func TestOperation_operate(t *testing.T) {
	t.Run("operation operate", func(t *testing.T) {
		operation := eq
		if operation.operate(1, 1) == false {
			t.Errorf("operation.operate(1, 1) = false; want true")
		}

		if operation.operate(3.14, 3.14) == false {
			t.Errorf("operation.operate(3.14, 3.14) = false; want true")
		}

		if operation.operate("hello", "hello") == false {
			t.Errorf("operation.operate(\"hello\", \"hello\") = false; want true")
		}

		if operation.operate(1, 2) == true {
			t.Errorf("operation.operate(1, 2) = true; want false")
		}

		if operation.operate(3.14, 2.71) == true {
			t.Errorf("operation.operate(3.14, 2.71) = true; want false")
		}

		if operation.operate("hello", "world") == true {
			t.Errorf("operation.operate(\"hello\", \"world\") = true; want false")
		}

		operation = neq
		if operation.operate(1, 1) == true {
			t.Errorf("operation.operate(1, 1) = true; want false")
		}

		if operation.operate(3.14, 3.14) == true {
			t.Errorf("operation.operate(3.14, 3.14) = true; want false")
		}

		if operation.operate("hello", "hello") == true {
			t.Errorf("operation.operate(\"hello\", \"hello\") = true; want false")
		}

		if operation.operate(1, 2) == false {
			t.Errorf("operation.operate(1, 2) = false; want true")
		}

		if operation.operate(3.14, 2.71) == false {
			t.Errorf("operation.operate(3.14, 2.71) = false; want true")
		}

		if operation.operate("hello", "world") == false {
			t.Errorf("operation.operate(\"hello\", \"world\") = false; want true")
		}

		operation = gt
		if operation.operate(2, 1) == false {
			t.Errorf("operation.operate(2, 1) = false; want true")
		}

		if operation.operate(3.14, 2.71) == false {
			t.Errorf("operation.operate(3.14, 2.71) = false; want true")
		}

		if operation.operate(1, 2) == true {
			t.Errorf("operation.operate(1, 2) = true; want false")
		}

		if operation.operate(2.71, 3.14) == true {
			t.Errorf("operation.operate(2.71, 3.14) = true; want false")
		}

		operation = gte
		if operation.operate(2, 1) == false {
			t.Errorf("operation.operate(2, 1) = false; want true")
		}

		if operation.operate(3.14, 2.71) == false {
			t.Errorf("operation.operate(3.14, 2.71) = false; want true")
		}

		if operation.operate(1, 2) == true {
			t.Errorf("operation.operate(1, 2) = true; want false")
		}

		if operation.operate(2.71, 3.14) == true {
			t.Errorf("operation.operate(2.71, 3.14) = true; want false")
		}

		operator := lt
		if operator.operate(1, 2) == false {
			t.Errorf("operator.operate(1, 2) = false; want true")
		}

		if operator.operate(2.71, 3.14) == false {
			t.Errorf("operator.operate(2.71, 3.14) = false; want true")
		}

		if operator.operate(2, 1) == true {
			t.Errorf("operator.operate(2, 1) = true; want false")
		}

		if operator.operate(3.14, 2.71) == true {
			t.Errorf("operator.operate(3.14, 2.71) = true; want false")
		}

		operator = lte
		if operator.operate(1, 2) == false {
			t.Errorf("operator.operate(1, 2) = false; want true")
		}

		if operator.operate(2.71, 3.14) == false {
			t.Errorf("operator.operate(2.71, 3.14) = false; want true")
		}

		if operator.operate(2, 1) == true {
			t.Errorf("operator.operate(2, 1) = true; want false")
		}

		if operator.operate(3.14, 2.71) == true {
			t.Errorf("operator.operate(3.14, 2.71) = true; want false")
		}

		operation = in
		if operation.operate(1, []int{1, 2, 3}) == false {
			t.Errorf("operation.operate(1, []int{1, 2, 3}) = false; want true")
		}

		if operation.operate(1.1, []float64{1.1, 2.2, 3.3}) == false {
			t.Errorf("operation.operate(1.1, []float64{1.1, 2.2, 3.3}) = false; want true")
		}

		if operation.operate("hello", []string{"hello", "world"}) == false {
			t.Errorf("operation.operate(\"hello\", []string{\"hello\", \"world\"}) = false; want true")
		}

		if operation.operate(4, []int{1, 2, 3}) == true {
			t.Errorf("operation.operate(4, []int{1, 2, 3}) = true; want false")
		}

		if operation.operate(4.4, []float64{1.1, 2.2, 3.3}) == true {
			t.Errorf("operation.operate(4.4, []float64{1.1, 2.2, 3.3}) = true; want false")
		}

		if operation.operate("china", []string{"hello", "world"}) == true {
			t.Errorf("operation.operate(\"world\", []string{\"hello\", \"world\"}) = true; want false")
		}
	})
}
