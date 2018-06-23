package github

import (
	"github.com/google/go-github/github"
	"github.com/mpppk/gitany"
)

func toGitHubListOptions(opt gitany.ListOptions) github.ListOptions {
	return github.ListOptions{
		Page:    opt.Page,
		PerPage: opt.PerPage,
	}
}

func toGitHubRepositoryListByOrgOptions(opt *gitany.RepositoryListByOrgOptions) *github.RepositoryListByOrgOptions {
	return &github.RepositoryListByOrgOptions{
		Type:        opt.Type,
		ListOptions: toGitHubListOptions(opt.ListOptions),
	}
}

func toGitHubIssueListOptions(opt *gitany.IssueListOptions) *github.IssueListOptions {
	return &github.IssueListOptions{
		Filter:      opt.Filter,
		State:       opt.State,
		Labels:      opt.Labels,
		Sort:        opt.Sort,
		Direction:   opt.Direction,
		Since:       opt.Since,
		ListOptions: toGitHubListOptions(opt.ListOptions),
	}
}
