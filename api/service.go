package api

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/vanng822/go-solr/solr"
)

// Service represent the services
type Service interface {
	Create(data ProductForm) (int, error)
	Select(key, value string) ([]map[string]interface{}, int, error)
	Delete(id string) (int, error)
}

type implService struct {
	log        *logrus.Entry
	repository Repository
}

// NewService will create an object that represent the Service interface
func NewService(log *logrus.Entry, r Repository) Service {
	return &implService{log: log, repository: r}
}

func (s *implService) Create(data ProductForm) (int, error) {
	docs := []solr.Document{}
	doc := solr.Document{
		"id":          data.ID,
		"title":       data.Title,
		"author":      data.Author,
		"price":       data.Price,
		"description": data.Description,
	}
	docs = append(docs, doc)
	params := &url.Values{}
	params.Add("commitWithin", "1000")
	params.Add("overwrite", "true")
	err := s.repository.Create(docs, 0, params)
	if err != nil {
		s.log.Errorf("can't create data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}
	return 0, nil
}

func (s *implService) Select(key, value string) ([]map[string]interface{}, int, error) {
	if len(key) == 0 {
		key = "*"
	}

	if len(value) == 0 {
		value = "*"
	}

	docs, err := s.repository.Select(key, value)
	if err != nil {
		s.log.Errorf("can't select data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}
	data := []map[string]interface{}{}
	for _, d := range docs {
		data = append(data, d)
	}
	return data, 0, nil
}

func (s *implService) Delete(id string) (int, error) {
	data := map[string]interface{}{
		"id": id,
	}
	params := &url.Values{}
	params.Add("commitWithin", "1000")
	params.Add("overwrite", "true")
	err := s.repository.Delete(data, params)
	if err != nil {
		s.log.Errorf("can't select data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}
	return 0, nil
}
