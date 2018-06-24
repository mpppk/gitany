package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/mpppk/gitany"
	"github.com/pkg/errors"
)

type issuesService struct {
	rawClient   RawClient
	client      gitany.Client
	ListOptions *github.ListOptions
}

func (i *issuesService) GetIssuesURL(owner, repo string) (string, error) {
	if err := checkOwnerAndRepo(owner, repo); err != nil {
		return "", errors.Wrap(err, "Invalid owner or repo was passed to GetIssuesURL")
	}

	repoUrl, err := i.client.GetRepositories().GetURL(owner, repo)
	return repoUrl + "/issues", errors.Wrap(err, "Error occurred in github.Client.GetIssuesURL")
}

func (i *issuesService) GetURL(owner, repo string, id int) (string, error) {
	url, err := i.GetIssuesURL(owner, repo)
	return fmt.Sprintf("%s/%d", url, id), errors.Wrap(err, "Error occurred in github.Client.GetIssueURL")
}

func (i *issuesService) ListByRepo(ctx context.Context, owner, repo string) (serviceIssues []gitany.Issue, err error) {
	opt := &github.IssueListByRepoOptions{ListOptions: *i.ListOptions}
	issues, err := i.getGitHubIssues(ctx, owner, repo, opt)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to get Issues by rawClient client in github.Client.GetIssues")
	}

	for _, issue := range issues {
		serviceIssues = append(serviceIssues, &Issue{Issue: issue})
	}

	return serviceIssues, errors.Wrap(err, "Error occurred in github.Client.GetIssues")
}

func (i *issuesService) ListByOrg(ctx context.Context, org string, opt *gitany.IssueListOptions) ([]gitany.Issue, gitany.Response, error) {
	githubOpt := toGitHubIssueListOptions(opt)
	githubIssues, res, err := i.rawClient.GetIssues().ListByOrg(ctx, org, githubOpt)
	if err != nil {
		return nil, &Response{Response: res}, errors.Wrap(err, "failed to list github organizations")
	}

	var issues []gitany.Issue
	for _, githubIssue := range githubIssues {
		issues = append(issues, &Issue{Issue: githubIssue})
	}

	return issues, &Response{Response: res}, nil

}

func (i *issuesService) getGitHubIssues(ctx context.Context, owner, repo string, opt *github.IssueListByRepoOptions) (issues []*github.Issue, err error) {
	issuesAndPRs, _, err := i.rawClient.GetIssues().ListByRepo(ctx, owner, repo, opt)

	if err != nil {
		return nil, errors.Wrap(err, "Error occurred in github.Client.getGitHubIssues")
	}

	for _, issueOrPR := range issuesAndPRs {
		if issueOrPR.PullRequestLinks == nil {
			issues = append(issues, issueOrPR)
		}
	}
	return issues, nil
}

func (i *issuesService) ListLabels(ctx context.Context, owner string, repo string, opt *gitany.ListOptions) (labels []gitany.Label, res gitany.Response, err error) {
	githubOpt := toGitHubListOptions(opt)

	githubLabels, response, err := i.rawClient.GetIssues().ListLabels(ctx, owner, repo, &githubOpt)
	if err != nil {
		return nil, &Response{Response: response}, errors.Wrap(err, "failed to fetch labels in github.Client.ListLabels")
	}

	for _, githubLabel := range githubLabels {
		labels = append(labels, githubLabel)
	}

	return labels, &Response{Response: response}, nil
}
