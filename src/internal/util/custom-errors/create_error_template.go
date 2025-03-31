package customerrors

import (
	"fmt"
)

// CreateErrorTemplate returns a closure with a configured topLevelErrorMessage.
// When the closure is executed with a relevant errorDetails input, a error object combining both messages will be returned.
// Pre-requisite: topLevelErrorMessage must have exactly one "%s".
func CreateErrorTemplate(topLevelErrorMessage string) func(errorDetails string) error {
	return func(errorDetails string) error {
		return fmt.Errorf(topLevelErrorMessage, errorDetails)
	}
}
