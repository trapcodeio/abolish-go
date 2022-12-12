package abolish

import (
	"errors"
	"fmt"
)

type Map map[string]interface{}
type Any interface{}

type ValidatorFunc[T any] func(value T, option *any) *ValidationError

type ValidationError struct {
	Validator string
	Code      string
	Message   string
}

func (e ValidationError) Error() string {
	return e.Message
}

type Validator[T any] struct {
	Name        string
	Validate    ValidatorFunc[T]
	Description string
	Error       ValidationError
}

var validators = make(map[string]any)

func RegisterValidator[T any](v Validator[T]) error {
	// check if v already exists
	if _, ok := validators[v.Name]; ok {
		return errors.New("validator already exists")
	}

	validators[v.Name] = v

	return nil
}

func RegisterValidators(validators []Validator[any]) error {
	for _, v := range validators {
		err := RegisterValidator(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReplaceValidator(name string, v Validator[any]) error {
	// check if v already exists
	if _, ok := validators[name]; !ok {
		return errors.New("validator does not exist")
	}

	validators[name] = v

	return nil
}

func HasValidator(name string) bool {
	_, ok := validators[name]
	return ok
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
		validator, ok := validators[validatorName].(Validator[any])
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
			err.Validator = validator.Name
			return err
		}
	}

	return nil
}
