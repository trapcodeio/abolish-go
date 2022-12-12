package main

import (
	ab "abolish"
	"fmt"
)

var rule = ab.StringToRules("required|exact:hello")

type DataType struct {
	Name string `json:"name"`
}

func main() {
	addValidators()

	//var text = "test"
	// print address of text
	//data := DataType{
	//	Name: "My name is ewo!",
	//}

	err := ab.Validate[string]("", &rule)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//fmt.Println(abolish.StringToRules(`required|email:"hello"`))
	//fmt.Println(abolish.StringToRules(`required|email:'hello'`))
	//fmt.Println(abolish.StringToRules("required|email:`hello`"))
}

func addValidators() {
	// struct validator
	err := ab.RegisterValidator(func() ab.Validator[any] {
		vErr := &ab.ValidationError{
			Validator: "required",
			Code:      "required",
			Message:   "This field is required",
		}

		return ab.Validator[any]{
			Name:  "required",
			Error: *vErr,
			Validate: func(value any, option *any) *ab.ValidationError {
				// check if value is nil
				if value == nil {
					return vErr
				}

				return nil
			},
		}
	}())

	// exact validator
	err = ab.RegisterValidator(ab.Validator[string]{
		Name: "exact",
		Validate: func(value string, option *any) *ab.ValidationError {

			if value != *option {
				return &ab.ValidationError{
					Message: fmt.Sprintf("value must be exactly %v", *option),
				}
			}

			return nil
		},
	})

	if err != nil {
		return
	}
}
