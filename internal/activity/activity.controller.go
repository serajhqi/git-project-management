package activity

import (
	"context"
	"git-project-management/migrations"

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

	activityDto := &migrations.ActivityLog{
		ActionType:       "epic",
		UserID:           1,
		Details:          "sth",
		AssociatedEntity: "task",
		EntityID:         1,
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
