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

func (r *Repo) getAll(limit, offset int) ([]ActivityDto, error) {
	var activities []ActivityDto

	err := r.db.Model(&activities).Limit(limit).Offset(offset).Select()

	if err != nil {
		return activities, err
	}
	return activities, nil

}
