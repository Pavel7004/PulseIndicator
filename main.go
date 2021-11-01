package main

import (
	"github.com/lawl/pulseaudio"

	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	client, err := pulseaudio.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-sig:
			return
		case <-tick.C:
			vol, err := client.Volume()
			if err != nil {
				panic(err)
			}
			fmt.Println(vol)
		}
	}
}
