package util

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"net/url"
	"strings"

	"github.com/zujko/mini-url/db"
)

const alpha = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaLen = len(alpha)

func ShortenURL(url string) string {
	// Check if URL already exists
	exists, val := Exists(url)
	if exists {
		return val
	}
	redis, err := db.RedisPool.Get()
	defer db.RedisPool.Put(redis)
	if err != nil {
		log.Fatal(err)
	}
	// Grab unique ID
	resp, err := redis.Cmd("INCR", "url.id").Int()
	if err != nil {
		log.Fatal(err)
	}
	shortURL := Encode(resp)
	StoreURL(shortURL, url)

	return shortURL
}

// Exists checks if the URL already exists
// If it does, just use the already shortened URL
func Exists(url string) (bool, string) {
	// Grab connection from pool
	redis, err := db.RedisPool.Get()
	defer db.RedisPool.Put(redis)

	if err != nil {
		log.Fatal(err)
	}
	// Check if it exists
	response := redis.Cmd("EXISTS", url)
	// Grab response as an int
	respVal, err := response.Int()
	if err != nil {
		log.Fatal(err)
	}
	if response.Err != nil {
		log.Fatal(response.Err)
	}
	// If it exists, return the shortened link
	if respVal == 1 {
		response, err := redis.Cmd("GET", fmt.Sprintf("%s", url)).Str()
		if err != nil {
			log.Fatal(err)
		}
		return true, response
	}
	return false, ""

}

func StoreURL(shortURL string, longURL string) {
	// Check if the URL is absolute or not
	url, _ := url.Parse(longURL)
	if !url.IsAbs() {
		longURL = "https://" + longURL
	}

	redis, err := db.RedisPool.Get()
	if err != nil {
		log.Fatal(err)
	}
	defer db.RedisPool.Put(redis)

	if redis.Cmd("MULTI").Err != nil {
		log.Fatal("Failed to create Transaction")
	}
	if redis.Cmd("SET", fmt.Sprintf("url:%s", shortURL), longURL).Err != nil {
		log.Fatal("Failed to set short URL")
	}
	if redis.Cmd("SET", longURL, shortURL).Err != nil {
		log.Fatal("Failed to set long URL")
	}
	if redis.Cmd("EXEC").Err != nil {
		log.Fatal("Failed to exec Transaction")
	}
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
