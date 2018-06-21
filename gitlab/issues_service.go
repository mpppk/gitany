package gitlab

import (
	"context"
	"fmt"

	"github.com/mpppk/gitany"
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

type issuesService struct {
	projectsService gitany.RepositoriesService
	rawClient       rawClient
	raw             IssuesService
	ListOptions     *gitlab.ListOptions
}

func (i *issuesService) ListByRepo(ctx context.Context, owner, repo string) (serviceIssues []gitany.Issue, err error) {
	opt := &gitlab.ListProjectIssuesOptions{ListOptions: *i.ListOptions}
	issues, _, err := i.raw.ListProjectIssues(owner+"/"+repo, opt)

	for _, issue := range issues {
		serviceIssues = append(serviceIssues, &Issue{Issue: issue})
	}

	return serviceIssues, errors.Wrap(err, "Failed to get Issues by raw client in gitlab.Client.GetIssues")
}

func (i *issuesService) ListLabels(ctx context.Context, owner string, repo string) (labels []gitany.Label, err error) {
	gitlabLabels, _, err := i.rawClient.Labels.ListLabels(owner+"/"+repo, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch labels in gitlab.Client.ListLabels")
	}

	for _, githubLabel := range gitlabLabels {
		labels = append(labels, &Label{Label: githubLabel})
	}

	return labels, nil
}

func (i *issuesService) GetIssuesURL(owner, repo string) (string, error) {
	repoUrl, err := i.projectsService.GetURL(owner, repo)
	return repoUrl + "/issues", errors.Wrap(err, "Error occurred in gitlab.Client.GetIssuesURL")
}

func (i *issuesService) GetURL(owner, repo string, id int) (string, error) {
	url, err := i.GetIssuesURL(owner, repo)
	return fmt.Sprintf("%s/%d", url, id), errors.Wrap(err, "Error occurred in gitlab.Client.GetIssueURL")
}
