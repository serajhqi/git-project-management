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

	d := 12
	activityDto := &migrations.Activity{
		TaskID:      1,
		Duration:    &d,
		Action:      "make things clear",
		Description: "the body of the message",
		CreatedBy:   1,
	}
	// &migrations.Timelog{
	// 	TaskID:    1,
	// 	Duration:  time.Duration(time.Hour),
	// 	CreateBy:  1,
	// 	CreatedAt: time.Time{},
	// }

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
