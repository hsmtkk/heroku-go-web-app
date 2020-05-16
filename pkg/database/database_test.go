package database_test

import (
	"database/sql"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hsmtkk/heroku-go-web-app/pkg/database"
	"github.com/hsmtkk/heroku-go-web-app/pkg/post"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "database")
	defer os.RemoveAll(tmpDir)
	db, _ := sql.Open("sqlite3", "./test.db")
	defer db.Close()
	ope := database.New(db)

	err := ope.Create(post.Post{ID: 0, Content: "alpha", Author: "bravo"})
	assert.Nil(t, err, "should be nil")

	p, err := ope.Retrieve(0)
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, "alpha", p.Content, "should be equal")
	assert.Equal(t, "bravo", p.Author, "should be equal")

	err = ope.Update(post.Post{ID: 0, Content: "charlie", Author: "delta"})
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, "charlie", p.Content, "should be equal")
	assert.Equal(t, "delta", p.Author, "should be equal")

	err = ope.Delete(p)
	assert.Nil(t, err, "should be nil")
}
