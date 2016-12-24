package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/aranair/remindbot/config"
	"github.com/aranair/remindbot/handlers"
	router "github.com/aranair/remindbot/router"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/BurntSushi/toml"
	"github.com/justinas/alice"
)

type Reminder struct {
	Id      int64
	Content string
}

func main() {
	var conf config.Config

	_, err := toml.DecodeFile("configs.toml", &conf)
	checkErr(err)
	fmt.Println(conf)

	db, err := sql.Open("sqlite3", conf.DB.Datapath+"/reminders.db")
	checkErr(err)

	defer db.Close()
	CreateTable(db)

	// pqStr := "user=" + conf.DB.User + " password='" + conf.DB.Password + "' dbname=remindbot host=localhost sslmode=disable"
	// fmt.Println(pqStr)

	// db, err := sql.Open("postgres", pqStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	buf := bytes.NewBuffer(nil)
	ac := handlers.NewAppContext(db, conf, buf)
	stack := alice.New()

	r := router.New()
	r.POST("/reminders", stack.ThenFunc(ac.CommandHandler))

	http.ListenAndServe(":8080", r)
	fmt.Println("Server starting at port 8080.")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS reminders(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Content TEXT,
		Chat_id INTEGER,
		Created DATETIME
	);
	`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}
