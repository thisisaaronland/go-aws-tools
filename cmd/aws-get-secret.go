package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"encoding/base64"
	"log"
)

func main() {

	dsn := flag.String("dsn", "", "...")
	secret_name := flag.String("secret-name", "", "...")
	
	flag.Parse()
	
	ses := nil
	
	svc := secretsmanager.New(sess)

	for _, secret_name := range flag.Args() {


		secret_value, err := getSecret(svc, secret_name)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func getSecret(secret_name string) {
	
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secret_name),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)
	
	if err != nil {
		return "", err
	}

	var secretString, decodedBinarySecret string
	
	if result.SecretString != nil {
		
		secretString = *result.SecretString
		
	} else {
		
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)

		return "", err

		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
	}

	return secretString, nil
}
