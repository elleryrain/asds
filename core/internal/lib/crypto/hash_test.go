package crypto_test

import (
	"testing"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
)

func TestPassword(t *testing.T) {
	password := "Password123"
	hash, err := crypto.Password(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hash == "" {
		t.Fatalf("expected hash to be generated, got empty string")
	}

	t.Logf("Generated hash: %s", hash)
}
