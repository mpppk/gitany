package gitany

import "time"

type Milestone interface {
	GetID() int64
	GetNumber() int
	GetTitle() string
	GetStartDate() *time.Time
	GetDueDate() *time.Time
	GetState() string
}
