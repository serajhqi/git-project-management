package task

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

func (r *Repo) getAll(projectId int64, limit, offset int) ([]TaskDTO, error) {
	var tasks []TaskDTO

	err := r.db.Model(&tasks).Where("project_id = ?", projectId).Limit(limit).Offset(offset).Select()

	if err != nil {
		return tasks, err
	}
	return tasks, nil

}
