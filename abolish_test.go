package abolish

import (
	"fmt"
	"testing"
)

// A test validator
var testValidator = Validator{
	Name:  "required",
	Error: &ValidationError{Message: ":param is required"},
	Validate: func(value any, option any) *ValidationError {
		// check if value is nil
		if value == nil {
			return DefaultError
		}
		return nil
	},
}

func afterTest() {
	fmt.Println("=== cleanup after test ===")
	// remove validator
	removeAllValidators()
}

func Test_HasValidator(t *testing.T) {

	// check if validator exists
	if HasValidator(testValidator.Name) {
		t.Errorf("expecting [%v] but got [%v]", false, true)
	}

	// register validator
	err := RegisterValidator(testValidator)
	if err != nil {
		t.Errorf("expecting [%v] but got [%v]", nil, err)
	}

	// check if validator exists
	if !HasValidator(testValidator.Name) {
		t.Errorf("expecting [%v] but got [%v]", true, false)
	}

	t.Cleanup(afterTest)
}

func Test_RegisterValidator(t *testing.T) {

	t.Run("validator must have a name", func(t *testing.T) {
		// check that error is thrown if name is empty
		err := RegisterValidator(Validator{})
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("registered validator must have an error", func(t *testing.T) {
		err := RegisterValidator(Validator{Name: "test"})
		if err != nil {
			t.Fatal(err)
		}

		// check if validator exists
		if !HasValidator("test") {
			t.Error("expecting validator to exist")
		}

		// get validator
		v, err := GetRegisteredValidator("test")
		if err != nil {
			t.Fatal(err)
		}

		// check that error is set to default if nil
		if v.Error == nil {
			t.Fatalf("Error is nil")
		}
	})

	t.Cleanup(afterTest)
}

func Test_RegisterValidators(t *testing.T) {
	// register validators
	err := RegisterValidators([]Validator{
		{Name: "test1"},
		{Name: "test2"},
	})

	if err != nil {
		t.Fatal(err)
	}

	// check if validators exist
	if !HasValidator("test1") {
		t.Error("expecting validator to exist")
	}

	if !HasValidator("test2") {
		t.Error("expecting validator to exist")
	}

	t.Cleanup(afterTest)
}
