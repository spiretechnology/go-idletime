package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/spiretechnology/go-idletime"
)

func main() {
	// Root context for the program
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	// Trigger `idleProcess` once the machine has been idle for 5 seconds
	if err := idletime.RunWhileIdle(ctx, 5*time.Second, idleProcess); err != nil {
		log.Println(err.Error())
	}
}

func idleProcess(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
		}
		log.Println("Machine is idle: ", time.Now())
	}
}
