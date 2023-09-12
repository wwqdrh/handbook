package cookbook

import (
	"errors"
	"fmt"
)

var (
	ErrorInternal = errors.New("internal error")
)

func getError(level int) error {
	level1Err := fmt.Errorf("[getData] level 1 error: %w", ErrorInternal)
	if level == 1 {
		return level1Err
	}
	if level == 2 {
		return fmt.Errorf("[getData] level 2 error: %w", level1Err)
	}

	return ErrorInternal
}

func ErrorSimple() {
	err := getError(1)
	if errors.Is(err, ErrorInternal) {
		fmt.Printf("is error internal: %v\n", err)
	}
	fmt.Printf("unwrapped error: %v\n", errors.Unwrap(err))

	fmt.Printf("---\n")

	err = getError(2)
	if errors.Is(err, ErrorInternal) {
		fmt.Printf("is error internal: %v\n", err)
	}
	unwrapped := errors.Unwrap(err)
	fmt.Printf("unwrapped error: %v\n", unwrapped)
	fmt.Printf("unwrapped unwrapped error: %v\n", errors.Unwrap(unwrapped))
}
