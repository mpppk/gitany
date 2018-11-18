package gitlab

import (
	"context"
	"fmt"

	"github.com/mpppk/gitany"
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

type repositoriesService struct {
	raw           RawClient
	host          string
	serviceConfig *gitany.ServiceConfig
}

func (r *repositoriesService) GetURL(owner, repo string) (string, error) {
	err := checkOwnerAndRepo(owner, repo)
	return fmt.Sprintf("%s://%s/%s/%s", r.serviceConfig.Protocol, r.host, owner, repo), errors.Wrap(err, "Error occurred in gitlab.Client.GetRepositoryURL")
}

func (r *repositoriesService) Get(ctx context.Context, owner, repo string) (gitany.Repository, gitany.Response, error) {
	project, res, err := r.raw.GetProjects().GetProject(owner + "/" + repo)

	if err != nil {
		return nil, &Response{Response: res}, errors.Wrap(err, "Failed to get Repository by raw client in gitlab.Client.GetRepository")
	}

	return &Repository{Project: project}, &Response{Response: res}, err
}

func (r *repositoriesService) Create(ctx context.Context, org string, repo gitany.Repository) (gitany.Repository, gitany.Response, error) {
	opt := &gitlab.CreateProjectOptions{Name: gitlab.String(repo.GetName())}
	if org != "" {
		group, res, err := r.raw.GetGroups().GetGroup(org) // orgからgroup IDを取得してoptに設定する
		if err != nil {
			return nil, &Response{Response: res}, errors.Wrap(err, fmt.Sprintf("failed to fetch group which name is %v", org))
		}

		opt.NamespaceID = gitlab.Int(group.ID)
	}
	retRepository, res, err := r.raw.GetProjects().CreateProject(opt)
	return &Repository{retRepository}, &Response{Response: res}, err
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

func (r *repositoriesService) CreateRelease(ctx context.Context, owner, repo string, newRelease *gitany.NewRelease) (gitany.Release, gitany.Response, error) {
	panic("Not Implemented Yet")
	//opt := &gitlab.CreateTagOptions{}
	//tag, _, err := r.rawClient.GetTags().CreateTag(owner+"/"+repo, opt)
	//return tag, err
}

func (r *repositoriesService) ListByOrg(ctx context.Context, org string, options *gitany.RepositoryListByOrgOptions) (repos []gitany.Repository, res gitany.Response, err error) {
	gitlabOpt := toGitLabListGroupProjectsOptions(options)
	projects, response, err := r.raw.GetGroups().ListGroupProjects(org, gitlabOpt) // FIXME
	if err != nil {
		return nil, &Response{Response: response}, errors.Wrap(err, "Error occurred in gitlab.Client.RepositoriesService.ListByOrg")
	}

	for _, project := range projects {
		repos = append(repos, &Repository{project})
	}
	return repos, &Response{Response: response}, nil
}
