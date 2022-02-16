package idletime

import (
	"context"
	"log"
	"time"
)

// DefaultSampleInterval is the default amount of time between
var DefaultSampleInterval = time.Second

// Action is a function that is executed with a context
type Action func(context.Context) error

// RunWhileIdle waits for the machine to become idle, then triggers the given action. If the machine becomes
// active (no longer idle) while the action function is being executed, the action function's context
// will automatically be cancelled.
func RunWhileIdle(ctx context.Context, threshold time.Duration, action Action) error {

	// Loop indefinitely until the context is cancelled
	for {

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(DefaultSampleInterval):
		}

		// Get the idle time
		idleTime, err := GetIdleTime()
		if err != nil {
			log.Println("idletime error: ", err.Error())
			continue
		}

		// If the idle time exceeds the threshold, trigger the action
		if idleTime >= threshold {
			if err := runUntilActive(ctx, DefaultSampleInterval, action); err != nil {
				log.Println("idletime action error: ", err.Error())
			}
		}

	}

}

// runUntilActive runs the given action while the machine is idle. If the machine becomes active, the
// context is cancelled.
func runUntilActive(ctx context.Context, sampleInterval time.Duration, action func(context.Context) error) error {

	// Create a cancellation context for the action
	ctx, cancel := NotifyContextWhenActive(ctx, sampleInterval)
	defer cancel()

	// Run the action in the context
	return action(ctx)

}
