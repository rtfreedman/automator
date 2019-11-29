package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rtfreedman/automator/watcher"
)

var ctx, cancel = context.WithCancel(context.Background())

func main() {
	fmt.Println("TODO: implement me")
	go watcher.Watch(ctx)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	cancel()
}
