package main

import (
	"strings"
)

const (
	a = 'a'
	z = 'z'
	A = 'A'
	Z = 'Z'
)

func LowerCaseASCII(char rune) bool {
	if a <= char && char <= z {
		return true
	}
	return false
}

func UpperCaseASCII(char rune) bool {
	if A <= char && char <= Z {
		return true
	}
	return false
}

// Clamps a value between min and max - assumes min > 0, max > 0
func Clamp(min, max, c rune) rune {
	return min + (c-min)%(max-min+1)
}

func AddASCII(char, addend rune) rune {
	if LowerCaseASCII(char) {
		return Clamp(a, z, char+(addend-a))
	} else if UpperCaseASCII(char) {
		return Clamp(A, Z, char+(addend-A))
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

func Decrypt(key, cipher string) string {
	var message strings.Builder
	i := 0

	for _, char := range cipher {
		k := key[i%(len(key))]
		if LowerCaseASCII(char) || UpperCaseASCII(char) {
			message.WriteRune(AddASCII(char, rune(-k)))
			i += 1
			continue
		}
		message.WriteRune(char)
	}

	return message.String()
}
