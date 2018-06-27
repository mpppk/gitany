package gitlab

import (
	"github.com/mpppk/gitany"
	"github.com/xanzy/go-gitlab"
)

func toGitLabListOptions(opt *gitany.ListOptions) gitlab.ListOptions {
	if opt == nil {
		return gitlab.ListOptions{}
	}

	return gitlab.ListOptions{
		Page:    opt.Page,
		PerPage: opt.PerPage,
	}
}

func toGitLabListGroupProjectsOptions(opt *gitany.RepositoryListByOrgOptions) *gitlab.ListGroupProjectsOptions {
	if opt == nil {
		return nil
	}

	return &gitlab.ListGroupProjectsOptions{
		// FIXME support Type field
		ListOptions: toGitLabListOptions(&opt.ListOptions),
	}
}

func toGitLabListProjectIssuesOptions(opt *gitany.IssueListByRepoOptions) *gitlab.ListProjectIssuesOptions {
	if opt == nil {
		return nil
	}

	state := opt.State
	if opt.State == "open" {
		state = "opened"
	}

	// FIXME Add more options
	return &gitlab.ListProjectIssuesOptions{
		State:       gitlab.String(state),
		Labels:      gitlab.Labels(opt.Labels),
		ListOptions: toGitLabListOptions(&opt.ListOptions),
	}
}

func toGitLabListGroupIssuesOptions(opt *gitany.IssueListOptions) *gitlab.ListGroupIssuesOptions {
	if opt == nil {
		return nil
	}

	state := opt.State
	if opt.State == "open" {
		state = "opened"
	}

	// FIXME Add more options
	return &gitlab.ListGroupIssuesOptions{
		State:       gitlab.String(state),
		Labels:      gitlab.Labels(opt.Labels),
		ListOptions: toGitLabListOptions(&opt.ListOptions),
	}
}

func toGitLabListGroupMilestonesOptions(opt *gitany.MilestoneListOptions) *gitlab.ListGroupMilestonesOptions {
	if opt == nil {
		return nil
	}

	state := opt.State
	if opt.State == "open" {
		state = "opened"
	}

	return &gitlab.ListGroupMilestonesOptions{
		State:       state, // FIXME
		ListOptions: toGitLabListOptions(&opt.ListOptions),
	}
}
