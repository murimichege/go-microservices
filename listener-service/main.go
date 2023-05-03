package main

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// start listening to the messages

	// create consumer

	// listen and watch the queue and consume events
}
func connect() (*amqp091.Connection, *amqp091.Connection) {
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp091.Connection

	// dont continue until rabbit is ready
	for {
		c, err := amqp091.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("Rabbit not yet ready")
			counts ++
		}
		else {
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, nil
		}
		backoff = time.Duration(math.Pow(float64((counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backoff)
		continue

	}
     return connection, nil
}
