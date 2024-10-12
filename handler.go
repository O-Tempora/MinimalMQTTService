package main

import (
	"context"
	"database/sql"
	"log"
	"strconv"

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
		fl, err := strconv.ParseFloat(string(m.Payload()), 64)
		if err != nil {
			log.Printf("%v - parsed %v (%s as string)", err.Error(), m.Payload(), string(m.Payload()))
		}
		c.pgClient.Exec(`
				insert into gotest (time, data)
				values (now(), $1)
			`,
			fl,
		)
	}
}
