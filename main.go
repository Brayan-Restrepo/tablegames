package main

import (
	"tablegames/config"
	"tablegames/databases"
	"tablegames/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
