package main

import (
	riak "github.com/basho/riak-go-client"
	"log"
)

type App struct {
	Routes     []Router
	riakClient *riak.Client
}

func (app *App) RiakClient() *riak.Client {
	if app.riakClient == nil {
		opts := &riak.NewClientOptions{
			RemoteAddresses: RIAK_ADDRESSES,
		}

		client, err := riak.NewClient(opts)
		if err != nil {
			log.Fatalln(err)
		}

		app.riakClient = client
	}

	return app.riakClient
}
