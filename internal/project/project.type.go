package project

import "time"

type Project struct {
	tableName   struct{}  `pg:"project"`
	ID          int64     `pg:"id,pk"`                    // Unique identifier
	Name        string    `pg:"name,notnull"`             // Project name
	Description string    `pg:"description"`              // Optional project description
	StartDate   time.Time `pg:"start_date"`               // Project start date
	EndDate     time.Time `pg:"end_date"`                 // Project end date
	CreatedBy   int64     `pg:"created_by"`               // User ID who created the project
	CreatedAt   time.Time `pg:"created_at,default:now()"` // Timestamp when the project was created
}

// ---

type ProjectDto struct {
	ID          int64     `pg:"id,pk"`
	Name        string    `pg:"name,notnull"`
	Description string    `pg:"description"`
	StartDate   time.Time `pg:"start_date"`
	EndDate     time.Time `pg:"end_date"`
	CreatedBy   int64     `pg:"created_by"`
	CreatedAt   time.Time `pg:"created_at,default:now()"`
	// task count
	// done task count
	// activity count
}
type GetAllRequest struct {
	Limit  int `path:"limit"`
	Offset int `path:"offset"`
}

type GetAllResponse struct {
	Body []ProjectDto
}
