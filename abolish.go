package abolish

import (
	"errors"
	"fmt"
	"strings"
)

// Map - abolish map
type Map map[string]interface{}

// Any - abolish any
type Any interface{}

// ValidatorFunc - abolish validator function
type ValidatorFunc func(value any, option *any) *ValidationError

// ValidationError - abolish validation error
type ValidationError struct {
	Validator string // validator name
	Code      string // error code
	Message   string // error message
}

// Error - return error message
func (e ValidationError) Error() string {
	return e.Message
}

// DefaultError - default error struct
// when returned in a validator, it will be replaced with the validator's error
var DefaultError = &ValidationError{
	Code: "__DEFAULT_ERROR__",
}

// Validator - abolish validator
type Validator struct {
	Name        string
	Validate    ValidatorFunc
	Description string
	Error       *ValidationError
}

// validatorsMap - abolish validatorsMap map
// key: validator name
// value: validator
var validatorsMap = make(map[string]Validator)

// HasValidator - check if validator a validator exists
func HasValidator(name string) bool {
	_, ok := validatorsMap[name]
	return ok
}

// RegisterValidator - register a validator
func RegisterValidator(v Validator) error {
	// check if name is empty
	if v.Name == "" {
		return errors.New("validator name cannot be empty")
	}

	// check if v already exists
	if HasValidator(v.Name) {
		return errors.New("validator already exists")
	}

	// if error is nil, set default error
	if v.Error == nil {
		v.Error = &ValidationError{
			Code:      "validation",
			Validator: v.Name,
			Message:   ":param failed [" + v.Name + "] validation.",
		}

	} else {
		// default error if none
		if v.Error.Code == "" {
			v.Error.Code = v.Name
		}

		// default validator if none
		if v.Error.Validator == "" {
			v.Error.Validator = v.Name
		}

		if v.Error.Message == "" {
			v.Error.Message = ":param failed [" + v.Name + "] validation."
		}
	}

	// register validator
	validatorsMap[v.Name] = v

	return nil
}

// RegisterValidators - register multiple validatorsMap
func RegisterValidators(validators []Validator) error {
	for _, v := range validators {
		err := RegisterValidator(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetRegisteredValidator - get a validator
func GetRegisteredValidator(name string) (Validator, error) {

	validator, ok := validatorsMap[name]
	if !ok {
		return Validator{}, errors.New("validator does not exist")
	}

	return validator, nil
}

// ReplaceValidator - replace a validator
func ReplaceValidator(name string, v Validator) error {
	// check if v already exists
	if !HasValidator(name) {
		return errors.New("validator does not exist")
	}

	// delete old validator
	delete(validatorsMap, name)

	// register new validator
	return RegisterValidator(v)
}

// removeAllValidators - remove all validatorsMap
func removeAllValidators() {
	validatorsMap = make(map[string]Validator)
}

// removeValidator - remove a validator
func removeValidator(name string) {
	delete(validatorsMap, name)
}

// Validate - validate a value
// value: value to validate
// rules: rules to validate value with
func Validate[T any](value T, rules *Rules) error {

	// loop through rules
	for validatorName, option := range *rules {
		// check if validatorName exists
		if !HasValidator(validatorName) {
			return &ValidationError{
				Code:    "validation",
				Message: fmt.Sprintf("validator [%v] does not exist.", validatorName),
			}
		}

		// get validator
		validator, ok := validatorsMap[validatorName]
		if !ok {
			// get real type of validator
			validatorType := fmt.Sprintf("%T", validatorsMap[validatorName])

			return &ValidationError{
				Code:    "validation",
				Message: fmt.Sprintf("validator [%v] is expecting [%v] but got [%T] instead", validatorName, validatorType, value),
			}
		}

		// run validatorName
		err := validator.Validate(value, &option)
		if err != nil {

			// if error is default error, set default error in the validator
			if err == DefaultError {
				err = validator.Error
			}

			// parse error message
			// replace :param with value name
			// replace :option with option if string-able
			err.Message = strings.ReplaceAll(err.Message, paramPlaceholder, "Variable")
			err.Message = strings.ReplaceAll(err.Message, optionPlaceholder, optionToString(option))

			return err
		}
	}

	return nil
}
