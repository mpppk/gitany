package github

import (
	"time"

	"github.com/google/go-github/github"
)

type Milestone struct {
	*github.Milestone
}

func (m *Milestone) GetStartDate() *time.Time {
	return nil
}

func (m *Milestone) GetDueDate() *time.Time {
	return nil
}

func (m *Milestone) GetState() string {
	return *m.State
}
