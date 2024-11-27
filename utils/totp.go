package utils

import "github.com/pquerna/otp/totp"

func GenerateTOTPSecret(account string) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MimicProxy",
		AccountName: account,
	})
	if err != nil {
		return "", "", err
	}
	return key.URL(), key.Secret(), nil
}

func VerifyTOTPCode(secret, code string) bool {
	return totp.Validate(code, secret)
}
