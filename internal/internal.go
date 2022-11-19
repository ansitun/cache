package internal

import "fmt"

// Private application and library code.
// This is the code you don't want others importing in their applications or libraries.
// Note that this layout pattern is enforced by the Go compiler itself.
func internal() {
	fmt.Println("This code is internal")
}
