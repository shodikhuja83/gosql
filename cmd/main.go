package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/shodikhuja83/gosql/cmd/app"
	"github.com/shodikhuja83/gosql/pkg/customers"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	host := "0.0.0.0"
	port := "9999"
	dbConnectionString := "postgres://app:pass@localhost:5432/db"
	if err := execute(host, port, dbConnectionString); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func execute(host, port, dbConnectionString string) (err error) {
	db, err := sql.Open("pgx", dbConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	mux := http.NewServeMux()
	customerService := customers.NewService(db)
	server := app.NewServer(mux, customerService)
	server.Init()

	httpServer := &http.Server{
		Addr:    host + ":" + port,
		Handler: server,
	}

	return httpServer.ListenAndServe()
}
