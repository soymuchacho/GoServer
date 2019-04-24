package Component

type IComponent interface {
	Init() bool
	Update(tm uint64) bool
	End() bool
}

type Base struct {
}
