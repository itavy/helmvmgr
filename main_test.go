package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
)

func TestMainProgram(t *testing.T) {
	// don't launch etcd server when invoked via go test
	if strings.HasSuffix(os.Args[0], "helmvmgr.test") {
		return
	}

	notifier := make(chan os.Signal, 1)
	signal.Notify(notifier, syscall.SIGINT, syscall.SIGTERM)
	go main()
	<-notifier
}
