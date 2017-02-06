package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	s "strings"
	"time"

	"github.com/aranair/remindbot/commands"
	"github.com/aranair/remindbot/config"
)

type Update struct {
	Id  int64   `json:"update_id"`
	Msg Message `json:"message"`
}

type Message struct {
	Id   int64  `json:"message_id"`
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type AppContext struct {
	db   *sql.DB
	conf config.Config
	cmds commands.Commands
}

type Reminder struct {
	Id      int64     `sql:id`
	Content string    `sql:content`
	Created time.Time `sql:created`
	ChatId  int64     `sql:chat_id`
}

func NewAppContext(db *sql.DB, conf config.Config, cmds commands.Commands) AppContext {
	return AppContext{db: db, conf: conf, cmds: cmds}
}

func (ac *AppContext) CommandHandler(w http.ResponseWriter, r *http.Request) {
	var update Update

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&update); err != nil {
		log.Println(err)
	} else {
		log.Println(update.Msg.Text)
	}

	cmd, txt := ac.cmds.Extract(update.Msg.Text)
	chatId := update.Msg.Chat.Id

	switch s.ToLower(cmd) {
	case "hazel":
		ac.sendText(chatId, "안녕~")
	case "remind":
		ac.save(txt, chatId)
	case "list":
		ac.list(chatId)
	case "renum":
		ac.renum(chatId)
	case "clear":
		i, _ := strconv.Atoi(txt)
		ac.clear(i, chatId)
	case "clearall":
		ac.clearall(chatId)
	}
}

func (ac *AppContext) save(txt string, chatId int64) {
	_, err := ac.db.Exec(`INSERT INTO reminders(content, created, chat_id) VALUES ($1, $2, $3)`, txt, time.Now(), chatId)
	checkErr(err)
	ac.sendText(chatId, "I remember liao!")
}

func (ac *AppContext) clear(id int, chatId int64) {
	_, err := ac.db.Exec(`DELETE FROM reminders WHERE chat_id=$1 AND id=$2`, chatId, id)
	checkErr(err)
	// "&#127881;"
	ac.sendText(chatId, "Pew!")
}

func (ac *AppContext) clearall(chatId int64) {
	_, err := ac.db.Exec(`DELETE FROM reminders WHERE chat_id=$1`, chatId)
	checkErr(err)
	ac.sendText(chatId, "Pew Pew Pew!")
}

func (ac *AppContext) list(chatId int64) {
	rows, err := ac.db.Query(`SELECT id, content, created FROM reminders WHERE chat_id=$1`, chatId)
	checkErr(err)
	defer rows.Close()

	var arr []string

	for rows.Next() {
		var c string
		var i int64
		var d time.Time
		_ = rows.Scan(&i, &c, &d)

		label = timeSinceLabel(d)
		line := "- " + c + label + "`" + strconv.Itoa(int(i)) + "`"
		arr = append(arr, line)
	}
	text := s.Join(arr, "\n")

	if len(text) < 5 {
		text = "No current reminders, hiak~"
	}

	ac.sendText(chatId, text)
}

func timeSinceLabel(d time.Time) string {
	var duration = time.Since(d)
	var durationNum int
	var unit string

	if int(duration.Hours()) == 0 {
		durationNum = int(duration.Minutes())
		unit = "min"
	} else if duration.Hours() < 24 {
		durationNum = int(duration.Hours())
		unit = "hour"
	} else {
		durationNum = int(duration.Hours()) / 24
		unit = "day"
	}

	if durationNum > 1 {
		unit = unit + "s"
	}

	return " `" + strconv.Itoa(int(durationNum)) + " " + unit + "`"
}

// This resets numbers for everyone!
func (ac *AppContext) renum(chatId int64) {
	rows, err := ac.db.Query(`SELECT content, created, chat_id FROM reminders`)
	checkErr(err)
	defer rows.Close()

	var arr []Reminder
	var c string
	var d time.Time
	var cid int64

	for rows.Next() {
		_ = rows.Scan(&c, &d, &cid)
		arr = append(arr, Reminder{Content: c, Created: d, ChatId: cid})
	}

	_, err = ac.db.Exec(`DELETE FROM reminders`)
	checkErr(err)

	_, err = ac.db.Exec(`DELETE FROM sqlite_sequence WHERE name='reminders';`)
	checkErr(err)

	for _, r := range arr {
		_, err := ac.db.Exec(`INSERT INTO reminders(content, created, chat_id) VALUES ($1, $2, $3)`, r.Content, r.Created, r.ChatId)
		checkErr(err)
	}

	ac.list(chatId)
}

func (ac *AppContext) sendText(chatId int64, text string) {
	link := "https://api.telegram.org/bot{botId}:{apiKey}/sendMessage?chat_id={chatId}&text={text}&parse_mode=Markdown"
	link = s.Replace(link, "{botId}", ac.conf.BOT.BotId, -1)
	link = s.Replace(link, "{apiKey}", ac.conf.BOT.ApiKey, -1)
	link = s.Replace(link, "{chatId}", strconv.FormatInt(chatId, 10), -1)
	link = s.Replace(link, "{text}", url.QueryEscape(text), -1)

	fmt.Println(link)

	_, _ = http.Get(link)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
