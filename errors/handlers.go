package errors

import (
	"fmt"
	"log"
)

// HandleError takes care of error handling to reduce boilerplate code
func HandleError(err error) {
	e := fmt.Sprintf("\x1b[31m[ERROR]:\x1b[0m %v", err)
	log.Println(e)
}
