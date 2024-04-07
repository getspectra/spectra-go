package spectra

import "testing"

func Test_Spectra_Validate(t *testing.T) {
	loader := func(fields []FieldName) (Data, error) {
		return Data{
			"user.id": 1,
		}, nil
	}

	allowPolicy := NewPolicy(Eq("user.id", 1), Allow, []string{"EDIT_FILE"}, "Allow edit file for user 1")
	denyPolicy := NewPolicy(Eq("user.id", 1), Deny, []string{"EDIT_FILE"}, "Deny edit file for user 1")
	isValid := Validate([]*Policy{allowPolicy, denyPolicy}, DataLoaderFunc(loader), "EDIT_FILE")
	if isValid {
		t.Error("Expected false, got true")
	}

	allowPolicy = NewPolicy(Eq("user.id", 1), Allow, []string{"EDIT_FILE"}, "Allow edit file for user 1")
	denyPolicy = NewPolicy(Eq("user.id", 2), Deny, []string{"EDIT_FILE"}, "Deny edit file for user 2")
	isValid = Validate([]*Policy{allowPolicy, denyPolicy}, DataLoaderFunc(loader), "EDIT_FILE")
	if !isValid {
		t.Error("Expected true, got false")
	}
}
