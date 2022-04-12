//go:build !(widows || darwin)

package idletime

import (
	"log"
	"runtime"
	"time"
)

// GetIdleTime gets the amount of time that the machine has been idle. This is the amount of time since a
// user input of any kind.
func GetIdleTime() (time.Duration, error) {
	log.Printf("idletime.GetIdleTime() is not implemented for this platform: %s-%s\n", runtime.GOOS, runtime.GOARCH)
	return time.Duration(0)
}
