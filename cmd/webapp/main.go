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

	stack := alice.New()

	fmt.Println("test")
	r := router.New()
	r.POST("/reminders", stack.ThenFunc(ac.CommandHandler))

	http.ListenAndServe(":8080", r)
	fmt.Println("Server starting at port 8080.")
}

func initDB(datapath string) *sql.DB {
	db, err := sql.Open("sqlite3", datapath+"/reminders.db")
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
