package gitany

import "time"

type Issue interface {
	GetBody() string
	GetDueDate() *time.Time
	GetHTMLURL() string
	GetID() int64
	GetNumber() int
	GetMilestone() Milestone
	GetRepositoryID() int64
	//GetRepository() Repository
	GetTitle() string
	GetLabels() []string // FIXME
}
