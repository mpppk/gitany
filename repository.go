package gitany

type Repository interface {
	GetID() int64
	GetName() string
	GetHTMLURL() string
	GetGitURL() string
	GetCloneURL() string
}
