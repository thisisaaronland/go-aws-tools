package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-aws-tools/config"
	"github.com/go-ini/ini"
	"log"
	"os"
)

func main() {

	profile := flag.String("profile", "default", "A valid AWS credentials profile")

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

	for _, cred := range flag.Args() {

		k, err := sect.GetKey(cred)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(k.Value())
	}

	os.Exit(0)
}
