package main

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	// Import the `pq` package with a preceding underscore since it is imported as a side
	// effect. The `pq` package is a GO Postgres driver for the `database/sql` package.
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	staticFileDirectory := http.Dir("./static/")
	staticFileServer := http.FileServer(staticFileDirectory)
	staticFileHandler := http.StripPrefix("/", staticFileServer)
	router.Handle("/", staticFileHandler).Methods("GET")
	router.Handle("/add_devotional.html", staticFileHandler).Methods("GET")
	router.HandleFunc("/devotional", getDevotionalHandler).Methods("GET")
	router.HandleFunc("/devotional", createDevotionalHandler).Methods("POST")

	return router
}

func main() {
	connString := `user=devotionaluser 
				   password=Jesus
				   host=localhost
				   port=5432
				   dbname=devotionalinfo
				   sslmode=disable`
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	store = &dbStore{db: db}
	router := newRouter()
	http.ListenAndServe(":8000", router)
}