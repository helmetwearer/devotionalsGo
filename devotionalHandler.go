package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


type Devotional struct {
	Source		string
	Content		string
	DatePosted	time.Time
	AuthorName 	string
}

func getDevotionalHandler(w http.ResponseWriter, r *http.Request) {

	devotionalList, err := store.GetDevotional()

	devotionalListBytes, err := json.Marshal(devotionalList)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(devotionalListBytes)

}

func createDevotionalHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	devotional := Devotional{}
	devotional.Source = r.Form.Get("source")
	devotional.Content = r.Form.Get("content")
	devotional.DatePosted = time.Now()
	devotional.AuthorName = r.Form.Get("authorname")

	err = store.CreateDevotional(&devotional)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}