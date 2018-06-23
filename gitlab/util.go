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

func toGitLabListGroupIssuesOptions(opt *gitany.IssueListOptions) *gitlab.ListGroupIssuesOptions {
	// FIXME
	return &gitlab.ListGroupIssuesOptions{
		State:  gitlab.String(opt.State),
		Labels: gitlab.Labels(opt.Labels),
		//IIDs:            opt.IIDs,
		//Milestone:       opt.Milestone,
		//Scope:           opt.Scope,
		//AuthorID:        opt.AuthorID,
		//AssigneeID:      opt.AssigneeID,
		//MyReactionEmoji: opt.MyReactionEmoji,
		//OrderBy:         opt.OrderBy,
		Sort: gitlab.String(opt.Sort),
		//Search: opt.Search,
		//CreatedAfter:    opt.CreatedAfter,
		//CreatedBefore:   opt.CreatedBefore,
		//UpdatedAfter:    opt.UpdatedAfter,
		//UpdatedBefore:   opt.UpdatedBefore,
	}
}
