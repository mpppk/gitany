package gitlab

import (
	"context"
	"fmt"

	"github.com/mpppk/gitany"
	"github.com/mpppk/gitany/etc"
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

type repositoriesService struct {
	raw           RawClient
	host          string
	serviceConfig *etc.ServiceConfig
}

func (r *repositoriesService) GetURL(owner, repo string) (string, error) {
	err := checkOwnerAndRepo(owner, repo)
	return fmt.Sprintf("%s://%s/%s/%s", r.serviceConfig.Protocol, r.host, owner, repo), errors.Wrap(err, "Error occurred in gitlab.Client.GetRepositoryURL")
}

func (r *repositoriesService) Get(ctx context.Context, owner, repo string) (gitany.Repository, error) {
	project, _, err := r.raw.GetProjects().GetProject(owner + "/" + repo)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to get Repository by raw client in gitlab.Client.GetRepository")
	}

	return &Repository{Project: project}, err
}

func (r *repositoriesService) Create(ctx context.Context, repo string) (gitany.Repository, error) {
	opt := &gitlab.CreateProjectOptions{Name: &repo}
	retRepository, _, err := r.raw.GetProjects().CreateProject(opt)
	return &Repository{retRepository}, err
}

func (r *repositoriesService) GetMilestonesURL(owner, repo string) (string, error) {
	repoUrl, err := r.GetURL(owner, repo)
	return repoUrl + "/milestones", errors.Wrap(err, "Error occurred in gitlab.Client.GetMilestonesURL")
}

func (r *repositoriesService) GetMilestoneURL(owner, repo string, id int) (string, error) {
	url, err := r.GetMilestonesURL(owner, repo)
	return fmt.Sprintf("%s/%d", url, id), errors.Wrap(err, "Error occurred in gitlab.Client.GetMilestoneURL")
}

func (r *repositoriesService) GetWikisURL(owner, repo string) (string, error) {
	repoUrl, err := r.GetURL(owner, repo)
	return repoUrl + "/wikis", errors.Wrap(err, "Error occurred in gitlab.Client.GetWikisURL")
}

func (r *repositoriesService) GetCommitsURL(owner, repo string) (string, error) {
	repoUrl, err := r.GetURL(owner, repo)
	return repoUrl + "/commits/master", errors.Wrap(err, "Error occurred in gitlab.Client.GetCommitsURL")
}

func (r *repositoriesService) CreateRelease(ctx context.Context, owner, repo string, newRelease *gitany.NewRelease) (gitany.Release, error) {
	panic("Not Implemented Yet")
	//opt := &gitlab.CreateTagOptions{}
	//tag, _, err := r.rawClient.GetTags().CreateTag(owner+"/"+repo, opt)
	//return tag, err
}

func (r *repositoriesService) ListByOrg(ctx context.Context, org string) (repos []gitany.Repository, err error) {
	projects, _, err := r.raw.GetGroups().ListGroupProjects(org, nil) // FIXME
	if err != nil {
		return nil, errors.Wrap(err, "Error occurred in gitlab.Client.RepositoriesService.ListByOrg")
	}

	for _, project := range projects {
		repos = append(repos, &Repository{project})
	}
	return repos, nil
}
