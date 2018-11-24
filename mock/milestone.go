package mock

import "time"

type milestone struct {
	ID        int64
	Number    int
	Title     string
	StartDate *time.Time
	DueDate   *time.Time
	State     string
}

func NewMilestone() *milestone {
	return &milestone{}
}

func (m *milestone) GetID() int64 {
	return m.ID
}

func (m *milestone) GetNumber() int {
	return m.Number
}

func (m *milestone) GetTitle() string {
	return m.Title
}

func (m *milestone) GetStartDate() *time.Time {
	return m.StartDate
}

func (m *milestone) GetDueDate() *time.Time {
	return m.DueDate
}

func (m *milestone) GetState() string {
	return m.State
}
