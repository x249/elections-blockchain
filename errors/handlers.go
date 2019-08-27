package errors

import (
	"fmt"
	"log"
)

// HandleError takes care of error handling to reduce boilerplate code
func HandleError(err error) {
	e := fmt.Sprintf("[ERROR]: %v", err)
	log.Println(e)
}
