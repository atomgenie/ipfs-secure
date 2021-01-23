package utils

import "crypto/sha256"

// HashSha256 Hash password
func HashSha256(password []byte) [32]byte {
	sha := sha256.New()
	sha.Write([]byte(password))

	var key [32]byte
	hash := sha.Sum(nil)

	copy(key[:], hash[:32])

	return key
}
