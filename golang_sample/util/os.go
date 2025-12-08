package util

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// HandleGracefulShutdown -
func HandleGracefulShutdown(fn func()) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	fn()
	log.Println("HandleGracefulShutdown successfully")
}
