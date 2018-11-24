package mock

import (
	"context"

	"github.com/mpppk/gitany"
)

type repositoriesService struct {
	GotRepository     gitany.Repository
	URL               string
	WikisURL          string
	MilestonesURL     string
	MilestoneURL      string
	CommitsURL        string
	CreatedRepository gitany.Repository
	CreatedRelease    gitany.Release
	OrgRepositories   []gitany.Repository
}

func NewRepositoriesSerivce() *repositoriesService {
	return &repositoriesService{}
}

func (rs *repositoriesService) Get(ctx context.Context, owner, repo string) (gitany.Repository, gitany.Response, error) {
	return rs.GotRepository, nil, nil
}

func (rs *repositoriesService) GetURL(owner, repo string) (string, error) {
	return rs.URL, nil
}

func (rs *repositoriesService) GetWikisURL(owner, repo string) (string, error) {
	return rs.WikisURL, nil
}

func (rs *repositoriesService) GetMilestonesURL(owner, repo string) (string, error) {
	return rs.MilestonesURL, nil
}

func (rs *repositoriesService) GetMilestoneURL(owner, repo string, no int) (string, error) {
	return rs.MilestoneURL, nil
}

func (rs *repositoriesService) GetCommitsURL(owner, repo string) (string, error) {
	return rs.CommitsURL, nil
}

func (rs *repositoriesService) Create(ctx context.Context, org string, repo gitany.Repository) (gitany.Repository, gitany.Response, error) {
	return rs.CreatedRepository, nil, nil
}

func (rs *repositoriesService) CreateRelease(ctx context.Context, owner, repo string, newRelease *gitany.NewRelease) (gitany.Release, gitany.Response, error) {
	return rs.CreatedRelease, nil, nil
}

func (rs *repositoriesService) ListByOrg(ctx context.Context, org string, opt *gitany.RepositoryListByOrgOptions) ([]gitany.Repository, gitany.Response, error) {
	return rs.OrgRepositories, nil, nil
}
