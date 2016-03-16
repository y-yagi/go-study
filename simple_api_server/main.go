package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Post struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func getPosts(db *sql.DB) ([]Post, error) {
	var posts []Post = make([]Post, 100)
	var i = 0

	rows, err := db.Query("SELECT * FROM posts limit 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&posts[i].Id, &posts[i].Title, &posts[i].Body, &posts[i].Created_at, &posts[i].Updated_at); err != nil {
			return nil, err
		}
		i += 1
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func main() {
	db, err := sql.Open("postgres", "user=yaginuma dbname=api_test")
	if err != nil {
		log.Fatal(err)
		return
	}

	posts, err := getPosts(db)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, post := range posts {
		mapPost, _ := json.Marshal(post)
		fmt.Println(string(mapPost))
	}
}
