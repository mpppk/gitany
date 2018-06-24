package gitany

type Issue interface {
	GetBody() string
	GetHTMLURL() string
	GetID() int64
	GetNumber() int
	GetRepositoryID() int64
	//GetRepository() Repository
	GetTitle() string
	GetLabels() []string // FIXME
}
