package gitany

type Repository interface {
	GetName() string
	GetHTMLURL() string
	GetGitURL() string
	GetCloneURL() string
}
