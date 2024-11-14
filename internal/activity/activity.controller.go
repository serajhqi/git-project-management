package activity

import (
	"context"

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

	activityDto := ActivityDto{
		Title:       req.Body.Title,
		Description: req.Body.Description,
		TaskId:      req.TaskId,
		TimeLog:     req.Body.TimeLog,
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
