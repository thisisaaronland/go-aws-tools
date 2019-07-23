package main

import (
	"flag"
	"github.com/aaronland/go-aws-tools/config"
	"github.com/go-ini/ini"
	"log"
	"os"
)

func main() {

	profile := flag.String("profile", "default", "A valid AWS credentials profile")
	token := flag.Bool("session-token", true, "Require AWS_SESSION_TOKEN")

	flag.Parse()

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	ini_config, err := ini.Load(cfg.Path)

	if err != nil {
		log.Fatal(err)
	}

	sect := ini_config.Section(*profile)

	creds := map[string]string{
		"aws_access_key_id":     "AWS_ACCESS_KEY_ID",
		"aws_secret_access_key": "AWS_SECRET_ACCESS_KEY",
	}

	if *token {
		creds["aws_session_token"] = "AWS_SESSION_TOKEN"
	}

	for ini_prop, env_var := range creds {

		k, err := sect.GetKey(ini_prop)

		if err != nil {
			log.Fatal(err)
		}

		err = os.Setenv(env_var, k.Value())

		if err != nil {
			log.Fatal(err)
		}
	}

}
