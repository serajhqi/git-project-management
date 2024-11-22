package activity

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-pg/pg/v10"
)

func Setup(api *huma.API, db *pg.DB) {

	controller := NewController(NewRepo(db))

	huma.Register(*api, huma.Operation{
		OperationID: "get-one-activity",
		Method:      http.MethodGet,
		Path:        "/activities/{id}",
		Summary:     "one activity",
		Description: "",
		Tags:        []string{"activity"},
	}, controller.getOne)

	huma.Register(*api, huma.Operation{
		OperationID: "add-activity",
		Method:      http.MethodPost,
		Path:        "/tasks/{task_id}/activities",
		Summary:     "add activity",
		Description: "",
		Tags:        []string{"activity"},
	}, controller.create)
}
