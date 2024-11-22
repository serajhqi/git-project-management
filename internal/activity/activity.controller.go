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

func (c *Controller) getOne(_ context.Context, req *GetOneRequest) (*GetOneResponse, error) {

	activity, err := c.repo.getByID(req.Id)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[activity] get all")
	}

	return &GetOneResponse{
		Body: ToActivityDTO(*activity),
	}, nil
}

func ToActivityDTO(model ActivityEntity) ActivityDTO {
	return ActivityDTO{
		TaskId:      model.ID,
		Title:       model.Title,
		Description: model.Description,
		Duration:    model.Duration,
		CreatedBy:   model.CreatedBy,
		CreateAt:    model.CreatedAt,
	}
}
