package main

import (
	"fmt"
	root "github/kubemq-io/json-streamer/cmd"
	"os"
)

func main() {
	if err := root.Execute(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
