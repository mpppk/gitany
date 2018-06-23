package gitany

import "time"

type RepositoryListByOrgOptions struct {
	// Type of repositories to list. Possible values are: all, public, private,
	// forks, sources, member. Default is "all".
	Type string `url:"type,omitempty"`
	ListOptions
}

type ListOptions struct {
	Page    int
	PerPage int
}

type IssueListOptions struct {
	// Filter specifies which issues to list. Possible values are: assigned,
	// created, mentioned, subscribed, all. Default is "assigned".
	Filter string

	// State filters issues based on their state. Possible values are: open,
	// closed, all. Default is "open".
	State string

	// Labels filters issues based on their label.
	Labels []string

	// Sort specifies how to sort issues. Possible values are: created, updated,
	// and comments. Default value is "created".
	Sort string

	// Direction in which to sort issues. Possible values are: asc, desc.
	// Default is "desc".
	Direction string

	// Since filters issues by time.
	Since time.Time

	ListOptions
}
