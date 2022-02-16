package idletime

import "time"

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation -framework AppKit

#include <CoreFoundation/CoreFoundation.h>
#include <AppKit/AppKit.h>

double desktop_idle_get_time() {
	return CGEventSourceSecondsSinceLastEventType(kCGEventSourceStateHIDSystemState, kCGAnyInputEventType);
}
*/
import "C"

// GetIdleTime gets the amount of time that the machine has been idle. This is the amount of time since a
// user input of any kind.
func GetIdleTime() (time.Duration, error) {
	secondsFloat := C.desktop_idle_get_time()
	duration := time.Duration(secondsFloat) * time.Second
	return duration, nil
}
