package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
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
	flag.Var(&paths, "path", "A path to invalidate (you can pass multiple -path flags for multiple paths)")

	dist := flag.String("distribution-id", "", "A valid CloudFront distribution ID")
	profile := flag.String("profile", "session", "A valid AWS credentials profile")

	flag.Parse()

	if *dist == "" {
		log.Fatal("Invalid distribution ID")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	aws_cfg, err := cfg.AWSConfigWithProfile(*profile)

	if err != nil {
		log.Fatal(err)
	}

	cf := cloudfront.New(aws_cfg)

	caller := fmt.Sprintf("%d#%s", *dist, paths.String())

	hash := sha1.Sum([]byte(caller))
	enc_caller := hex.EncodeToString(hash[:])

	quantity := int64(len(paths))

	batch_paths := cloudfront.Paths{
		Items:    paths,
		Quantity: &quantity, // I have no idea what this is supposed to be...
	}

	batch := cloudfront.InvalidationBatch{
		CallerReference: &enc_caller,
		Paths:           &batch_paths,
	}

	params := cloudfront.CreateInvalidationInput{
		DistributionId:    dist,
		InvalidationBatch: &batch,
	}

	req := cf.CreateInvalidationRequest(&params)

	rsp, err := req.Send(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp)
}
