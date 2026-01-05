package main

import (
	"database/sql"
	"log"

	"github.com/HadeedTariq/go-ecommerce-api/cmd/api"
	"github.com/HadeedTariq/go-ecommerce-api/config"
	"github.com/HadeedTariq/go-ecommerce-api/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	database, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(database)

	server := api.NewApiServer(":8080", database)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Db: connected successfully")
}
