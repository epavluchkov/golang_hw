package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "TCP connect timeout")
	flag.Parse()

	telnetClient := NewTelnetClient(net.JoinHostPort(flag.Arg(0), flag.Arg(1)), *timeout, os.Stdin, os.Stdout)
	defer telnetClient.Close()

	if err := telnetClient.Connect(); err != nil {
		fmt.Println("failed to connect: ", err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		telnetClient.Receive()
		cancel()
	}()

	go func() {
		telnetClient.Send()
		cancel()
	}()

	<-ctx.Done()
}
