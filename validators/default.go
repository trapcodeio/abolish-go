package validators

import (
	ab "abolish"
)

var Required = ab.Validator{
	Name:  "required",
	Error: &ab.ValidationError{Message: ":param is required"},
	Validate: func(value any, option any) *ab.ValidationError {
		// check if value is nil
		if value == nil {
			return ab.DefaultError
		}

		return nil
	},
}

var Exact = ab.Validator{
	Name:       "exact",
	Error:      &ab.ValidationError{Message: ":param value must be exactly :option"},
	ValueTypes: &[]string{"string", "int", "bool"},
	Validate: func(value any, option any) *ab.ValidationError {
		// get type of value using reflect
		valueType := ab.TypeOf(value)

		// check if value type is not the same as option type
		if valueType != ab.TypeOf(option) {
			return ab.DefaultError
		}

		// if string
		if valueType == "string" {
			if value.(string) != option.(string) {
				return ab.DefaultError
			}
		} else if valueType == "int" {
			if value.(int) != option.(int) {
				return ab.DefaultError
			}
		} else if valueType == "bool" {
			if value.(bool) != option.(bool) {
				return ab.DefaultError
			}
		}

		return nil
	},
}
