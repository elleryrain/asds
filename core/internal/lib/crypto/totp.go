package crypto

import (
	"encoding/base32"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var opts = totp.ValidateOpts{
	Period:    900, // 15 мин
	Skew:      1,
	Digits:    otp.DigitsEight,
	Algorithm: otp.AlgorithmSHA256,
}

// GenerateTOTPCode generates a TOTP code from a totp_salt and returns hex code
func GenerateTOTPCode(salt string) (string, error) {
	salt32 := base32.StdEncoding.EncodeToString([]byte(salt))
	code, err := totp.GenerateCodeCustom(salt32, time.Now(), opts)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString([]byte(code)), nil
}

// ValidateTOTP validates the TOTP code 
func ValidateTOTP(passcode string, salt string) error {
	salt32 := base32.StdEncoding.EncodeToString([]byte(salt))

	code, err := hex.DecodeString(passcode)
	if err != nil {
		return fmt.Errorf("hex decode: %v", err)
	}

	valid, err := totp.ValidateCustom(string(code), salt32, time.Now(), opts)
	if err != nil {
		return fmt.Errorf("validate: %v", err)
	}

	if !valid {
		return errors.New("invalid totp")
	}

	return nil
}
