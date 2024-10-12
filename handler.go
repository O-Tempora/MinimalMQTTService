package main

import (
	"context"
	"database/sql"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	client   mqtt.Client
	pgClient *sql.DB
}

func (c *Client) EventHandler(ctx context.Context) mqtt.MessageHandler {
	return func(cl mqtt.Client, m mqtt.Message) {
		if !cl.IsConnected() {
			log.Println("Event handler: connection is not established")
			return
		}
		log.Printf("It's fine")
		log.Print(c.pgClient.Ping())
	}
}
