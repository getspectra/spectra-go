package spectra

import "testing"

func Test_isValidExpressionTuple(t *testing.T) {
	t.Run("isValidExpressionTuple", func(t *testing.T) {
		if isValidExpressionTuple([]any{"user.is_admin", "=", true}) == false {
			t.Errorf("isValidExpressionTuple([\"user.is_admin\", \"=\", true]) = false; want true")
		}

		if isValidExpressionTuple([]any{"user.is_admin", "!", "true"}) == true {
			t.Errorf("isValidExpressionTuple([\"user.is_admin\", \"!\", \"true\"]) = true; want false")
		}
	})
}
