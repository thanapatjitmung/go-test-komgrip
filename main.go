package main

import (
	"context"
	"thanapatjitmung/go-test-komgrip/config"
	"thanapatjitmung/go-test-komgrip/databases"
	"thanapatjitmung/go-test-komgrip/server"
)

func main() {
	ctx := context.Background()
	conf := config.ConfigGetting()
	mariaDb := databases.NewMariaDatabase(conf.MariaDB)
	mongoDb := databases.NewMongoDatabase(ctx, conf.MongoDB)
	server := server.NewEchoServer(conf, mariaDb, mongoDb)
	server.Start()
}
