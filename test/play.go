package main

import (
	ab "abolish"
	"abolish/validators"
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

	err = ab.Validate[any](20, &rule)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func addValidators() {
	err := ab.RegisterValidators([]ab.Validator{
		validators.Required,
		validators.Exact,
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
