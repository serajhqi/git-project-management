package activity

import "time"

type ActivityEntity struct {
	tableName   struct{}  `pg:"activity"`
	ID          int64     `pg:"id,pk"`                    // Unique identifier
	TaskId      int64     `pg:"task_id,notnull"`          // ID of the associated task
	Title       string    `pg:"title,notnull"`            // Action performed (e.g., "Worked On sth")
	Description string    `pg:"description"`              // Optional description of the action
	Duration    *int      `pg:"duration"`                 // Duration in minutes (optional for non-timelog actions)
	CreatedBy   int64     `pg:"created_by,notnull"`       // CreatedBy of the user who performed the action
	CreatedAt   time.Time `pg:"created_at,default:now()"` // Timestamp when the action occurred
}

type IdBody struct {
	Id int `json:"id"`
}

// Create ---
type CreateActivityDto struct {
	CommitMessage string `json:"commit_message"`
}
type ActivityCreateRequest struct {
	Body CreateActivityDto
}
type ActivityCreateResponse struct {
	Body IdBody
}

// Get ---
type ActivityDto struct {
	TaskId      int64       `json:"task_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Duration    string      `json:"duration"`
	CreatedBy   int64       `json:"created_by"`
	CreateAt    time.Ticker `json:"created_at"`
}

type GetAllRequest struct {
	Limit  int `path:"limit"`
	Offset int `path:"offset"`
}
type GetAllResponse struct {
	Body []ActivityDto
}
