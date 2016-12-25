package main

import (
	"fmt"
	"net/http"

	"github.com/aranair/remindbot/commands"
	"github.com/aranair/remindbot/config"
	"github.com/aranair/remindbot/handlers"

	router "github.com/aranair/remindbot/router"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/BurntSushi/toml"
	"github.com/justinas/alice"
)

func main() {
	var conf config.Config

	_, err := toml.DecodeFile("configs.toml", &conf)
	checkErr(err)

	fmt.Println(conf)
	db := initDB(conf.DB.Datapath)

	// pqStr := "user=" + conf.DB.User + " password='" + conf.DB.Password + "' dbname=remindbot host=localhost sslmode=disable"
	// fmt.Println(pqStr)

	// db, err := sql.Open("postgres", pqStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ac := handlers.NewAppContext(db, conf, commands.NewCommandList())
	stack := alice.New()

	r := router.New()
	r.POST("/reminders", stack.ThenFunc(ac.CommandHandler))

	http.ListenAndServe(":8080", r)
	fmt.Println("Server starting at port 8080.")
}

func initDB(datapath string) *sql.DB {
	db, err := sql.Open("sqlite3", datapath+"/reminders.db")
	checkErr(err)

	defer db.Close()

	err = createTable(db)
	checkErr(err)

	return db
}

// Create table if not exists
func createTable(db *sql.DB) (err error) {
	sql_table := `
	CREATE TABLE IF NOT EXISTS reminders(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		chat_id INTEGER,
		created DATETIME
	);
	`
	_, err = db.Exec(sql_table)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
