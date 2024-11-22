package task

import (
	"git-project-management/internal/activity"

	"github.com/go-pg/pg/v10"
)

type Repo struct {
	db *pg.DB
}

func NewRepo(db *pg.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) getByID(id int64) (*TaskEntity, error) {
	project := &TaskEntity{}
	err := r.db.Model(project).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (r *Repo) getAllActivities(taskId int64, limit, offset int) ([]activity.ActivityEntity, error) {
	var activities []activity.ActivityEntity

	err := r.db.Model(&activities).Where("task_id = ?", taskId).Limit(limit).Offset(offset).Select()

	if err != nil {
		return activities, err
	}
	return activities, nil

}
