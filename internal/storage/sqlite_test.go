package storage

import (
	"os"
	"testing"
)

func TestStorage(t *testing.T) {
	os.Remove("test.db")
	st, err := NewStorage("test.db")
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	err = st.CreateUser("testuser", "password")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	user, err := st.GetUserByLogin("testuser")
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	if user.Login != "testuser" {
		t.Fatalf("unexpected user: %v", user.Login)
	}

	err = st.SaveCalculation(user.ID, "2+2", 4)
	if err != nil {
		t.Fatalf("failed to save calculation: %v", err)
	}
}
