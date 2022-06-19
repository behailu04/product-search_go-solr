package api

import (
	"net/url"

	"github.com/vanng822/go-solr/solr"
)

type Repository interface {
	Create(docs []solr.Document, chunkSize int, params *url.Values) error
	Select(key, value string) ([]solr.Document, error)
	Delete(data map[string]interface{}, params *url.Values) error
}

type solrRepository struct {
	Client *solr.SolrInterface
}

func NewSolrRepository(Client *solr.SolrInterface) Repository {
	return &solrRepository{Client}
}

func (s *solrRepository) Create(docs []solr.Document, chunkSize int, params *url.Values) error {
	_, err := s.Client.Add(docs, chunkSize, params)
	return err
}

func (s *solrRepository) Select(key, value string) ([]solr.Document, error) {
	query := solr.NewQuery()
	query.Q(key + ":" + value)
	search := s.Client.Search(query)
	r, err := search.Result(nil)
	if err != nil {
		return nil, err
	}
	return r.Results.Docs, nil
}

func (s *solrRepository) Delete(data map[string]interface{}, params *url.Values) error {
	_, err := s.Client.Delete(data, params)
	return err
}
