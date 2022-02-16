package idletime

import (
	"context"
	"log"
	"time"
)

// NotifyContextWhenActive cancels the given context when the machine is no longer idle. The sampleInterval provided
// determines how often the system idle time should be checked.
func NotifyContextWhenActive(ctx context.Context, sampleInterval time.Duration) (context.Context, context.CancelFunc) {

	// Create the cancellation context
	ctx, cancel := context.WithCancel(ctx)

	// In the background, poll the idle time until
	go func() {

		// The last idle time
		lastIdleTime := time.Duration(0)

		for {

			// Allow context cancellation to breakout of the sleep interval
			select {
			case <-ctx.Done():
				cancel()
				return
			case <-time.After(sampleInterval):
			}

			// Get the idle time
			idleTime, err := GetIdleTime()
			if err != nil {
				log.Println("idletime error: ", err.Error())
				return
			}

			// If the idle time is unexpectedly low, cancel the context
			if idleTime < lastIdleTime+sampleInterval/2 {
				cancel()
				return
			}

			// Update the last idle time
			lastIdleTime = idleTime

		}

	}()

	// Return the context and our cancellation function
	return ctx, cancel

}
