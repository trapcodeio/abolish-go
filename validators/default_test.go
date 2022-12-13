package validators

import (
	ab "abolish"
	"testing"
)

func Test_Required(t *testing.T) {
	// Add required validator
	err := ab.RegisterValidator(Required)
	if err != nil {
		t.Fatal(err)
	}

	rule, err := ab.StringToRulesCompiled("required")
	if err != nil {
		t.Fatal(err)
	}

	// check nil
	err = ab.Validate[any](nil, &rule)
	if err == nil {
		t.Fatal("[nil] is expected to fail")
	}

	// check empty string
	err = ab.Validate[string]("", &rule)
	if err != nil {
		t.Fatal("[empty string] is expected to pass")
	}

	// check empty array
	err = ab.Validate[[]any]([]any{}, &rule)
	if err != nil {
		t.Fatal("[empty array] is expected to pass")
	}

	// empty map
	err = ab.Validate[map[any]any](map[any]any{}, &rule)
	if err != nil {
		t.Fatal("[empty map] is expected to pass")
	}

	// empty struct
	type testStruct struct {
		Name string
	}

	err = ab.Validate[testStruct](testStruct{}, &rule)
	if err != nil {
		t.Fatal("[empty struct] is expected to pass")
	}
}

func Test_Exact(t *testing.T) {
	// Add exact validator
	err := ab.RegisterValidator(Exact)
	if err != nil {
		t.Fatal(err)
	}

	// check string
	rule := ab.StringToRules("exact:hello")
	err = ab.Validate[string]("hello", &rule)
	if err != nil {
		t.Fatal(err)
	}

	// check number
	rule = ab.Rules{"exact": 10}
	err = ab.Validate[int](10, &rule)
	if err != nil {
		t.Fatal(err)
	}

	err = ab.Validate[int](20, &rule)
	if err == nil {
		t.Fatal("Expects validation to fail!")
	}

	// check boolean
	rule = ab.Rules{"exact": false}
	err = ab.Validate(false, &rule)
	if err != nil {
		t.Fatal(err)
	}

	err = ab.Validate(true, &rule)
	if err == nil {
		t.Fatal("Expects validation to fail!")
	}
}
