package main

import (
	"flag"
	_ "fmt"
	"github.com/aaronland/go-aws-tools/auth"
	"github.com/aaronland/go-aws-tools/config"
	"github.com/aaronland/go-aws-tools/utils"
	"log"
)

func main() {

	profile := flag.String("profile", "default", "...")
	session_profile := flag.String("session-profile", "session", "...")
	duration := flag.Int64("duration", 3600, "...")

	flag.Parse()

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	aws_cfg, err := cfg.AWSConfigWithProfile(*profile)

	if err != nil {
		log.Fatal(err)
	}

	code := utils.Readline("Enter your MFA token code:")

	if code == "" {
		log.Fatal("Invalid token code")
	}

	creds, err := auth.GetCredentialsWithMFA(aws_cfg, code, *duration)

	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetSessionCredentialsWithProfile(*session_profile, creds)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updated session credentials for '%s' profile (expires %s)\n", *session_profile, *creds.Expiration)
}
