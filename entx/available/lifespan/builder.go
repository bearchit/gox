package lifespan

import "time"

type Builder struct {
	Lifespan
}

func (b *Builder) StartAt(v time.Time) *Builder {
	b.startAt = v
	return b
}

func (b *Builder) EndAt(v time.Time) *Builder {
	b.endAt = v
	return b
}

func (b *Builder) Get() Lifespan {
	return b.Lifespan
}

func Build() *Builder {
	return &Builder{}
}
