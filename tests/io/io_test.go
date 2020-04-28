package io

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestPingPong(t *testing.T) {
	go server(t)

	// wait server startup
	time.Sleep(1 * time.Second)
	for i := 0; i < conc; i++ {
		go client(t)
	}

	log.SetFlags(log.LstdFlags | log.Llongfile)

	sg := make(chan os.Signal)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	<-sg
}
