package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func Salting(password string) string {
	salt := strconv.FormatInt(time.Now().UnixNano(), 10)
	saltedPassword := password + salt

	sha := sha256.Sum256([]byte(saltedPassword))

	return hex.EncodeToString(sha[:])
}
