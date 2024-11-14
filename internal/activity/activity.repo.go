package activity

import "database/sql"

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) create(payload ActivityDto) (*int, error) {

	var id int
	err := r.db.QueryRow(`INSERT INTO activity(title, description, task_id, timelog, by) 
											  VALUES($1, $2, $3, $4, $5) RETURNING id`, payload.Title, payload.Description, payload.TaskId, payload.TimeLog, 1).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
