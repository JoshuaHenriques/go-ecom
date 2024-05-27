package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash to be not empty")
	}

	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if err := ComparePassword(hash, "password"); err != nil {
		t.Errorf("expected password to match hash")
	}

	if err := ComparePassword(hash, "notpassword"); err == nil {
		t.Errorf("expected password to not match hash")
	}
}
