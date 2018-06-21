package main

import (
	"flag"
	_ "fmt"
	"github.com/aaronland/go-aws-tools/auth"
	"github.com/aaronland/go-aws-tools/config"
	"github.com/aaronland/go-aws-tools/utils"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/go-ini/ini"
	"log"
	"os"
)

// const admDuration = 3600
const standardDuration = 36000

func main() {

	profile := flag.String("profile", "default", "...")

	flag.Parse()

	cfg, err := config.LoadSharedConfigWithProfile(*profile)

	if err != nil {
		log.Fatal(err)
	}

	code := utils.Readline("Enter your MFA token code:")

	if code == "" {
		log.Fatal("Invalid token code")
	}

	creds, err := auth.GetCredentialsWithMFA(cfg, code, standardDuration)

	if err != nil {
		log.Fatal(err)
	}

	// sudo put me somewhere else..

	var ini_path string
	var ini_config *ini.File

	for _, path := range external.DefaultSharedConfigFiles {

		_, err := os.Stat(path)

		if err != nil {
			continue
		}

		i, err := ini.Load(path)

		if err != nil {
			log.Fatal(err)
		}

		ini_path = path
		ini_config = i
		break
	}

	if ini_config == nil {
		log.Fatal("Unable to load config file")
	}

	sect := ini_config.Section(*profile)

	sect.Key("aws_access_key_id").SetValue(*creds.AccessKeyId)
	sect.Key("aws_secret_access_key").SetValue(*creds.SecretAccessKey)
	sect.Key("aws_session_token").SetValue(*creds.SessionToken)

	err = ini_config.SaveTo(ini_path)

	if err != nil {
		log.Fatal(err)
	}

	// end of sudo put me somewhere else..

	log.Println(*creds.Expiration)
}
