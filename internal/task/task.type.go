package task

import "time"

type Task struct {
	tableName   struct{}   `pg:"task"`
	ID          int64      `pg:"id,pk"`                    // Unique identifier
	ParentID    int64      `pg:"parent_id"`                // Unique identifier
	Title       string     `pg:"title,notnull"`            // Task title
	Description string     `pg:"description"`              // Task description
	Status      string     `pg:"status,notnull"`           // Task status (e.g., "Pending", "Completed")
	Priority    string     `pg:"priority,notnull"`         // Task priority (e.g., "Low", "Medium", "High")
	AssigneeID  int64      `pg:"assignee_id"`              // ID of the user assigned to this task
	ProjectID   int64      `pg:"project_id,notnull"`       // ID of the project this task belongs to
	DueDate     time.Time  `pg:"due_date"`                 // Due date for the task
	CreatedBy   int64      `pg:"created_by"`               // User ID who created the task
	CreatedAt   time.Time  `pg:"created_at,default:now()"` // Timestamp when the task was created
	DeletedAt   *time.Time `pg:"deleted_at"`               // Soft delete timestamp
}
