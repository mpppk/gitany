package mock

type label struct {
	ID          int64
	Name        string
	Description string
}

func NewLabel() *label {
	return &label{}
}

func (l *label) GetID() int64 {
	return l.ID
}

func (l *label) GetName() string {
	return l.Name
}

func (l *label) GetDescription() string {
	return l.Description
}
