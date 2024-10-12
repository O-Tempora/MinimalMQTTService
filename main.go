package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
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

	conn, err := sql.Open("postgres", "postgresql://postgres@localhost:5432/test")
	if err != nil {
		log.Fatal(err.Error())
	}
	c.pgClient = conn

	subToken := c.client.Subscribe("test", 1, c.EventHandler(ctx))
	if ok := subToken.Wait(); !ok || subToken.Error() != nil {
		log.Fatalf(
			"Failed to subscribe to a topic [%s]: wait - [%t], error - [%v]",
			"sound", ok, subToken.Error(),
		)
	}

	log.Println(http.ListenAndServe("localhost:7070", nil))
}
