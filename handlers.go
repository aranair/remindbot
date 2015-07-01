package remindbot

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	s "strings"

	"github.com/lib/pq"
)

type AppContext struct {
	db *sql.DB
}

func NewAppContext(db *sql.DB) AppContext {
	return AppContext{db: db}
}

type Update struct {
	Id  int64   `json:"update_id"`
	Msg Message `json:"message"`
}

type Message struct {
	Id   int64  `json:"message_id"`
	Text string `json:"text"`
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
	cmd := arr[0]
	txt := s.Join(arr[1:len(arr)], " ")

	fmt.Println(cmd)
	fmt.Println(txt)

	_, err := ac.db.Exec(`INSERT INTO reminders(content) VALUES ($1)`, txt)

	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
	}
}
