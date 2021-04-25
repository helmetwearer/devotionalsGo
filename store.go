package main

import (
	"database/sql"
)

type Store interface {
	CreateDevotional(devotional *Devotional) error
	GetDevotional() ([]*Devotional, error)
}

type dbStore struct {
	db *sql.DB
}

var store Store

func (store *dbStore) CreateDevotional(devotional *Devotional) error {
	_, err := store.db.Query(
		"INSERT INTO devotionalinfo(source,content,dateposted,authorname) VALUES ($1,$2,$3, $4)",
		devotional.Source, devotional.Content, devotional.DatePosted, devotional.AuthorName)
	return err
}

func (store *dbStore) GetDevotional() ([]*Devotional, error) {
	rows, err := store.db.Query("SELECT source, content, dateposted, authorname FROM devotionalinfo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devotionalList := []*Devotional{}
	for rows.Next() {
		devotional := &Devotional{}
		if err := rows.Scan(&devotional.Source, &devotional.Content, &devotional.DatePosted, &devotional.AuthorName); err != nil {
			return nil, err
		}
		devotionalList = append(devotionalList, devotional)
	}
	return devotionalList, nil
}
