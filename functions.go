package abolish

import (
	"fmt"
)

var paramPlaceholder = ":param"
var optionPlaceholder = ":option"

// optionToString - check if variable is string-able
func optionToString(option any) string {
	return fmt.Sprintf("%v", option)
}
