package project

import (
	"errors"

	"github.com/go-pg/pg/v10"
)

type Repo struct {
	db *pg.DB
}

func NewRepo(db *pg.DB) *Repo {
	return &Repo{db: db}
}

// Create a new project
func (r *Repo) Create(project *ProjectEntity) error {
	_, err := r.db.Model(project).Insert()
	return err
}

// Get a project by ID
func (r *Repo) GetByID(id int64) (*ProjectEntity, error) {
	project := &ProjectEntity{}
	err := r.db.Model(project).Where("id = ?", id).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return project, nil
}

func (r *Repo) getAll(limit, offset int) ([]ProjectEntity, error) {
	var projects []ProjectEntity

	err := r.db.Model(&projects).Limit(limit).Offset(offset).Select()

	if err != nil {
		return projects, err
	}
	return projects, nil

}
