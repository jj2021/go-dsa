package dsa

import "testing"

func TestVerify(t *testing.T) {
	// Generate Key Pair
	pair := GenerateKeyPair()

	// Sign Content
	content := []byte("abc")
	sig := Sign(content, pair.Private, pair.Params)

	// Verify Content
	valid, err := Verify(sig, content, pair.Public, pair.Params)
	if err != nil || valid == false {
		t.Errorf("Message signature verification failed")
	}

	// Verify modified content
	modifiedContent := []byte("bca")
	valid, err = Verify(sig, modifiedContent, pair.Public, pair.Params)
	if err == nil || valid == true {
		t.Errorf("Expected verification to fail; Got valid signature")
	}
}
