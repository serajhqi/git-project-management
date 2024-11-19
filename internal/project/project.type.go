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

type CreateRequest struct {
	Body struct {
	}
}
type CreateResponse struct{}
