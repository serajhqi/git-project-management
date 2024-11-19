package project

type Controller struct {
	repo *Repo
}

func NewController(repo *Repo) *Controller {
	return &Controller{
		repo: repo,
	}
}
