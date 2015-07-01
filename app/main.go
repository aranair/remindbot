package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	rb "github.com/aranair/remindbot"
	router "github.com/aranair/remindbot/router"

	_ "github.com/lib/pq"

	"github.com/BurntSushi/toml"
	"github.com/justinas/alice"
)

type Config struct {
	DB database `toml:"database"`
}

type database struct {
	User     string
	Password string
}

type Reminder struct {
	Id      int64
	Content string
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("configs.toml", &conf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)

	pqStr := "user=" + conf.DB.User + " password='" + conf.DB.Password + "' dbname=remindbot host=localhost sslmode=disable"
	fmt.Println(pqStr)
	db, err := sql.Open("postgres", pqStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ac := rb.NewAppContext(db)
	stack := alice.New()

	r := router.New()
	r.POST("/reminders", stack.ThenFunc(ac.CommandHandler))

	fmt.Println("Server starting at port 8080.")
	http.ListenAndServe(":8080", r)
}
