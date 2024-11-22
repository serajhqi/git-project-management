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

func (r *Repo) create(payload *ActivityEntity) (*int, error) {

	result, err := r.db.Model(payload).Insert()
	if err != nil {
		return nil, err
	}
	id := result.RowsReturned()
	return &id, nil
}

func (r *Repo) getByID(id int64) (*ActivityEntity, error) {
	project := &ActivityEntity{}
	err := r.db.Model(project).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return project, nil
}
