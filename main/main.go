package main

import (
	"fmt"
	"path"
	"time"

	"github.com/gardusig/clientfix"
	"github.com/gardusig/clientfix/internal/application"
)

func main() {
	fmt.Println("Starting client...")
	client, err := clientfix.NewClient(
		path.Join("config", "fix.cfg"),
		application.Application{},
	)
	if err != nil {
		panic(err)
	}
	client.Start()
	defer client.Stop()
	fmt.Println("Started client")
	time.Sleep(5 * time.Second)
}
