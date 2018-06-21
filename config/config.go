package config

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/go-ini/ini"
	"os"
)

type Config struct {
	Path string
}

func NewConfig() (*Config, error) {

	var ini_config *ini.File
	var ini_path string

	for _, path := range external.DefaultSharedConfigFiles {

		_, err := os.Stat(path)

		if err != nil {
			continue
		}

		i, err := ini.Load(path)

		if err != nil {
			return nil, err
		}

		ini_path = path
		ini_config = i
		break
	}

	if ini_config == nil {
		return nil, errors.New("Unable to load config file")
	}

	c := Config{
		Path: ini_path,
	}

	return &c, nil
}

func (c *Config) AWSConfigWithProfile(profile string) (aws.Config, error) {

	cfg := external.WithSharedConfigProfile(profile)
	return external.LoadDefaultAWSConfig(cfg)
}

func (c *Config) SetSessionCredentialsWithProfile(profile string, creds *sts.Credentials) error {

	ini_path := c.Path
	ini_config, err := ini.Load(ini_path)

	if err != nil {
		return err
	}

	sect := ini_config.Section(profile)

	sect.Key("aws_access_key_id").SetValue(*creds.AccessKeyId)
	sect.Key("aws_secret_access_key").SetValue(*creds.SecretAccessKey)
	sect.Key("aws_session_token").SetValue(*creds.SessionToken)

	err = ini_config.SaveTo(ini_path)

	if err != nil {
		return err
	}

	return nil
}
