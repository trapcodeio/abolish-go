package abolish

import (
	"testing"
)

// stringAble - check if variable is string-able
func Test_OptionToString(t *testing.T) {
	isString := optionToString("hello")
	if isString != "hello" {
		t.Errorf("expecting [%v] but got [%v]", "hello", isString)
	}

	isInt := optionToString(1)
	if isInt != "1" {
		t.Errorf("expecting [%v] but got [%v]", "1", isInt)
	}

	isFloat := optionToString(1.1)
	if isFloat != "1.1" {
		t.Errorf("expecting [%v] but got [%v]", "1.1", isFloat)
	}

	isBool := optionToString(true)
	if isBool != "true" {
		t.Errorf("expecting [%v] but got [%v]", "true", isBool)
	}

	isNil := optionToString(nil)
	if isNil != "<nil>" {
		t.Errorf("expecting [%v] but got [%v]", "<nil>", isNil)
	}

	isStruct := optionToString(struct{}{})
	if isStruct != "{}" {
		t.Errorf("expecting [%v] but got [%v]", "{}", isStruct)
	}

	isArray := optionToString([]string{"hello", "world"})
	if isArray != "[hello world]" {
		t.Errorf("expecting [%v] but got [%v]", "[hello world]", isArray)
	}
}
