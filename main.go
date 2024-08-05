// This project is made by Gautam from IIT Bhubaneswar
// This projects shorten the proivded URL using hash func md5 that uses checksum algo

package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	hash := hex.EncodeToString(data)
	hash = hash[:8]
	fmt.Println(hash)
	fmt.Println(data)
	return hash
}

func CreateURL(OrignalURL string) string {
	hash := generateShortURL(OrignalURL)
	urlDB[hash] = URL{
		ID:           hash,
		OrignalURL:   OrignalURL,
		CreationDate: time.Now(),
		ShortURL:     hash,
	}
	return hash
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id] // accessing the value from map return two values, value, boolean, whether key is present or not
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func URLshortner(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	shortner := CreateURL(data.URL)
	// fmt.Fprint(w, shortner)
	w.Header().Set("Content-Type", "application/json")
	var response struct {
		ShortURL string `json:"url"`
	}
	response.ShortURL = shortner
	json.NewEncoder(w).Encode(response)
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusNotFound)
	}
	http.Redirect(w, r, url.OrignalURL, http.StatusFound)
}

func main() {
	fmt.Println("Starting url shortner!!")
	// OrignalURL := "https://google.com"
	// generateShortURL(OrignalURL)
	fmt.Println("Server starting on port 3000!!")
	http.HandleFunc("/", handler)
	http.HandleFunc("/shortner", URLshortner)
	http.HandleFunc("/redirect", redirectURL)
	// http.ListenAndServe :: return an error and takes two input, port number and handler func
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting server: ", err)
	}
}
