package main

import (
	ab "abolish"
	"fmt"
)

func main() {
	addValidators()
	rule, err := ab.StringToRulesCompiled("required|exact:hello")
	if err != nil {
		panic(err)
	}

	err = ab.Validate[any](nil, &rule)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ab.Validate[string]("", &rule)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func addValidators() {
	// struct validator
	err := ab.RegisterValidator(ab.Validator{
		Name: "required",
		Error: &ab.ValidationError{
			Validator: "required",
			Code:      "required",
			Message:   "This field is required",
		},
		Validate: func(value any, option *any) *ab.ValidationError {
			// check if value is nil
			if value == nil {
				return ab.DefaultError
			}

			return nil
		},
	})

	// exact validator
	err = ab.RegisterValidator(ab.Validator{
		Name: "exact",
		Validate: func(value any, option *any) *ab.ValidationError {

			if value != *option {
				return &ab.ValidationError{
					Message: fmt.Sprintf("value must be exactly %v", *option),
				}
			}

			return nil
		},
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
