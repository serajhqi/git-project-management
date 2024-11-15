package migrations

import (
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Entity Struct Definitions

type Project struct {
	ID          int64    `pg:"id,pk"`
	Name        string   `pg:"name,notnull"`
	Description string   `pg:"description"`
	StartDate   string   `pg:"start_date"`
	EndDate     string   `pg:"end_date"`
	Status      string   `pg:"status"`
	Priority    string   `pg:"priority"`
	Owner       int64    `pg:"owner"`
	Budget      float64  `pg:"budget"`
	Tags        []string `pg:"tags,array"`
}

type Task struct {
	ID            int64    `pg:"id,pk"`
	Title         string   `pg:"title,notnull"`
	Description   string   `pg:"description"`
	Status        string   `pg:"status"`
	Priority      string   `pg:"priority"`
	StartDate     string   `pg:"start_date"`
	DueDate       string   `pg:"due_date"`
	Assignee      int64    `pg:"assignee"`
	Dependencies  []int64  `pg:"dependencies,array"`
	EstimatedTime string   `pg:"estimated_time"`
	LoggedTime    string   `pg:"logged_time"`
	Subtasks      []int64  `pg:"subtasks,array"`
	Tags          []string `pg:"tags,array"`
	ProjectID     int64    `pg:"project_id"`
}

type User struct {
	ID                   int64    `pg:"id,pk"`
	Name                 string   `pg:"name,notnull"`
	Email                string   `pg:"email,unique,notnull"`
	Role                 string   `pg:"role"`
	Avatar               string   `pg:"avatar"`
	AvailabilityStatus   string   `pg:"availability_status"`
	Skills               []string `pg:"skills,array"`
	NotificationsEnabled bool     `pg:"notifications_enabled,default:true"`
}

type Milestone struct {
	ID             int64  `pg:"id,pk"`
	Title          string `pg:"title,notnull"`
	Description    string `pg:"description"`
	DueDate        string `pg:"due_date"`
	Status         string `pg:"status"`
	Priority       string `pg:"priority"`
	CompletionDate string `pg:"completion_date"`
	ProjectID      int64  `pg:"project_id"`
}

type Resource struct {
	ID              int64   `pg:"id,pk"`
	Type            string  `pg:"type"`
	Name            string  `pg:"name"`
	URL             string  `pg:"url"`
	Description     string  `pg:"description"`
	UploadedBy      int64   `pg:"uploaded_by"`
	AssociatedTasks []int64 `pg:"associated_tasks,array"`
}

type Comment struct {
	ID        int64   `pg:"id,pk"`
	Author    int64   `pg:"author"`
	Content   string  `pg:"content"`
	Timestamp string  `pg:"timestamp,default:now()"`
	TaskID    int64   `pg:"task_id"`
	ProjectID int64   `pg:"project_id"`
	Replies   []int64 `pg:"replies,array"`
}

type TimeLog struct {
	ID         int64  `pg:"id,pk"`
	TaskID     int64  `pg:"task_id"`
	UserID     int64  `pg:"user_id"`
	Date       string `pg:"date"`
	StartTime  string `pg:"start_time"`
	EndTime    string `pg:"end_time"`
	TotalHours string `pg:"total_hours"`
	Notes      string `pg:"notes"`
}

type ActivityLog struct {
	ID               int64  `pg:"id,pk"`
	ActionType       string `pg:"action_type"`
	UserID           int64  `pg:"user_id"`
	Timestamp        string `pg:"timestamp,default:now()"`
	Details          string `pg:"details"`
	AssociatedEntity string `pg:"associated_entity"`
	EntityID         int64  `pg:"entity_id"`
}

type Dependency struct {
	ID               int64  `pg:"id,pk"`
	DependentTaskID  int64  `pg:"dependent_task_id"`
	DependencyTaskID int64  `pg:"dependency_task_id"`
	Type             string `pg:"type"`
	Status           string `pg:"status"`
}

type Notification struct {
	ID               int64  `pg:"id,pk"`
	UserID           int64  `pg:"user_id"`
	Type             string `pg:"type"`
	Content          string `pg:"content"`
	Timestamp        string `pg:"timestamp,default:now()"`
	ReadStatus       bool   `pg:"read_status,default:false"`
	AssociatedEntity string `pg:"associated_entity"`
	EntityID         int64  `pg:"entity_id"`
}

// CreateTables function for running migrations
func createTables(db *pg.DB) error {
	models := []interface{}{
		(*Project)(nil),
		(*Task)(nil),
		(*User)(nil),
		(*Milestone)(nil),
		(*Resource)(nil),
		(*Comment)(nil),
		(*TimeLog)(nil),
		(*ActivityLog)(nil),
		(*Dependency)(nil),
		(*Notification)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func Migrate(db *pg.DB) {

	// Run the migration to create tables
	err := createTables(db)
	if err != nil {
		log.Printf("Migration encountered some non-fatal issues: %v", err)
	}

	fmt.Println("Tables created successfully!")
}
