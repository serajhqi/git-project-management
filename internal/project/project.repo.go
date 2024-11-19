package project

import (
	"errors"
	"time"

	"github.com/go-pg/pg/v10"
)

type Repo struct {
	db *pg.DB
}

func NewRepo(db *pg.DB) *Repo {
	return &Repo{db: db}
}

// Create a new project
func (r *Repo) Create(project *Project) error {
	_, err := r.db.Model(project).Insert()
	return err
}

// Get a project by ID
func (r *Repo) GetByID(id int64) (*Project, error) {
	project := &Project{}
	err := r.db.Model(project).Where("id = ?", id).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return project, nil
}

// Update a project
func (r *Repo) Update(project *Project) error {
	_, err := r.db.Model(project).Where("id = ?", project.ID).Update()
	return err
}

// Delete a project (soft delete)
func (r *Repo) Delete(id int64) error {
	_, err := r.db.Model(&Project{}).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).
		Update()
	return err
}

// List all projects (excluding deleted)
func (r *Repo) List() ([]*Project, error) {
	var projects []*Project
	err := r.db.Model(&projects).Where("deleted_at IS NULL").Select()
	return projects, err
}
