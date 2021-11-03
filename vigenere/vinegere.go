package vigenere

import (
	"strings"
)

const (
	a = 'a'
	z = 'z'
	A = 'A'
	Z = 'Z'
)

// Returns true if char is an ASCII char between a and z (lowercase alphabetical)
func LowerCaseASCII(char rune) bool {
	if a <= char && char <= z {
		return true
	}
	return false
}

// Returns true if char is an ASCII char between A and Z (uppercase alphabetical)
func UpperCaseASCII(char rune) bool {
	if A <= char && char <= Z {
		return true
	}
	return false
}

// Clamps a char code between min and max with wrap-around
func Clamp(min, max, c rune) rune {
	offset := (c - min) % (max - min + 1)
	if c < min {
		return max + 1 + offset
	}
	return min + offset
}

// Shifts an ASCII char by an offset
func AddASCII(char, offset rune) rune {
	if LowerCaseASCII(char) {
		return Clamp(a, z, char+offset)
	} else if UpperCaseASCII(char) {
		return Clamp(A, Z, char+offset)
	}

	return char
}

func SubtractASCII(char, offset rune) rune {
	if LowerCaseASCII(char) {
		return Clamp(a, z, char-offset)
	} else if UpperCaseASCII(char) {
		return Clamp(A, Z, char-offset)
	}

	return char
}

// Performs a Vinegere substitution
// of a string of ASCII characters between a-z or A-Z
// using the given key
func Encrypt(key, message string) string {
	var cipher strings.Builder
	i := 0

	for _, char := range message {
		k := key[i%(len(key))]

		// Use character offset for key character if alphabetical ASCII
		if LowerCaseASCII(rune(k)) {
			k = k - a
		} else if UpperCaseASCII(rune(k)) {
			k = k - A
		}

		if LowerCaseASCII(char) || UpperCaseASCII(char) {
			cipher.WriteRune(AddASCII(char, rune(k)))
			i += 1
			// Do not use up key if character is not a valid ASCII
			continue
		}
		cipher.WriteRune(char)
	}

	return cipher.String()
}

// Decrypts a Vinegere cypher
// of a string of ASCII characters between a-z or A-Z
// using the given key
func Decrypt(key, cipher string) string {
	var message strings.Builder
	i := 0

	for _, char := range cipher {
		k := key[i%(len(key))]

		// Use character offset for key character if alphabetical ASCII
		if LowerCaseASCII(rune(k)) {
			k = k - a
		} else if UpperCaseASCII(rune(k)) {
			k = k - A
		}

		if LowerCaseASCII(char) || UpperCaseASCII(char) {
			message.WriteRune(SubtractASCII(char, rune(k)))
			i += 1
			// Do not use up key if character is not a valid ASCII
			continue
		}
		message.WriteRune(char)
	}

	return message.String()
}
