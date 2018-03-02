package snippets

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash just hashes given string
func MD5Hash(in string) string {
	hasher := md5.New()
	hasher.Write([]byte(in))
	return hex.EncodeToString(hasher.Sum(nil))
}
