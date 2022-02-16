package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/spiretechnology/go-idletime"
)

func main() {

	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	// Trigger `MyIdleAction` once the machine has been idle for 5 seconds
	if err := idletime.RunWhileIdle(ctx, 5*time.Second, MyIdleHttpServer); err != nil {
		log.Println(err.Error())
	}

}

// MyIdleAction is a simulated long-running action
func MyIdleAction(ctx context.Context) error {
	// Simulate some long action
	log.Println("Doing some work...")
	select {
	case <-ctx.Done():
		log.Println("System is no longer idle, or the parent context was cancelled")
		return ctx.Err()
	case <-time.After(10 * time.Second):
	}
	return nil
}

// MyIdleHttpServer is an HTTP server that only runs when the machine is idle. Once the machine is
// no longer idle, the server immediately shuts down.
func MyIdleHttpServer(ctx context.Context) error {

	log.Println("Starting HTTP server")
	server := http.Server{
		Addr: "127.0.0.1:4444",
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("Hello world"))
		}),
	}

	go func() {
		<-ctx.Done()
		server.Close()
	}()

	return server.ListenAndServe()

}
