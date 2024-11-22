package task

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-pg/pg/v10"
)

func Setup(api *huma.API, db *pg.DB) {

	controller := NewController(NewRepo(db))

	huma.Register(*api, huma.Operation{
		OperationID: "get-one-task",
		Method:      http.MethodGet,
		Path:        "/tasks/{id}",
		Summary:     "get one task",
		Description: "",
		Tags:        []string{"tasks"},
	}, controller.getOne)

	huma.Register(*api, huma.Operation{
		OperationID: "get-all-task-activities",
		Method:      http.MethodGet,
		Path:        "/tasks/{id}/activities",
		Summary:     "get all task activities",
		Description: "",
		Tags:        []string{"tasks"},
	}, controller.getAllActivities)

}
