package config

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func LoadSharedConfigWithProfile(profile string) (aws.Config, error) {

	cfg := external.WithSharedConfigProfile(profile)
	return external.LoadDefaultAWSConfig(cfg)
}
