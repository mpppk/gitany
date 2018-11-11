package github

import (
	"github.com/google/go-github/github"
	"github.com/mpppk/gitany"
)

func toGitHubListOptions(opt *gitany.ListOptions) github.ListOptions {
	if opt == nil {
		return github.ListOptions{}
	}

	return github.ListOptions{
		Page:    opt.Page,
		PerPage: opt.PerPage,
	}
}

func toGitHubRepositoryListByOrgOptions(opt *gitany.RepositoryListByOrgOptions) *github.RepositoryListByOrgOptions {
	if opt == nil {
		return nil
	}

	return &github.RepositoryListByOrgOptions{
		Type:        opt.Type,
		ListOptions: toGitHubListOptions(&opt.ListOptions),
	}
}

func toGitHubIssueListByRepoOptions(opt *gitany.IssueListByRepoOptions) *github.IssueListByRepoOptions {
	if opt == nil {
		return nil
	}

	return &github.IssueListByRepoOptions{
		State:       opt.State,
		Labels:      opt.Labels,
		ListOptions: toGitHubListOptions(&opt.ListOptions),
	}
}

func toGitHubIssueListOptions(opt *gitany.IssueListOptions) *github.IssueListOptions {
	if opt == nil {
		return nil
	}

	return &github.IssueListOptions{
		Filter:      opt.Filter,
		State:       opt.State,
		Labels:      opt.Labels,
		ListOptions: toGitHubListOptions(&opt.ListOptions),
	}
}
