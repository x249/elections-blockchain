package errors_test

import (
	"fmt"
	"testing"

	"github.com/x249/elections-blockchain/errors"
)

func Test(t *testing.T) {
	t.Run("throw error and catch", func(t *testing.T) {
		err := fmt.Errorf("Error!")
		errors.HandleError(&err)
	})
}
