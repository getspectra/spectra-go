package main

import (
	"fmt"
	. "github.com/getspectra/spectra"
)

func main() {
	loader := func(fields []FieldName) (Data, error) {
		return Data{
			"user.id": 1,
		}, nil
	}

	allowPolicy := NewPolicy(Eq("user.id", 1), Allow, []string{"EDIT_FILE"}, "Allow edit file for user 1")
	denyPolicy := NewPolicy(Eq("user.id", 1), Deny, []string{"EDIT_FILE"}, "Deny edit file for user 1")
	isValid := Validate([]*Policy{allowPolicy, denyPolicy}, DataLoaderFunc(loader), "EDIT_FILE")

	fmt.Println("Is valid:", isValid) // Output: Is valid: false

	allowPolicy = NewPolicy(Eq("user.id", 1), Allow, []string{"EDIT_FILE"}, "Allow edit file for user 1")
	denyPolicy = NewPolicy(Eq("user.id", 2), Deny, []string{"EDIT_FILE"}, "Deny edit file for user 2")
	isValid = Validate([]*Policy{allowPolicy, denyPolicy}, DataLoaderFunc(loader), "EDIT_FILE")

	fmt.Println("Is valid:", isValid) // Output: Is valid: true
}
