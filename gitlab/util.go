package gitlab

import (
	"github.com/mpppk/gitany"
	"github.com/xanzy/go-gitlab"
)

func toGitLabListOptions(opt gitany.ListOptions) gitlab.ListOptions {
	return gitlab.ListOptions{
		Page:    opt.Page,
		PerPage: opt.PerPage,
	}
}

func toGitLabListGroupProjectsOptions(opt *gitany.RepositoryListByOrgOptions) *gitlab.ListGroupProjectsOptions {
	return &gitlab.ListGroupProjectsOptions{
		// FIXME support Type field
		ListOptions: toGitLabListOptions(opt.ListOptions),
	}
}
