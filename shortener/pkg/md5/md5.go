package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Sum returns the md5 hash of the input data as a hex string
func Sum(data []byte) string {
	// 1. create a new md5 hash
	h := md5.New()
	// 2. write the data to the hash
	h.Write(data)
	// 3. return the hex string of the hash
	return hex.EncodeToString(h.Sum(nil))
}