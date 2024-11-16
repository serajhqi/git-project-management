package activity

import (
	"git-project-management/migrations"

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

func (r *Repo) create(payload *migrations.Activity) (*int, error) {

	result, err := r.db.Model(payload).Insert()
	if err != nil {
		return nil, err
	}
	id := result.RowsReturned()
	return &id, nil
}
