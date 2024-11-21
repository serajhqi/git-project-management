package activity

import (
	"context"
	"time"

	"gitea.com/logicamp/lc"
)

type Controller struct {
	repo *Repo
}

func NewController(repo *Repo) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) create(_ context.Context, req *ActivityCreateRequest) (*ActivityCreateResponse, error) {

	duration := 12
	activityDto := &ActivityEntity{
		TaskId:      2,
		Title:       "",
		Description: "",
		Duration:    &duration,
		CreatedBy:   0,
		CreatedAt:   time.Time{},
	}

	id, err := c.repo.create(activityDto)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[activity] create")
	}

	return &ActivityCreateResponse{
		Body: IdBody{
			Id: *id,
		},
	}, nil
}

func (c *Controller) getAll(_ context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	limit := 100
	offset := 0
	if req.Limit > 0 {
		limit = req.Limit
	}

	if req.Offset > 0 {
		offset = req.Offset
	}

	activities, err := c.repo.getAll(limit, offset)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[activity] get all")
	}

	return &GetAllResponse{
		Body: activities,
	}, nil
}
