package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	s "strings"
	"time"

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
	buf  *bytes.Buffer
}

type Reminder struct {
	Id      int64  `db:id`
	Content string `db:content`
	ChatId  int64  `db:chat_id`
}

func NewAppContext(db *sql.DB, conf config.Config, buf *bytes.Buffer) AppContext {
	return AppContext{db: db, conf: conf, buf: buf}
}

func (ac *AppContext) CommandHandler(w http.ResponseWriter, r *http.Request) {
	var update Update

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&update); err != nil {
		log.Println(err)
	} else {
		log.Println(update.Msg.Text)
	}

	arr := s.Split(update.Msg.Text, " ")
	cmd := s.Replace(arr[0], "/", "", -1)
	chatId := update.Msg.Chat.Id

	switch cmd {
	case "remindme":
		txt := s.Join(arr[1:len(arr)], " ")
		ac.save(txt, chatId)
	case "remind":
		txt := s.Join(arr[1:len(arr)], " ")
		ac.save(txt, chatId)
	case "list":
		ac.list(chatId)
	case "clear":
		id := s.Join(arr[1:len(arr)], " ")
		i, _ := strconv.Atoi(id)
		ac.clear(i, chatId)
	case "clearall":
		ac.clearall(chatId)
	default:
		fmt.Println("Ignoring update.")
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
	ac.sendText(chatId, "Done!")
}

func (ac *AppContext) clearall(chatId int64) {
	_, err := ac.db.Exec(`DELETE FROM reminders WHERE chat_id=$1`, chatId)
	checkErr(err)
	ac.sendText(chatId, "All reminders cleared!")
}

func (ac *AppContext) list(chatId int64) {
	rows, err := ac.db.Query(`SELECT content, id FROM reminders where chat_id=$1`, chatId)
	checkErr(err)
	defer rows.Close()

	var arr []string
	for rows.Next() {
		var r Reminder
		_ = rows.Scan(&r)
		strId := strconv.Itoa(int(r.Id))
		arr = append(arr, "- "+r.Content+"("+strId+")")
	}
	text := s.Join(arr, "\n")

	if len(text) < 5 {
		text = "No current reminders, hiak~"
	}

	ac.sendText(chatId, text)
}

func (ac *AppContext) sendText(chatId int64, text string) {
	link := "https://api.telegram.org/bot{botId}:{apiKey}/sendMessage?chat_id={chatId}&text={text}"
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
