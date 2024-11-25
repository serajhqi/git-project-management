package activity

import (
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

func (r *Repo) create(activity *ActivityEntity) (*ActivityEntity, error) {

	_, err := r.db.Model(activity).Returning("*").Insert()
	return activity, err
}

func (r *Repo) getByID(id int64) (*ActivityEntity, error) {
	project := &ActivityEntity{}
	err := r.db.Model(project).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (r *Repo) findByUserIDAndTaskID(userID, taskID int64) (*ActivityEntity, error) {
	project := &ActivityEntity{}
	err := r.db.Model(project).Where("created_by = ? AND task_id = ?", userID, taskID).Order("created_at DESC").First()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (r *Repo) getAll(taskID int64, limit, offset int) ([]ActivityEntity, error) {
	var activities []ActivityEntity

	err := r.db.Model(&activities).Where("task_id = ?", taskID).Limit(limit).Offset(offset).Select()

	if err != nil {
		return activities, err
	}
	return activities, nil

}
