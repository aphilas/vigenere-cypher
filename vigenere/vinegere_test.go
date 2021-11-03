package vigenere_test

import (
	"testing"

	"github.com/nevilleomangi/vigenere-cypher/vigenere"
)

func TestVigenere(t *testing.T) {
	key := "ugali"

	cases := []struct {
		test, message, cipher string
	}{
		{"special characters", "lorem!#", "furpu!#"},
		{"lowercase ASCII", "lorem ipsum dolor sit amet", "furpu cvsfu xulzz mot luyz"},
		{"uppercase ASCII", "LOREM IPSUM DOLOR SIT AMET", "FURPU CVSFU XULZZ MOT LUYZ"},
	}

	for _, c := range cases {
		t.Run("Encrypts "+c.test, func(t *testing.T) {
			got := vigenere.Encrypt(key, c.message)
			if got != c.cipher {
				t.Fatalf("\nMessage:%q\n Got:%q\n Want:%q\n", c.message, got, c.cipher)
			}
		})
		t.Run("Decrypts "+c.test, func(t *testing.T) {
			got := vigenere.Decrypt(key, c.cipher)
			if got != c.message {
				t.Fatalf("\nCipher:%q\n Got:%q\n Want:%q\n", c.cipher, got, c.message)
			}
		})
	}
}
