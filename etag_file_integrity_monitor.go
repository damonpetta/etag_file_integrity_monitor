package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/SparkPost/gosparkpost"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	baseURL := os.Getenv("FIM_URL")
	for {
		log.Printf("Requesting new hash for: [%s]", baseURL)
		baseHash := getEtag(baseURL)
		comapreHash(baseHash, baseURL)
	}
}

func comapreHash(baseHash string, baseURL string) {
	for {
		time.Sleep(6 * time.Second)
		currentHash := getEtag(baseURL)
		if baseHash != currentHash {
			log.Printf("Hash mismatch [%s] != [%s]", currentHash, baseHash)
			spNotify(currentHash)
			return
		}
	}
}

func getEtag(targetURL string) string {

	baseURL, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Add("cache", RandStringBytes(10))

	baseURL.RawQuery = params.Encode()
	log.Println(baseURL)

	response, err := http.Head(targetURL)
	if err != nil {
		log.Println("Error while checking", baseURL, ":", err)
	}

	hash := response.Header.Get("Etag")
	log.Printf("Got hash [%s] for [%s]", hash, baseURL)

	return hash
}

func spNotify(hash string) {

	apiKey := os.Getenv("FIM_SPARKPOST_API_KEY")
	notificationToEmail := os.Getenv("FIM_NOTIFICATION_TO_EMAIL")
	notificationFromEmail := os.Getenv("FIM_NOTIFICATION_FROM_EMAIL")
	cfg := &gosparkpost.Config{
		BaseUrl:    "https://api.sparkpost.com",
		ApiKey:     apiKey,
		ApiVersion: 1,
	}
	var client gosparkpost.Client
	err := client.Init(cfg)
	if err != nil {
		log.Fatalf("SparkPost client init failed: %s\n", err)
	}

	tx := &gosparkpost.Transmission{
		Recipients: []string{notificationToEmail},
		Content: gosparkpost.Content{
			HTML:    "<i>File hash changed in PCI Environment</i>",
			From:    notificationFromEmail,
			Subject: "[ALERT] PCI file monitoring",
		},
	}

	id, _, err := client.Send(tx)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Transmission sent with id [%s] to [%s]\n", id, notificationToEmail)

}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
