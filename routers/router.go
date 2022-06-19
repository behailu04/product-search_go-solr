package routers

import (
	api "product-search_go_solr/api"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/sirupsen/logrus"
)

func GetRouter(log *logrus.Entry, svc api.Service) *martini.Martini {
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(render.Renderer((render.Options{
		IndentJSON: true,
	})))

	ctrl := api.NewCtrl(log, svc)

	r := martini.NewRouter()
	r.Post(`/create`, binding.Bind(api.ProductForm{}), ctrl.Create)
	r.Get(`/select`, ctrl.Select)
	r.Delete(`/delete/:id`, ctrl.Delete)

	m.Action(r.Handle)

	return m
}
