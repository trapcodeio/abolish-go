package abolish

import (
	"errors"
	"fmt"
	"strings"
)

type Map map[string]interface{}
type Any interface{}

type ValidatorFunc func(value any, option *any) *ValidationError

type ValidationError struct {
	Validator string
	Code      string
	Message   string
}

func (e ValidationError) Error() string {
	return e.Message
}

var DefaultError = &ValidationError{
	Code: "__DEFAULT_ERROR__",
}

type Validator struct {
	Name        string
	Validate    ValidatorFunc
	Description string
	Error       *ValidationError
}

var validators = make(map[string]any)

func HasValidator(name string) bool {
	_, ok := validators[name]
	return ok
}

func RegisterValidator(v Validator) error {
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

	validators[v.Name] = v

	return nil
}

//goland:noinspection GoUnusedExportedFunction
func RegisterValidators(validators []Validator) error {
	for _, v := range validators {
		err := RegisterValidator(v)
		if err != nil {
			return err
		}
	}

	return nil
}

//goland:noinspection GoUnusedExportedFunction
func ReplaceValidator(name string, v Validator) error {
	// check if v already exists
	if !HasValidator(name) {
		return errors.New("validator does not exist")
	}

	// delete old validator
	delete(validators, name)

	// register new validator
	return RegisterValidator(v)
}

func Validate[T any](variable T, rules *Rules) error {

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
		validator, ok := validators[validatorName].(Validator)
		if !ok {
			// get real type of validator
			validatorType := fmt.Sprintf("%T", validators[validatorName])

			return &ValidationError{
				Code:    "validation",
				Message: fmt.Sprintf("validator [%v] is expecting [%v] but got [%T] instead", validatorName, validatorType, variable),
			}
		}

		// run validatorName
		err := validator.Validate(variable, &option)
		if err != nil {

			// if error is default error, set default error in the validator
			if err == DefaultError {
				err = validator.Error
			}

			// parse error message
			// replace :param with variable name
			// replace :option with option if string-able
			err.Message = strings.ReplaceAll(err.Message, paramPlaceholder, "Variable")
			err.Message = strings.ReplaceAll(err.Message, optionPlaceholder, optionToString(option))

			return err
		}
	}

	return nil
}
