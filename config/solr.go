package config

import (
	"fmt"

	"github.com/vanng822/go-solr/solr"
)

func InitSolr() (*solr.SolrInterface, error) {
	client, err := solr.NewSolrInterface(Configuration.Solr.Addr, Configuration.Solr.Core)
	if err != nil {
		return nil, fmt.Errorf("Failed to ping connection to solr: %s", err.Error())
	}

	return client, nil
}
