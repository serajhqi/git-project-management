package task

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-pg/pg/v10"
)

func Setup(api *huma.API, db *pg.DB) {

	controller := NewController(NewRepo(db))

	huma.Register(*api, huma.Operation{
		OperationID: "get-all-tasks",
		Method:      http.MethodGet,
		Path:        "/projects/{project_id}/tasks",
		Summary:     "all tasks",
		Description: "",
		Tags:        []string{"tasks"},
	}, controller.getAll)

}
