package main

import (
	"backendo-go/internal/config"
	"backendo-go/internal/database"
	"backendo-go/internal/service/event"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := initConfig()
	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)

	db, err := database.NewDatabase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := event.NewEventService(db, cfg)

	httpService := event.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func initConfig() *config.Config {
	configFile := flag.String("config", "./config/config.yaml", "this is the service config.")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS events (
		id integer primary key autoincrement,
		name varchar,
		start varchar,
		end varchar,
		description varchar);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panic on error
	// insertEvent := `INSERT INTO events (name, start, end, description) VALUES (?,?,?,?)`
	// s := fmt.Sprintf("Event number %v", time.Now().Nanosecond())
	// db.MustExec(insertEvent, s)
	return nil
}

// list = append(list, &Event{0, "event1", "20/6/2020", "20/6/2020", "sarasa"})
