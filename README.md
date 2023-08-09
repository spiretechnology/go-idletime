# go-idletime

This is a Go library for checking the system idle time. It also includes utilities for running scheduled tasks once the machine has become idle, as well as automatic context cancellation when the machine is suddenly no longer idle.

## Examples

## Check the idle time of the machine

```go
duration, err := idletime.GetIdleTime()
```

### Run a task when the machine is idle

```go
func main() {
	// Trigger `idleProcess` once the machine has been idle for 5 seconds
    ctx := context.Background()
	if err := idletime.RunWhileIdle(ctx, 5*time.Second, idleProcess); err != nil {
		log.Println(err.Error())
	}
}

func idleProcess(ctx context.Context) error {
    // Log "Your machine is idle" every second until the machine is no longer idle
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
		}
		log.Println("Your machine is idle: ", time.Now())
	}
}
```

## Platform support

It currently supports MacOS and Windows. Linux support would be easy to add, but there's been no need yet.
