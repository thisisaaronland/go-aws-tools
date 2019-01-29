package main

import (
	"flag"
	_ "fmt"
	"github.com/aaronland/go-aws-tools/auth"
	"github.com/aaronland/go-aws-tools/config"
	"github.com/aaronland/go-aws-tools/utils"
	"github.com/whosonfirst/iso8601duration"
	"log"
	"strings"
	"time"
)

func main() {

	profile := flag.String("profile", "default", "A valid AWS credentials profile")
	code := flag.String("code", "", "A valid MFA code. If empty the application will block and prompt the user")
	session_profile := flag.String("session-profile", "session", "The name of the AWS credentials profile to update with session credentials")
	session_duration := flag.String("duration", "PT1H", "A valid ISO8601 duration string indicating how long the session should last (months are currently not supported)")

	flag.Parse()

	d, err := duration.FromString(*session_duration)

	if err != nil {
		log.Fatal(err)
	}

	ttl_fl := d.ToDuration().Seconds()
	ttl := int64(ttl_fl)

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	aws_cfg, err := cfg.AWSConfigWithProfile(*profile)

	if err != nil {
		log.Fatal(err)
	}

	*code = strings.TrimSpace(*code)
	
	if *code == "" {

		*code = utils.Readline("Enter your MFA token code:")

		if *code == "" {
			log.Fatal("Invalid token code")
		}
	}

	creds, err := auth.GetCredentialsWithMFA(aws_cfg, *code, ttl)

	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetSessionCredentialsWithProfile(*session_profile, creds)

	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	then := now.Add(d.ToDuration())

	log.Printf("Updated session credentials for '%s' profile, expires %s (%s)\n", *session_profile, then.Format(time.Stamp), *creds.Expiration)
}
