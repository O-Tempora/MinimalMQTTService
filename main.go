package main

import (
	"context"
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	ctx := context.Background()

	client := mqtt.NewClient(
		mqtt.NewClientOptions().
			AddBroker("mqtt://localhost:1883").
			SetClientID("proxy"),
	)
	if appToken := client.Connect(); appToken.Wait() && appToken.Error() != nil {
		log.Fatal("mqtt connection failed")
	}

	c := Client{
		client: client,
	}

	subToken := c.client.Subscribe("test", 1, c.EventHandler(ctx))
	if ok := subToken.Wait(); !ok || subToken.Error() != nil {
		log.Fatalf(
			"Failed to subscribe to a topic [%s]: wait - [%t], error - [%v]",
			"sound", ok, subToken.Error(),
		)
	}

	log.Println(http.ListenAndServe("localhost:7070", nil))
}
