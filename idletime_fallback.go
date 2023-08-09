//go:build !(windows || darwin)

package idletime

import (
	"errors"
	"runtime"
	"time"
)

// GetIdleTime gets the amount of time that the machine has been idle. This is the amount of time since a
// user input of any kind.
func GetIdleTime() (time.Duration, error) {
	return time.Duration(0), errors.New("not implemented for this platform: %s-%s", runtime.GOOS, runtime.GOARCH)
}
