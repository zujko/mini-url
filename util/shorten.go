package util

import (
	"bytes"
	"math"
	"strings"
)

const alpha = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaLen = len(alpha)

func ShortenURL(url string) string {
	// Do Redis stuff and get ID
	var urlID = 126781289
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
	for i := 0; hashCount > i; i++ {
		hashBuf.WriteByte(alpha[hashDigits[i]])
	}

	return hashBuf.String()
}

// Decode decodes short url to id
func Decode(url string) int {
	result := 0
	shortUrlLen := len(url)
	for i := 0; i < shortUrlLen; i++ {
		x := strings.IndexByte(alpha, url[i])
		result += x * int(math.Pow(float64(alphaLen), float64(shortUrlLen-i-1)))
	}
	return result
}
