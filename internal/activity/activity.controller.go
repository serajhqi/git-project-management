package activity

import (
	"context"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"gitea.com/logicamp/lc"
	"github.com/danielgtaylor/huma/v2"
)

type Controller struct {
	repo *Repo
}

func NewController(repo *Repo) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) create(_ context.Context, req *CreateActivityRequest) (*CreateActivityResponse, error) {

	data, _ := base64.StdEncoding.DecodeString(req.Body.CommitMessage)
	// ProcessCommitMessage(string(data))

	fmt.Println("-----------------------------")
	fmt.Println(string(data))
	fmt.Println("-----------------------------")

	parsedValues := ParseFullCommitMessage(string(data))

	// Print the extracted values
	// fmt.Println("Parsed Values:")
	// for key, value := range parsedValues {
	// 	fmt.Printf("%s: %s\n", key, value)
	// }

	// Access individual values
	commitMessage := parsedValues["Commit Message"]
	fmt.Println(commitMessage)
	fmt.Println("-----------------------------")
	author := parsedValues["Author"]
	hash := parsedValues["Hash"]
	branch := parsedValues["Branch"]
	date := parsedValues["Date"]
	isoDate := parsedValues["ISO Date"]

	fmt.Println("\nExtracted Fields:")
	fmt.Println("Author:", author)
	fmt.Println("Hash:", hash)
	fmt.Println("Branch:", branch)
	fmt.Println("Date:", date)
	fmt.Println("ISO Date:", isoDate)

	parsedData, err := ParseCommitMessage(commitMessage)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Parsed Data:")
	fmt.Println("Activity Title:", parsedData.ActivityTitle)
	fmt.Println("Activity Description:", parsedData.ActivityDescription)
	fmt.Println("Task ID:", parsedData.TaskID)
	fmt.Println("Task Status:", parsedData.TaskStatus)
	fmt.Println("Project ID:", parsedData.ProjectID)
	fmt.Println("New Task Title:", parsedData.NewTaskTitle)
	fmt.Println("New Task Description:", parsedData.NewTaskDescription)
	fmt.Println("New Task Status:", parsedData.NewTaskStatus)
	fmt.Println("Time Spent (minutes):", parsedData.TimeSpent)

	// duration := 12
	// activityDto := &ActivityEntity{
	// 	TaskId:      2,
	// 	Title:       "",
	// 	Description: "",
	// 	Duration:    &duration,
	// 	CreatedBy:   0,
	// 	CreatedAt:   time.Time{},
	// }

	// id, err := c.repo.create(activityDto)
	// if err != nil {
	// 	return nil, lc.SendInternalErrorResponse(err, "[activity] create")
	// }

	return nil, huma.Error400BadRequest("wrong task id")
}

func (c *Controller) getOne(_ context.Context, req *GetOneRequest) (*GetOneResponse, error) {

	activity, err := c.repo.getByID(req.Id)
	if err != nil {
		return nil, lc.SendInternalErrorResponse(err, "[activity] get all")
	}

	return &GetOneResponse{
		Body: ToActivityDTO(*activity),
	}, nil
}

// ----------------------------------------------------
func ToActivityDTO(model ActivityEntity) ActivityDTO {
	return ActivityDTO{
		TaskId:      model.ID,
		Title:       model.Title,
		Description: model.Description,
		Duration:    model.Duration,
		CreatedBy:   model.CreatedBy,
		CreateAt:    model.CreatedAt,
	}
}

func ParseFullCommitMessage(fullCommitMessage string) map[string]string {
	// Create a map to store the extracted values
	values := make(map[string]string)

	// Split the message into key-value pairs using " | " as the delimiter
	parts := strings.Split(fullCommitMessage, " ||| ")

	// Iterate over each part and split it into key and value
	for _, part := range parts {
		keyValue := strings.SplitN(part, ": ", 2) // Split by ": " into key and value
		if len(keyValue) == 2 {
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])
			values[key] = value
		}
	}

	return values
}

type ParsedData struct {
	ActivityTitle       string
	ActivityDescription string
	TaskID              string
	TaskStatus          string
	ProjectID           string
	NewTaskTitle        string
	NewTaskDescription  string
	NewTaskStatus       string
	TimeSpent           int // Time spent in minutes
}

// ParseCommitMessage extracts relevant details from a structured commit message.
func ParseCommitMessage(commitMessage string) (*ParsedData, error) {
	lines := strings.Split(commitMessage, "\n")
	data := &ParsedData{}

	// Regex for [task-id] line
	taskRegex := regexp.MustCompile(`\[task-(\d+)\]\s*(.+)?`)
	// Regex for [proj-id] line (new task format)
	projectRegex := regexp.MustCompile(`\[proj-(\d+)\]\s*([^|]+?)\s*(?:\|\s*([^|]*))?\s*(?:\|\s*(.*))?$`)
	// Regex for [spent] line
	spentRegex := regexp.MustCompile(`\[spent\]\s*(\d+h)?\s*(\d+m)?`)

	// Extract the title and description from the first two non-empty lines
	titleSet := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if !titleSet {
			data.ActivityTitle = line
			titleSet = true
		} else if data.ActivityDescription == "" {
			data.ActivityDescription = line
		} else {
			break
		}
	}

	// Parse the rest of the lines for task, project, and spent details
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Match [task-id] details
		if matches := taskRegex.FindStringSubmatch(line); matches != nil {
			data.TaskID = matches[1]
			data.TaskStatus = matches[2]
		}

		// Match [proj-id] and new task details
		if matches := projectRegex.FindStringSubmatch(line); matches != nil {
			data.ProjectID = matches[1]
			data.NewTaskTitle = matches[2] // Always present
			// If the description or status exists, capture them, otherwise leave them empty
			if len(matches) > 3 {
				data.NewTaskDescription = matches[3]
			}
			if len(matches) > 4 {
				data.NewTaskStatus = matches[4]
			}
		}

		// Match [spent] time
		if matches := spentRegex.FindStringSubmatch(line); matches != nil {
			// Parse hours and minutes
			var totalMinutes int
			if matches[1] != "" {
				totalMinutes += parseTimeToMinutes(matches[1])
			}
			if matches[2] != "" {
				totalMinutes += parseTimeToMinutes(matches[2])
			}
			data.TimeSpent = totalMinutes
		}
	}

	if data.NewTaskTitle == "" {
		return nil, fmt.Errorf("invalid commit message format: new task title is required")
	}

	return data, nil
}

// parseTimeToMinutes converts time string (e.g., 1h 30m) to total minutes
func parseTimeToMinutes(timeStr string) int {
	re := regexp.MustCompile(`(\d+)(h|m)`)
	matches := re.FindAllStringSubmatch(timeStr, -1)

	totalMinutes := 0
	for _, match := range matches {
		value := match[1]
		unit := match[2]
		if unit == "h" {
			totalMinutes += parseInt(value) * 60
		} else if unit == "m" {
			totalMinutes += parseInt(value)
		}
	}

	return totalMinutes
}

// parseInt converts a string to an integer
func parseInt(str string) int {
	result := 0
	fmt.Sscanf(str, "%d", &result)
	return result
}
