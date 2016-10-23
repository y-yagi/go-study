package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type Post struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type postHandler struct {
	db *sql.DB
}

func getPosts(db *sql.DB) ([]Post, error) {
	var posts []Post = make([]Post, 100)
	var i = 0

	rows, err := db.db.Query("SELECT * FROM posts limit 100")
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

func (ph *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Started %s %s for %s at %s\n", r.Method, r.RequestURI, r.RemoteAddr, time.Now().Format(time.RFC3339))

	var buffer bytes.Buffer

	posts, err := getPosts(ph.db)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, post := range posts {
		mapPost, _ := json.Marshal(post)
		buffer.WriteString(string(mapPost))
	}

	fmt.Fprint(w, buffer.String())
}

func main() {
	db, err := sql.Open("postgres", "user=yaginuma dbname=api_test")

	if err != nil {
		log.Fatal(err)
		return
	}

	ph := &postHandler{db: db}
	http.Handle("/posts", ph)
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}

}
