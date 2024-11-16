package migrations

import (
	"log"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Project represents a project with tasks.
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

// Task represents a work item within a project.
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

// User represents an individual working in the system.
type User struct {
	tableName struct{}  `pg:"user"`
	ID        int64     `pg:"id,pk"`                    // Unique identifier
	Name      string    `pg:"name,notnull"`             // User's full name
	Email     string    `pg:"email,notnull,unique"`     // User's email address
	Password  string    `pg:"password,notnull"`         // User's hashed password
	CreatedAt time.Time `pg:"created_at,default:now()"` // Timestamp when the user was created
}

// Notification represents a message sent to a user.
type Notification struct {
	tableName  struct{}  `pg:"notification"`
	ID         int64     `pg:"id,pk"`                    // Unique identifier
	UserID     int64     `pg:"user_id,notnull"`          // ID of the user receiving the notification
	EntityType string    `pg:"entity_type,notnull"`      // Type of entity triggering the notification (e.g., "Task", "Project")
	EntityID   int64     `pg:"entity_id,notnull"`        // ID of the entity triggering the notification
	Message    string    `pg:"message,notnull"`          // Notification message
	Type       string    `pg:"type,notnull"`             // Type of notification (e.g., "Info", "Warning", "Critical")
	IsRead     bool      `pg:"is_read"`                  // Whether the notification has been read
	CreatedBy  int64     `pg:"created_by"`               // User ID who created the notification
	CreatedAt  time.Time `pg:"created_at,default:now()"` // Timestamp when the notification was created
}

// Timelog tracks the time spent on tasks by users.
type Activity struct {
	tableName   struct{}  `pg:"activity"`
	ID          int64     `pg:"id,pk"`                    // Unique identifier
	TaskID      int64     `pg:"task_id,notnull"`          // ID of the associated task
	Duration    *int      `pg:"duration"`                 // Duration in minutes (optional for non-timelog actions)
	Action      string    `pg:"action,notnull"`           // Action performed (e.g., "Worked On sth")
	Description string    `pg:"description"`              // Optional description of the action
	CreatedBy   int64     `pg:"created_by,notnull"`       // CreatedBy of the user who performed the action
	CreatedAt   time.Time `pg:"created_at,default:now()"` // Timestamp when the action occurred
}

func Migrate(db *pg.DB) error {
	models := []interface{}{
		(*Project)(nil),
		(*Task)(nil),
		(*User)(nil),
		(*Notification)(nil),
		(*Activity)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true, // Skip if table already exists
		})
		if err != nil {
			log.Printf("Warning: could not create table for %T: %v", model, err)
		}
	}

	// Create indexes
	indexQueries := []string{
		"CREATE INDEX IF NOT EXISTS idx_task_project_id ON task(project_id)",
		"CREATE INDEX IF NOT EXISTS idx_task_assignee_id ON task(assignee_id)",
		"CREATE INDEX IF NOT EXISTS idx_notification_user_id ON notification(user_id)",
		"CREATE INDEX IF NOT EXISTS idx_activity_task_id ON activity(task_id)",
		"CREATE INDEX IF NOT EXISTS idx_activity_created_by ON activity(created_by)",
	}

	for _, query := range indexQueries {
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Warning: could not create index: %v", err)
		}
	}

	return nil
}
