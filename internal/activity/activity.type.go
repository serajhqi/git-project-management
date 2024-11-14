package activity

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
