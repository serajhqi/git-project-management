package project

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-pg/pg/v10"
)

func Setup(api *huma.API, db *pg.DB) {

	controller := NewController(NewRepo(db))

	huma.Register(*api, huma.Operation{
		OperationID: "get-all-projects",
		Method:      http.MethodGet,
		Path:        "/projects",
		Summary:     "all projects",
		Description: "",
		Tags:        []string{"projects"},
	}, controller.getAll)

}
