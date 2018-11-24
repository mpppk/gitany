package mock

import (
	"time"

	"github.com/mpppk/gitany"
)

type issue struct {
	Body         string
	DueDate      *time.Time
	HTMLURL      string
	ID           int64
	Number       int
	Milestone    gitany.Milestone
	RepositoryID int64
	Title        string
	Labels       []string
}

func NewIssue() *issue {
	return &issue{}
}

func (i *issue) GetBody() string {
	return i.Body
}

func (i *issue) GetDueDate() *time.Time {
	return i.DueDate
}

func (i *issue) GetHTMLURL() string {
	return i.HTMLURL
}

func (i *issue) GetID() int64 {
	return i.ID
}

func (i *issue) GetNumber() int {
	return i.Number
}

func (i *issue) GetMilestone() gitany.Milestone {
	return i.Milestone
}

func (i *issue) GetRepositoryID() int64 {
	return i.RepositoryID
}

func (i *issue) GetTitle() string {
	return i.Title
}

func (i *issue) GetLabels() []string {
	return i.Labels
}
