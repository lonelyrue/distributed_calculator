package auth

import (
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := GenerateJWT(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	claims, err := ParseJWT(token)
	if err != nil {
		t.Fatalf("failed to parse token: %v", err)
	}

	if claims.UserID != 1 {
		t.Fatalf("unexpected userID: got %v", claims.UserID)
	}
}
