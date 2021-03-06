package gitany

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

type IssueListByRepoOptions struct {
	// State filters issues based on their state. Possible values are: open,
	// closed, all. Default is "open".
	State string `url:"state,omitempty"`

	// Labels filters issues based on their label.
	Labels []string `url:"labels,omitempty,comma"`
	ListOptions
}

type MilestoneListOptions struct {
	// State filters milestones based on their state. Possible values are:
	// open, closed, all. Default is "open".
	State string `url:"state,omitempty"`
	ListOptions
}

type IssueListOptions struct {
	Filter string
	// State filters issues based on their state. Possible values are: open,
	// closed, all. Default is "open".
	State string

	// Labels filters issues based on their label.
	Labels []string

	ListOptions
}
