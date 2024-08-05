package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OrignalURL   string    `json:"orignal_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

/*
we will use map DS to map one id with the link:
Say id = "9ef7ud" --> {
						ID: same
						OrignalURL = "https_link"
						ShortURL = "9ef7ud"
						CreationDate = time.Now()
						}
*/

var urlDB = make(map[string]URL)

// md5 is a hasher that uses checksum algo to hash string

func generateShortURL(OrignalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OrignalURL))
	fmt.Println("hasher: ", hasher)
	data := hasher.Sum(nil)
	fmt.Println(data)
	return "https://gautam.com"
}

func main() {
	fmt.Println("Starting url shortner!!")
	OrignalURL := "https://google.com"
	generateShortURL(OrignalURL)
}
