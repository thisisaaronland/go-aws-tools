package auth

// This is mostly code cribbed from here and refactored in to library code:
// https://github.com/DEEP-IMPACT-AG/skuld

// See also:
// https://docs.aws.amazon.com/cli/latest/reference/sts/get-session-token.html

import (
       "context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"strings"
)

func GetCredentialsWithMFA(cfg aws.Config, token string, duration int64) (*sts.Credentials, error) {

     ctx, cancel := context.WithCancel(context.Background())
     defer cancel()

     return GetCredentialsWithMFAWithContext(ctx, cfg, token, duration)
}

func GetCredentialsWithMFAWithContext(ctx context.Context, cfg aws.Config, token string, duration int64) (*sts.Credentials, error) {

	stsc := sts.New(cfg)

	username, err := username(ctx, stsc)

	if err != nil {
		return nil, err
	}

	iamc := iam.New(cfg)

	mfaDevice, err := mfaDevice(ctx, iamc, username)

	if err != nil {
		return nil, err
	}

	return sessionCredentials(ctx, stsc, mfaDevice, token, duration)
}

func username(ctx context.Context, stsc *sts.Client) (string, error) {

	req := stsc.GetCallerIdentityRequest(nil)

	callerIdResp, err := req.Send(ctx)

	if err != nil {
		return "", err
	}

	arn := callerIdResp.Arn

	return strings.Split(*arn, ":user/")[1], nil
}

func mfaDevice(ctx context.Context, iamc *iam.Client, userArn string) (string, error) {

	req := iamc.ListMFADevicesRequest(
		&iam.ListMFADevicesInput{UserName: &userArn},
	)

	mfaDevice, err := req.Send(ctx)

	if err != nil {
		return "", err
	}

	return *mfaDevice.MFADevices[0].SerialNumber, nil
}

func sessionCredentials(ctx context.Context, stsc *sts.Client, mfaDevice string, tokenCode string, duration int64) (*sts.Credentials, error) {

	req := stsc.GetSessionTokenRequest(&sts.GetSessionTokenInput{
		SerialNumber:    &mfaDevice,
		DurationSeconds: &duration,
		TokenCode:       &tokenCode,
	})

	token, err := req.Send(ctx)

	if err != nil {
		return nil, err
	}

	return token.Credentials, nil
}
