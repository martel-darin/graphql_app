package main

import (
	"context"
	"fmt"
	"log"

	"github.com/martel-darin/graphql_app/db"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Panic(err)
	}
	defer db.ConnectionPool.Close(context.Background())

	if users, err := db.FetchUsers(); err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("%+v\n", users)
	}
}
