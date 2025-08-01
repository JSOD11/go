package debug

import "fmt"

type Debugger struct {
	Active bool
}

// DebugLog
// Prints the formatted string to stdout if the debug flag is set.
func (d *Debugger) DebugLog(out string, args ...interface{}) {
	if d.Active {
		fmt.Printf(out, args...)
	}
}
