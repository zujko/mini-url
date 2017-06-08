package util

import (
	"bytes"
)

const alpha = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ShortenURL(url string) string {
	// Do Redis stuff and get ID
	var urlID = 1768
	return Encode(urlID)
}

// Encode encodes the URL's unique ID to Base62
func Encode(id int) string {
	hashDigits := []int{}
	var div = id
	var remainder = 0
	var hashBuf bytes.Buffer
	for div > 0 {
		remainder = div % 62
		div = div / 62
		hashDigits = append([]int{remainder}, hashDigits...)
	}

	var hashCount = len(hashDigits)
	i := 0
	for hashCount > i {
		hashBuf.WriteByte(alpha[hashDigits[i]])
		i++
	}

	return hashBuf.String()
}
