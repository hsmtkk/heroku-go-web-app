package webapp

import (
	"database/sql"
	"net/http"
	"path"
	"strconv"

	"github.com/hsmtkk/heroku-go-web-app/pkg/database"
	"github.com/hsmtkk/heroku-go-web-app/pkg/post"
)

type HandlerGenerator interface {
	GenerateHandler() func(http.ResponseWriter, *http.Request)
}

func New(db *sql.DB) HandlerGenerator {
	operator := database.New(db)
	return &handlerGeneratorImpl{operator: operator}
}

type handlerGeneratorImpl struct {
	operator database.Operator
}

func (impl *handlerGeneratorImpl) GenerateHandler() func(http.ResponseWriter, *http.Request) {
	return impl.HandleRequest
}

func (impl *handlerGeneratorImpl) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case http.MethodGet:
		err = impl.handleGet(w, r)
	case http.MethodPost:
		err = impl.handlePost(w, r)
	case http.MethodPut:
		err = impl.handlePut(w, r)
	case http.MethodDelete:
		err = impl.handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (impl *handlerGeneratorImpl) handleGet(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	p, err := impl.operator.Retrieve(id)
	if err != nil {
		return err
	}
	out, err := p.Marshal()
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return nil
}

func (impl *handlerGeneratorImpl) handlePost(w http.ResponseWriter, r *http.Request) error {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	p, err := post.Unmarshal(body)
	if err != nil {
		return err
	}
	err = impl.operator.Create(p)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func (impl *handlerGeneratorImpl) handlePut(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	p, err := impl.operator.Retrieve(id)
	if err != nil {
		return err
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	p, err = post.Unmarshal(body)
	if err != nil {
		return err
	}
	err = impl.operator.Update(p)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func (impl *handlerGeneratorImpl) handleDelete(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	p, err := impl.operator.Retrieve(id)
	if err != nil {
		return err
	}
	err = impl.operator.Delete(p)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func getID(r *http.Request) (int, error) {
	return strconv.Atoi(path.Base(r.URL.Path))
}
