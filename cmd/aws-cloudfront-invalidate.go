package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-aws-tools/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"log"
	"strings"
)

type Paths []string

func (e *Paths) String() string {
	return strings.Join(*e, "\n")
}

func (e *Paths) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func main() {

	var paths Paths
	flag.Var(&paths, "path", "One or paths to invalidate")

	dist := flag.String("distribution-id", "", "A valid CloudFront distribution ID")
	profile := flag.String("profile", "default", "...")

	flag.Parse()

	if *dist == "" {
		log.Fatal("Invalid distribution ID")
	}

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	aws_cfg, err := cfg.AWSConfigWithProfile(*profile)

	if err != nil {
		log.Fatal(err)
	}

	cf := cloudfront.New(aws_cfg)

	caller := "FIX ME"

	quantity := int64(len(paths))

	batch_paths := cloudfront.Paths{
		Items:    paths,
		Quantity: &quantity, // I have no idea what this is supposed to be...
	}

	batch := cloudfront.InvalidationBatch{
		CallerReference: &caller,
		Paths:           &batch_paths,
	}

	params := cloudfront.CreateInvalidationInput{
		DistributionId:    dist,
		InvalidationBatch: &batch,
	}

	req := cf.CreateInvalidationRequest(&params)

	rsp, err := req.Send()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp)
}
