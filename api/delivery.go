package api

import (
	"net/http"
	cons "product-search_go_solr/constant"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/sirupsen/logrus"
)

type ctrl struct {
	log *logrus.Entry
	svc Service
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl(log *logrus.Entry, svc Service) *ctrl {
	return &ctrl{log, svc}
}

func (ct *ctrl) Create(req ProductForm, r render.Render) {
	errs := req.Validate()
	if len(errs) > 0 {
		r.JSON(http.StatusBadRequest, NewGenericResponse(http.StatusBadRequest, cons.ERR, errs, nil))
		return
	}

	status, err := ct.svc.Create(req)
	if err != nil {
		r.JSON(status, NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}
	r.JSON(http.StatusCreated, NewGenericResponse(http.StatusCreated, cons.OK, []string{"Data has been saved"}, req))
}

func (ct *ctrl) Select(req *http.Request, r render.Render) {
	key := req.URL.Query().Get("key")
	value := req.URL.Query().Get("value")
	data, status, err := ct.svc.Select(key, value)
	if err != nil {
		r.JSON(status, NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}
	r.JSON(http.StatusOK, NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been retrieved"}, data))
}

func (ct *ctrl) Delete(params martini.Params, r render.Render) {
	id := params["id"]
	status, err := ct.svc.Delete(id)
	if err != nil {
		r.JSON(status, NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}
	r.JSON(http.StatusOK, NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been deleted"}, nil))
}
