package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Post struct {
	id         int
	title      string
	body       string
	created_at time.Time
	updated_at time.Time
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
		if err := rows.Scan(&posts[i].id, &posts[i].title, &posts[i].body, &posts[i].created_at, &posts[i].updated_at); err != nil {
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
		fmt.Printf("%s %s\n", post.title, post.body)
	}
}
