package util

import (
	"bytes"
	"log"
	"math"
	"net/url"
	"strings"

	"github.com/zujko/mini-url/db"
)

const alpha = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaLen = len(alpha)

// ShortenURL Checks if a long url already exists and returns the associated url
// otherwise, it will store the long url and return the short url.
func ShortenURL(url string) string {
	// Check if URL already exists
	exists, val := Exists(url)
	if exists {
		return val
	}
	shortURL := StoreURL(url)

	return shortURL
}

// Exists checks if the URL already exists. If it does, just use the already shortened URL
func Exists(longURL string) (bool, string) {
	urlObj, _ := url.Parse(longURL)
	if !urlObj.IsAbs() {
		longURL = "https://" + longURL
	}
	var shortURL string
	err := db.DBConn.QueryRow("SELECT short_url FROM url WHERE long_url = $1", longURL).Scan(&shortURL)
	if err != nil {
		return false, ""
	}
	return true, shortURL
}

// StoreURL store the long url and short url into postgres.
func StoreURL(longURL string) string {
	// Check if the URL is absolute or not
	url, _ := url.Parse(longURL)
	if !url.IsAbs() {
		longURL = "https://" + longURL
	}

	// Set short_url to an empty string because we need the id to encode
	var id int
	err := db.DBConn.QueryRow("INSERT INTO url(short_url,long_url) VALUES($1,$2) RETURNING url_id", "", longURL).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	shortURL := Encode(id)
	// Update the row to set the short_url
	result, err := db.DBConn.Exec("UPDATE url SET short_url = $1 WHERE url_id = $2", shortURL, id)
	if err != nil {
		log.Fatal(err)
	}
	// Check that the short_url was actually updated
	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal("Failed to update shorturl")
	}
	return shortURL
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
