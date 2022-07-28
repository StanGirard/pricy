package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/stangirard/pricy/internal/aws"
)

var (
	printVersion = flag.Bool("version", false, "print version and exit")
	sso          = flag.Bool("sso", false, "Use AWS SSO")
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

	var sess *session.Session
	if *sso {
		sess = aws.CreateSessionWithSSO()

	} else {
		sess = aws.CreateSessionWithCredentials()
	}
	aws.InitCostExplorer(sess)

	return nil
}
