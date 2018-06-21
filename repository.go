package gitany

type Repository interface {
	GetHTMLURL() string
	GetGitURL() string
	GetCloneURL() string
}
