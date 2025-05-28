package gms

import (
	"fmt"
	"os"
)

// Pass Massage to Error and if the Application had to exit or not
func (g *GMS) Error(_message string, _exit bool) {
	if _exit {
		// If the application had to exit then print the error message and exit
		fmt.Printf("Error: %s", _message)
		os.Exit(0)
	} else {
		// If the application had to continue then just print the error message
		fmt.Printf("Warning: %s", _message)
	}
}
