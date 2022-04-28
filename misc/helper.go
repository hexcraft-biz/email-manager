package misc

import (
	"crypto/rand"
	"crypto/sha512"
)

//================================================================
//
//================================================================
func StrInSlice(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}

	return false
}

//================================================================
//
//================================================================
func Sha512Sum(s string) []byte {
	hash := sha512.Sum512([]byte(s))
	sum := make([]byte, sha512.Size)
	for i := range hash {
		sum[i] = hash[i]
	}

	return sum
}

//================================================================
//
//================================================================
func GenRandomSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}

//================================================================
//
//================================================================
func ByteCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if (a[i] - b[i]) != 0 {
			return false
		}
	}

	return true
}
