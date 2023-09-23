package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashingSederhana(s string) string {
	sha := sha256.New()
	sha.Write([]byte(s))
	encrypted := sha.Sum(nil)
	return fmt.Sprintf("%x", encrypted)
}

func HashingSederhanaV2(s string) string {
	rahasia := sha256.Sum256([]byte(s))
	return hex.EncodeToString(rahasia[:])
}
