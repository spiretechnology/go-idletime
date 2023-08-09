package idletime

import (
	"context"
	"log"
	"time"
)

// NotifyContextWhenActive cancels the given context when the machine is no longer idle. The sampleInterval provided
// determines how often the system idle time should be checked.
func NotifyContextWhenActive(ctx context.Context) (context.Context, context.CancelFunc) {
	// Create the cancellation context
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		// Wait until the system is active
		select {
		case <-ctx.Done():
		case <-UntilActive(ctx):
		}

		// Cancel the context
		cancel()
	}()

	// Return the context and our cancellation function
	return ctx, cancel
}

// UntilActive returns a channel that is closed when the machine becomes active.
func UntilActive(ctx context.Context) <-chan struct{} {
	// Create the channel to return
	chanActive := make(chan struct{})

	// In the background, poll the idle time until
	go func() {
		// The last idle time
		var lastIdleTime time.Duration

		for {
			// Wait for the sample interval
			select {
			case <-ctx.Done():
				return
			case <-time.After(SampleInterval):
			}

			// Check how long the machine has been idle now.
			idleTime, err := GetIdleTime()
			if err != nil {
				log.Printf("idletime: %s", err)
				continue
			}

			// If the idle time didn't track with the amount of time that has passed, then the system
			// became active during the last sample interval.
			if idleTime < lastIdleTime+SampleInterval/2 {
				close(chanActive)
				return
			}

			// Update the last idle time
			lastIdleTime = idleTime
		}
	}()

	// Return the channel
	return chanActive
}

// UntilIdle returns a channel that is closed when the machine becomes idle.
func UntilIdle(ctx context.Context, threshold time.Duration) <-chan struct{} {
	// Create the channel to return
	chanIdle := make(chan struct{})

	// In the background, poll the idle time until
	go func() {
		for {
			// Wait for the sample interval
			select {
			case <-ctx.Done():
				return
			case <-time.After(SampleInterval):
			}

			// Check how long the machine has been idle now.
			idleTime, err := GetIdleTime()
			if err != nil {
				log.Printf("idletime: %s", err)
				continue
			}

			// If the idle time exceeds the threshold, close the channel
			if idleTime > threshold {
				close(chanIdle)
				return
			}
		}
	}()

	// Return the channel
	return chanIdle
}
