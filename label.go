package gitany

type Label interface {
	GetID() int64
	GetName() string
	GetDescription() string
}
