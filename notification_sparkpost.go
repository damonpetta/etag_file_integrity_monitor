package main

import (
	"log"
	"os"

	"github.com/SparkPost/gosparkpost"
)

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
