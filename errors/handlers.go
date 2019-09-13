package errors

import (
	"fmt"
	"log"
)

// HandleError takes care of error handling to reduce boilerplate code
func HandleError(err *error) {
	if *err != nil {
		e := fmt.Sprintf("\x1b[31m[ERROR]: %v\x1b[0m", *err)
		log.Println(e)
	}
}
