package utils

import "testing"

func TestGenerateSecureNonce(t *testing.T) {

	str1 := GenerateSecureNonce(5)
	str2 := GenerateSecureNonce(5)

	if len(str1) != 5 {
		t.Fatalf("Str1 len is not correct")
	}

	if len(str2) != 5 {
		t.Fatalf("Str2 len is not correct")
	}

	if str1 == str2 {
		t.Fatalf("Both strings are the same")
	}
}
