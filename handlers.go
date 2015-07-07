package remindbot

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	s "strings"

	"github.com/aranair/remindbot/config"
	"github.com/lib/pq"
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
}

func NewAppContext(db *sql.DB, conf config.Config) AppContext {
	return AppContext{db: db, conf: conf}
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

	fmt.Println("Command: ", cmd)

	switch cmd {
	case "remind":
		txt := s.Join(arr[1:len(arr)], " ")
		fmt.Println("Text: ", txt)
		ac.save(txt)
	case "list":
		ac.list(chatId)
	case "clear":
		// id := s.Join(arr[1:len(arr)], " ")
		// ac.clear(id)
	default:
		fmt.Println("Invalid command.")
	}
}

func (ac *AppContext) save(txt string) {
	_, err := ac.db.Exec(`INSERT INTO reminders(content) VALUES ($1)`, txt)

	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
	}
}

func (ac *AppContext) clear(id int) {
	_, err := ac.db.Exec(`DELETE FROM reminders WHERE id=$1`, id)

	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
	}
}

func (ac *AppContext) list(chatId int64) {
	rows, err := ac.db.Query(`SELECT content FROM reminders`)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
		return
	}
	defer rows.Close()
	var arr []string
	for rows.Next() {
		var content string
		_ = rows.Scan(&content)
		arr = append(arr, content)
	}
	text := s.Join(arr, "\n")

	link := "https://api.telegram.org/{botId}:{apiKey}/sendMessage?chat_id={chatId}&text={text}"
	link = s.Replace(link, "{botId}", ac.conf.BOT.BotId, -1)
	link = s.Replace(link, "{apiKey}", ac.conf.BOT.ApiKey, -1)
	link = s.Replace(link, "{chatId}", strconv.FormatInt(chatId, 10), -1)
	link = s.Replace(link, "{text}", url.QueryEscape(text), -1)

	_, _ = http.Get(link)
}
