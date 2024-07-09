package main

import (
	"fmt"
	"log"
	"os"

	"goravel/bootstrap"

	"github.com/goravel/framework/facades"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
	}()

	// Debug: print environment variables
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_DATABASE:", os.Getenv("DB_DATABASE"))
	fmt.Println("DB_USERNAME:", os.Getenv("DB_USERNAME"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))

	sqlDB, err := facades.Orm().Connection("mysql").DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}

	// Ping the database
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("successfully connected to the database")

	select {}
}
