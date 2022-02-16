package idletime

import (
	"syscall"
	"time"
	"unsafe"
)

var (
	user32           = syscall.MustLoadDLL("user32.dll")
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	getLastInputInfo = user32.MustFindProc("GetLastInputInfo")
	getTickCount     = kernel32.MustFindProc("GetTickCount")
	lastInputInfo    struct {
		cbSize uint32
		dwTime uint32
	}
)

// GetIdleTime gets the amount of time that the machine has been idle. This is the amount of time since a
// user input of any kind.
func GetIdleTime() (time.Duration, error) {
	lastInputInfo.cbSize = uint32(unsafe.Sizeof(lastInputInfo))
	currentTickCount, _, _ := getTickCount.Call()
	r1, _, err := getLastInputInfo.Call(uintptr(unsafe.Pointer(&lastInputInfo)))
	if r1 == 0 {
		return 0, err
	}
	duration := time.Duration((uint32(currentTickCount) - lastInputInfo.dwTime)) * time.Millisecond
	return duration, nil
}
