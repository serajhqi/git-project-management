package project

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

func (c *Controller) getAll(_ context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	limit := 100
	offset := 0
	if req.Limit > 0 {
		limit = req.Limit
	}

	if req.Offset > 0 {
		offset = req.Offset
	}

	projects, err := c.repo.getAll(limit, offset)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[activity] get all")
	}
	var projectsDTO []ProjectDTO
	for _, v := range projects {
		projectsDTO = append(projectsDTO, ToProjectDTO(v))
	}

	return &GetAllResponse{
		Body: projectsDTO,
	}, nil
}

func ToProjectDTO(model ProjectEntity) ProjectDTO {
	return ProjectDTO{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
		CreatedBy:   model.CreatedBy,
		CreatedAt:   model.CreatedAt,
	}
}
