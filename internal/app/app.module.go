package app

import (
	"git-project-management/internal/activity"
	"git-project-management/migrations"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-pg/pg/v10"
)

func Setup(api *huma.API, db *pg.DB) {

	migrations.Migrate(db)

	activity.Setup(api, db)
}
