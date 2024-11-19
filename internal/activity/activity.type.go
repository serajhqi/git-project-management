package activity

import "time"

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

type ActivityEntity struct {
	Id          int
	Title       string
	Description string
	ProjectId   string
	TimeLog     int
	By          int
	At          string
}

type ActivityDto struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TaskId      int    `json:"task_id"`
	TimeLog     int    `json:"timelog"`
}

// ---

type ActivityCreateRequest struct {
	TaskId int `path:"task_id"`
	Body   struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		TimeLog     int    `json:"timelog"`
	}
}

type IdBody struct {
	Id int `json:"id"`
}
type ActivityCreateResponse struct {
	Body IdBody
}
