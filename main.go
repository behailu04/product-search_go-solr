package main

import (
	api "product-search_go_solr/api"
	config "product-search_go_solr/config"
	"product-search_go_solr/routers"
)

func main() {
	client, err := config.InitSolr()
	if err != nil {
		panic(err)
	}

	log := config.InitLog()

	repo := api.NewSolrRepository(client)
	svc := api.NewService(log, repo)

	app := routers.GetRouter(log, svc)
	app.RunOnAddr(":" + config.Configuration.Port)
}
