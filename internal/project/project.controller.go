package project

import (
	"context"
	"git-project-management/internal/task"

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

func (c *Controller) getOne(_ context.Context, req *GetOneRequest) (*GetOneResponse, error) {

	project, err := c.repo.getByID(req.Id)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[activity] get all")
	}

	return &GetOneResponse{
		Body: ToProjectDTO(*project),
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

func (c *Controller) getAllTasks(_ context.Context, req *GetAllTasksRequest) (*GetAllTasksResponse, error) {
	limit := 100
	offset := 0
	if req.Limit > 0 {
		limit = req.Limit
	}

	if req.Offset > 0 {
		offset = req.Offset
	}

	tasks, err := c.repo.getAllTasks(req.ProjectId, limit, offset)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[activity] get all")
	}

	var tasksDTO []task.TaskDTO
	for _, v := range tasks {
		tasksDTO = append(tasksDTO, task.ToTaskDTO(v))
	}

	return &GetAllTasksResponse{
		Body: tasksDTO,
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
