package validators

import ab "abolish"

var Required = ab.Validator{
	Name:  "required",
	Error: &ab.ValidationError{Message: ":param is required"},
	Validate: func(value any, option *any) *ab.ValidationError {
		// check if value is nil
		if value == nil {
			return ab.DefaultError
		}

		return nil
	},
}

var Exact = ab.Validator{
	Name:  "exact",
	Error: &ab.ValidationError{Message: ":param value must be exactly :option"},
	Validate: func(value any, option *any) *ab.ValidationError {

		if value != *option {
			return ab.DefaultError
		}

		return nil
	},
}
