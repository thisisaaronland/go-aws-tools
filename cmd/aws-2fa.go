package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"log"
	"strings"
)

const admDuration = 3600
const standardDuration = 36000

func main() {

	profile := flag.String("profile", "default", "...")

	flag.Parse()

	cfg, err := external.LoadDefaultAWSConfig(
		external.WithSharedConfigProfile(*profile),
	)

	if err != nil {
		log.Fatal(err)
	}

	iamc := iam.New(cfg)
	stsc := sts.New(cfg)

	username := username(stsc)
	mfaDevice := mfaDevice(iamc, username)

	tokenCode := tokenCode()

	credentials := sessionCredentials(stsc, mfaDevice, tokenCode, duration(*profile))

	log.Println(credentials)
}

func username(stsc *sts.STS) string {

	callerIdResp, err := stsc.GetCallerIdentityRequest(nil).Send()

	if err != nil {
		log.Fatal("Unable to get the username.")
	}

	arn := callerIdResp.Arn

	return strings.Split(*arn, ":user/")[1]
}

func mfaDevice(iamc *iam.IAM, userArn string) string {

	req := iamc.ListMFADevicesRequest(
		&iam.ListMFADevicesInput{UserName: &userArn},
	)

	mfaDevice, err := req.Send()

	if err != nil {
		log.Fatal("Unable to fetch the MFA device", err)
	}

	return *mfaDevice.MFADevices[0].SerialNumber
}

func sessionCredentials(stsc *sts.STS, mfaDevice string, tokenCode string, duration int64) *sts.Credentials {

	req := stsc.GetSessionTokenRequest(&sts.GetSessionTokenInput{
		SerialNumber:    &mfaDevice,
		DurationSeconds: &duration,
		TokenCode:       &tokenCode,
	})

	token, err := req.Send()

	if err != nil {
		log.Fatal("Unable to create a new session.", err)
	}

	return token.Credentials
}

func duration(profile string) int64 {

	if strings.HasSuffix(profile, "-adm") {
		return admDuration
	}

	return standardDuration
}

func tokenCode() string {
	var tokenCode string
	fmt.Print("Enter your token: ")
	fmt.Scanf("%s", &tokenCode)
	return tokenCode
}
