package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

var (
	printVersion = flag.Bool("version", false, "print version and exit")
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//go:embed VERSION
var version string

func run() error {
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		return nil
	}
	aws.getCostUsage()

	return nil
}
