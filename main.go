package main

import (
	"fmt"
	"mywebserver/web"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go web.StartWebServer()
	ending()
}

func ending() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP)
	select {
	case s := <-c:
		fmt.Println("get signal:", s)
		web.Stop()
	}
}
