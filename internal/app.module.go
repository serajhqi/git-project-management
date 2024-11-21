package internal

import (
	"git-project-management/internal/activity"
	"git-project-management/internal/project"
	"git-project-management/internal/task"
	"git-project-management/migrations"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-pg/pg/v10"
)

func Setup(api *huma.API, db *pg.DB) {

	migrations.Migrate(db)

	project.Setup(api, db)
	task.Setup(api, db)
	activity.Setup(api, db)
}
