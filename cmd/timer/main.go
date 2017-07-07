package main

import (
	"fmt"

	"github.com/aranair/remindbot/commands"
	"github.com/aranair/remindbot/config"
	"github.com/aranair/remindbot/handlers"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/BurntSushi/toml"
	"github.com/jasonlvhit/gocron"
)

func task(ac handlers.AppContext, chatId int64, text string) {
	ac.SendText(chatId, text)
}

func main() {
	var conf config.Config

	_, err := toml.DecodeFile("configs.toml", &conf)
	checkErr(err)

	fmt.Println(conf)
	db := initDB(conf.DB.Datapath)
	defer db.Close()

	ac := handlers.NewAppContext(db, conf, commands.NewCommandList())
	chatId := -6894201
	gocron.Every(10).Seconds().Do(task, ac, int64(chatId), "Timer test")
	<-gocron.Start()
}

func initDB(datapath string) *sql.DB {
	db, err := sql.Open("sqlite3", datapath+"/reminders.db")
	checkErr(err)

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
