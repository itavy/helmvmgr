package errors

import "fmt"

func HError(code string, message string) error {
	return &herror{code, message}
}

type herror struct {
	code string
	message string
}

func (e *herror) Error() string {
	return fmt.Sprintf("HError: %s - %s", e.code, e.message)
}
