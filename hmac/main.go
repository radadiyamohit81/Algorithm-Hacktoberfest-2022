package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

// These are my helper library which I usually use when I need to check message integrity
// when data is passed between one system to another.
// It contains functions to generate the hash of a message and then to verify against given hash.
// Thinking of sharing it so that new developers don't have to figure out from scratch how to use the crypto libs for this

func ComputeHMACSHA256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func ComputeHMACSHA512(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func IsMatchedHMACSHA256(signature string, message string, secret string) bool {
	return ComputeHMACSHA256(message, secret) == signature
}

func IsMatchedHMACSHA512(signature string, message string, secret string) bool {
	return ComputeHMACSHA512(message, secret) == signature
}

// main function
func main() {
	secret := "this_is_super_secret_string"
	message := "this is super important message"
	expected := "wD+TzXu+kLtA9ybMy5YaQZIOyLwG+Gd7u41WGwcm9Z4=" // use the above at https://www.devglan.com/online-tools/hmac-sha256-online

	hmac256Result := ComputeHMACSHA256(message, secret)

	if hmac256Result == expected {
		fmt.Println("No one changed your important message")
	} else {
		fmt.Println("Oooh!!! someone is messing with you")
	}
}
