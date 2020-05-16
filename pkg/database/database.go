package database

import (
	"database/sql"

	"github.com/hsmtkk/heroku-go-web-app/pkg/post"
)

func New(db *sql.DB) Operator {
	return &operatorImpl{db: db}
}

type Operator interface {
	Retrieve(id int) (post.Post, error)
	Create(post.Post) error
	Update(post.Post) error
	Delete(post.Post) error
}

type operatorImpl struct {
	db *sql.DB
}

func (impl *operatorImpl) Retrieve(id int) (post.Post, error) {
	p := post.Post{}
	err := impl.db.QueryRow("SELECT id, content, author from posts where id = $1", id).Scan(&p.ID, &p.Content, &p.Author)
	if err != nil {
		return post.Post{}, err
	}
	return p, nil
}

func (impl *operatorImpl) Create(p post.Post) error {
	stStr := "INSERT INTO posts (content, author) values ($1, $2) returning id"
	st, err := impl.db.Prepare(stStr)
	if err != nil {
		return err
	}
	defer st.Close()
	return st.QueryRow(p.Content, p.Author).Scan(&p.ID)
}

func (impl *operatorImpl) Update(p post.Post) error {
	_, err := impl.db.Exec("UPDATE posts SET content = $2, author = $1 WHERE id = $1", p.ID, p.Content, p.Author)
	return err
}

func (impl *operatorImpl) Delete(p post.Post) error {
	_, err := impl.db.Exec("DELETE FROM posts WHERE id = $1", p.ID)
	return err
}
