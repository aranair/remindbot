package remindbot

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
	buf  *bytes.Buffer
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

	fmt.Println("Command: ", cmd)

	switch cmd {
	case "remind":
		txt := s.Join(arr[1:len(arr)], " ")
		b := []byte(txt + "\n")

		fmt.Println("Text: ", txt)
		if _, err := ac.buf.Write(b); err != nil {
			fmt.Println("Write error.")
		}
		// ac.save(txt)
	case "list":
		ac.listBuf(chatId)
	case "clear":
		// id := s.Join(arr[1:len(arr)], " ")
		ac.buf.Reset()
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

func (ac *AppContext) listBuf(chatId int64) {
	buf := ac.buf
	var data []byte
	var s []byte
	for {
		data, _ = buf.ReadBytes('\n')
		if len(data) == 0 {
			break
		}
		s = append(s, data...)
	}
	ac.sendText(chatId, string(s))
}

func (ac *AppContext) sendText(chatId int64, text string) {
	link := "https://api.telegram.org/{botId}:{apiKey}/sendMessage?chat_id={chatId}&text={text}"
	link = s.Replace(link, "{botId}", ac.conf.BOT.BotId, -1)
	link = s.Replace(link, "{apiKey}", ac.conf.BOT.ApiKey, -1)
	link = s.Replace(link, "{chatId}", strconv.FormatInt(chatId, 10), -1)
	if len(text) < 5 {
		link = s.Replace(link, "{text}", url.QueryEscape("No current reminders."), -1)
	} else {
		link = s.Replace(link, "{text}", url.QueryEscape(text), -1)
	}
	fmt.Println(link)

	_, _ = http.Get(link)
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
	ac.sendText(chatId, text)
}
