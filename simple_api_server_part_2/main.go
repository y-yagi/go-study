package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
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
	db *gorm.DB
}

func getPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post = make([]Post, 100)

	db.Find(&posts)

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
	db, err := gorm.Open("postgres", "user=yaginuma dbname=api_test")

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
