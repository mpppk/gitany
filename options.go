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
