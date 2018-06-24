package gitlab

import (
	"time"

	"github.com/xanzy/go-gitlab"
)

type Milestone struct {
	*gitlab.Milestone
}

func (m *Milestone) GetID() int64 {
	return int64(m.ID)
}

func (m *Milestone) GetNumber() int {
	return m.IID
}

func (m *Milestone) GetTitle() string {
	return m.Title
}

func (m *Milestone) GetStartDate() *time.Time {
	return (*time.Time)(m.StartDate)
}

func (m *Milestone) GetDueDate() *time.Time {
	return (*time.Time)(m.StartDate)
}

func (m *Milestone) GetState() string {
	return m.State
}
