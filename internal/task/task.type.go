package task

import "time"

type TaskEntity struct {
	tableName   struct{}  `pg:"task"`
	ID          int64     `pg:"id,pk"`                    // Unique identifier
	ParentID    int64     `pg:"parent_id"`                // Unique identifier
	Title       string    `pg:"title,notnull"`            // Task title
	Description string    `pg:"description"`              // Task description
	Status      string    `pg:"status,notnull"`           // Task status (e.g., "Pending", "Completed")
	Priority    string    `pg:"priority,notnull"`         // Task priority (e.g., "Low", "Medium", "High")
	AssigneeID  int64     `pg:"assignee_id"`              // ID of the user assigned to this task
	ProjectID   int64     `pg:"project_id,notnull"`       // ID of the project this task belongs to
	DueDate     time.Time `pg:"due_date"`                 // Due date for the task
	CreatedBy   int64     `pg:"created_by"`               // User ID who created the task
	CreatedAt   time.Time `pg:"created_at,default:now()"` // Timestamp when the task was created
}

// Get All ---
type TaskDTO struct {
	ID          int64     `json:"id"`
	ParentID    int64     `json:"parent_id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	AssigneeID  int64     `json:"assignee_id"`
	ProjectID   int64     `json:"project_id"`
	DueDate     time.Time `json:"due_date"`
	CreatedBy   int64     `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetAllRequest struct {
	Limit     int   `query:"limit"`
	Offset    int   `query:"offset"`
	ProjectId int64 `path:"project_id"`
}
type GetAllResponse struct {
	Body []TaskDTO
}
