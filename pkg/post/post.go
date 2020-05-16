package post

import "encoding/json"

type Post struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (p Post) Marshal() ([]byte, error) {
	return json.Marshal(&p)
}

func Unmarshal(js []byte) (Post, error) {
	p := Post{}
	if err := json.Unmarshal(js, &p); err != nil {
		return Post{}, err
	}
	return p, nil
}
