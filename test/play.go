package main

import (
	abolish "abolish"
	"fmt"
)

var rule = abolish.StringToRules("struct:test")

type DataType struct {
	Name string `json:"name"`
}

func main() {
	addValidators()

	//var text = "testd"
	// print address of text
	data := DataType{
		Name: "My name is ewo!",
	}

	err := abolish.Validate[DataType](data, &rule)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//fmt.Println(abolish.StringToRules(`required|email:"hello"`))
	//fmt.Println(abolish.StringToRules(`required|email:'hello'`))
	//fmt.Println(abolish.StringToRules("required|email:`hello`"))
}

func addValidators() {
	// struct validator
	err := abolish.RegisterValidator(abolish.Validator[any]{
		Name: "struct",
		Validate: func(value any, option *any) *abolish.ValidationError {
			// convert value to DataType struct
			//data, ok := value.(DataType)
			//if !ok {
			//	return &abolish.ValidationError{
			//		Message: "value is not a struct",
			//	}
			//}

			fmt.Println("struct validator called")
			fmt.Println("value:", value)
			fmt.Println("option:", *option)
			return nil
		},
	})

	// exact validator
	err = abolish.RegisterValidator(abolish.Validator[string]{
		Name: "exact",
		Validate: func(value string, option *any) *abolish.ValidationError {

			if value != *option {
				return &abolish.ValidationError{
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
