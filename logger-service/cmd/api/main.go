package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"log-service/data"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

const (
	webPort  = 80
	rpcPort  = 5001
	mongoURL = "mongodb://localhost:27017"
	gRpcPort = 50001
)

var client *mongo.Client

func main() {
	// connect to mongo db
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient
	// create context for disconnect

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}

	}()
	app := Config{
		Models: data.New(client),
	}
	// start the web server
	//go app.serve()
	//Register the rpc server
	err = rpc.Register(new(RPCServer))
	go app.rpcListen()
	go app.gRPCListen()

}

type Config struct {
	Models data.Models
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic()
	}
}

func (app *Config) rpcListen() error {
	log.Println("Starting RPC server on Port", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()
	for {
		rpcCon, err := listen.Accept()
		if err != nil {
			continue

		}
		go rpc.ServeConn(rpcCon)
	}
}
func connectToMongo() (*mongo.Client, error) {
	//create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, error := mongo.Connect(context.TODO(), clientOptions)
	if error != nil {
		log.Println("Error Connecting", error)
		return nil, error

	}
	return c, nil

}
