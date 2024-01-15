package utils

import "fmt"

func ExtendError(root error, cause error) error {
	return fmt.Errorf("%w: err=%w", root, cause)
}

func AppendErrorInfo(root error, info string) error {
	return fmt.Errorf("%w: %s", root, info)
}
