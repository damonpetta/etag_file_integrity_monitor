package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {

	errc := make(chan error)

	// HTTP ping to keep alive on Heroku
	go func() {
		p := getPort()
		log.Printf("Listening on %s...\n", p)
		http.HandleFunc("/ping", httpPing)
		errc <- http.ListenAndServe(p, nil)
	}()

	// Begin our monitoring
	monitor()

}

func monitor() {
	baseURL := os.Getenv("FIM_URL")
	for {
		log.Printf("Requesting new hash for: [%s]", baseURL)
		baseHash := getEtag(baseURL)
		comapreHash(baseHash, baseURL)
	}
}

func comapreHash(baseHash string, baseURL string) {
	for {
		time.Sleep(60 * time.Second)
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
