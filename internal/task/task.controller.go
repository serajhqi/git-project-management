package task

import (
	"context"
	"errors"
	"git-project-management/internal/activity"

	"gitea.com/logicamp/lc"
	"github.com/go-pg/pg/v10"
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

	task, err := c.repo.getByID(req.Id)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[activity] get one")
	}

	if errors.Is(err, pg.ErrNoRows) {
		return nil, errors.New("project not found")
	}

	return &GetOneResponse{
		Body: ToTaskDTO(*task),
	}, nil
}

func (c *Controller) setStatus(_ context.Context, req *SetTaskStatusRequest) (*SetTaskStatusResponse, error) {

	task, err := c.repo.getByID(req.TaskID)
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, errors.New("project not found")
		}
		return nil, lc.SendInternalErrorResponse(err, "[task] update task")
	}

	task.Status = req.Body.Status
	updatedTask, err := c.repo.Update(req.TaskID, task)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[task] update task")
	}
	return &SetTaskStatusResponse{
		Body: ToTaskDTO(*updatedTask),
	}, nil
}

func (c *Controller) getAllActivities(_ context.Context, req *GetAllActivitiesRequest) (*GetAllActivitiesResponse, error) {
	limit := 100
	offset := 0
	if req.Limit > 0 {
		limit = req.Limit
	}

	if req.Offset > 0 {
		offset = req.Offset
	}

	activities, err := c.repo.getAllActivities(req.TaskID, limit, offset)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[activity] get all")
	}

	var activitiesDTO []activity.ActivityDTO
	for _, v := range activities {
		activitiesDTO = append(activitiesDTO, activity.ToActivityDTO(v))
	}

	return &GetAllActivitiesResponse{
		Body: activitiesDTO,
	}, nil
}

func ToTaskDTO(model TaskEntity) TaskDTO {
	return TaskDTO{
		ID:          model.ID,
		ParentID:    model.ParentID,
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		Priority:    model.Priority,
		AssigneeID:  model.AssigneeID,
		ProjectID:   model.ProjectID,
		DueDate:     model.DueDate,
		CreatedBy:   model.CreatedBy,
		CreatedAt:   model.CreatedAt,
	}
}
