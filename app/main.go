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

func main() {
	var conf Config
	if _, err := toml.DecodeFile("configs.toml", &conf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)

	pqStr := "user=" + conf.DB.User + " password=" + conf.DB.Password + " dbname=remindbot sslmode=verify-full"
	db, err := sql.Open("postgres", pqStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// rows, err := db.Query("SELECT id FROM reminderes WHERE id = $1", 1)
	// fmt.Println(rows)

	ac := rb.NewAppContext(db)
	stack := alice.New()

	r := router.New()
	r.POST("/reminders", stack.ThenFunc(ac.CreateHandler))

	fmt.Println("Server starting at port 8080.")
	http.ListenAndServe(":8080", r)
}
