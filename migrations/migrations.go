package migrations

import (
	"git-project-management/internal/activity"
	"git-project-management/internal/notification"
	"git-project-management/internal/project"
	"git-project-management/internal/task"
	"git-project-management/internal/user"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func Migrate(db *pg.DB) error {
	models := []interface{}{
		(*project.ProjectEntity)(nil),
		(*task.TaskEntity)(nil),
		(*user.UserEntity)(nil),
		(*notification.NotificationEntity)(nil),
		(*activity.ActivityEntity)(nil),
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
