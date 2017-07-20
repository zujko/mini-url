package util

import (
	"bytes"
	"math"
	"net/url"
	"strings"
)

const alpha = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaLen = len(alpha)

// ShortenURL Checks if a long url already exists and returns the associated url
// otherwise, it will encode the long url then store it in Redis.
func ShortenURL(url string) string {
	// Check if URL already exists
	exists, val := Exists(url)
	if exists {
		return val
	}

	/*redis, err := db.RedisPool.Get()
	defer db.RedisPool.Put(redis)
	if err != nil {
		log.Fatal(err)
	}*/

	// Grab unique ID
	/*resp, err := redis.Cmd("INCR", "url.id").Int()
	if err != nil {
		log.Fatal(err)
	}
	shortURL := Encode(resp)*/
	//StoreURL(shortURL, url)

	return "test"
}

// Exists checks if the URL already exists. If it does, just use the already shortened URL
func Exists(longURL string) (bool, string) {
	urlObj, _ := url.Parse(longURL)
	if !urlObj.IsAbs() {
		longURL = "https://" + longURL
	}

	//redis, err := db.RedisPool.Get()
	//defer db.RedisPool.Put(redis)

	//if err != nil {
	//	log.Fatal(err)
	//}
	//response, err := redis.Cmd("GET", longURL).Str()
	//if err != nil {
	//	return false, ""
	//}
	return true, "test"
}

// StoreURL store the long url and short url into Redis.
// This runs in a transaction so both URL's are saved at the same time.
func StoreURL(shortURL string, longURL string) {
	// Check if the URL is absolute or not
	url, _ := url.Parse(longURL)
	if !url.IsAbs() {
		longURL = "https://" + longURL
	}

	//redis, err := db.RedisPool.Get()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.RedisPool.Put(redis)

	// Start transaction
	//if redis.Cmd("MULTI").Err != nil {
	//	log.Fatal("Failed to create Transaction")
	//}
	/*// Save short url
	if redis.Cmd("SET", fmt.Sprintf("url:%s", shortURL), longURL).Err != nil {
		log.Fatal("Failed to set short URL")
	}
	// Save long url
	if redis.Cmd("SET", longURL, shortURL).Err != nil {
		log.Fatal("Failed to set long URL")
	}
	// Commit transaction
	if redis.Cmd("EXEC").Err != nil {
		log.Fatal("Failed to exec Transaction")
	}*/
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

/*
// Create db connection
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	db = os.Getenv("DB_NAME")
	sslmode = os.Getenv("DB_SSL")
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", host, port, user, db, sslmode)
*/
