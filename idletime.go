package idletime

import (
	"context"
	"time"
)

// SampleInterval is the default amount of time between
var SampleInterval = 15 * time.Second

// Action is a function that is executed with a context
type Action = func(context.Context) error

// RunWhileIdle runs the given action whenever the machine has been idle for the specified threshold amount of time.
// When the machine becomes active again, the action will be cancelled. Once the machine becomes idle again, the
// action will be triggered again.
func RunWhileIdle(ctx context.Context, threshold time.Duration, action Action) error {
	for {
		// Wait for the machine to become idle
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-UntilIdle(ctx, threshold):
		}

		// Run the action until the machine becomes active again
		if err := runUntilActive(ctx, action); err != nil {
			return err
		}
	}
}

// runUntilActive runs the given action while the machine is idle. If the machine becomes active, the
// context is cancelled.
func runUntilActive(ctx context.Context, action Action) error {
	// Create a cancellation context for the action
	ctx, cancel := NotifyContextWhenActive(ctx)
	defer cancel()

	// Run the action in the context
	return action(ctx)
}
