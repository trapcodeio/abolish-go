package abolish

import (
	"testing"
)

// A test validator
var testValidator = Validator{
	Name:  "required",
	Error: &ValidationError{Message: ":param is required"},
	Validate: func(value any, option *any) *ValidationError {
		// check if value is nil
		if value == nil {
			return DefaultError
		}
		return nil
	},
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
}

func Test_RegisterValidatorProcess(t *testing.T) {

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

}
