package main

import (
	"context"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	client mqtt.Client
}

func (c *Client) EventHandler(ctx context.Context) mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		if !c.IsConnected() {
			log.Println("Event handler: connection is not established")
			return
		}
		log.Printf("It's fine")
	}
}
