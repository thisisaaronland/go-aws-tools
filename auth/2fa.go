package auth

func GetCredentialsWith2FA(cfg, token string) (*sts.Credentials, error) {

	iamc := iam.New(cfg)
	stsc := sts.New(cfg)

	username, err := username(stsc)

	if err != nil {
		return nil, err
	}

	mfaDevice := mfaDevice(iamc, username)

	if err != nil {
		return nil, err
	}

	return sessionCredentials(stsc, mfaDevice, token, duration(*profile))
}
