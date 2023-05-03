package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"listener/event"
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
	log.Println("Listening for and consuming rabbitmq messages")

	// create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// listen and watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}
func connect() (*amqp.Connection, *amqp.Connection) {
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	// dont continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("Rabbit not yet ready")
			counts++
		} else {
			log.Println("Connected to Rabbit MQ")
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, nil
		}
		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backoff)
		continue

	}
	return connection, nil
}
