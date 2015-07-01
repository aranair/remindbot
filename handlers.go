package remindbot

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type AppContext struct {
	db *sql.DB
}

func NewAppContext(db *sql.DB) AppContext {
	return AppContext{db: db}
}

func (ac *AppContext) CreateHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	if len(query.Get("location")) > 0 {
		lq := query.Get("location")[0]
		fmt.Fprint(w, "Your location query: ", lq)
	}

	fmt.Println(context.Get(r, "params"))
	params := context.Get(r, "params").(httprouter.Params)
	fmt.Println(params.ByName("location"))
}
