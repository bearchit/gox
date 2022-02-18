package lifespan

import (
	"time"
)

type Lifespan struct {
	startAt time.Time
	endAt   time.Time
}

func (ls *Lifespan) Started(at time.Time) bool {
	return ls.startAt.IsZero() || ls.startAt.Before(at)
}

func (ls *Lifespan) Ended(at time.Time) bool {
	if ls.endAt.IsZero() {
		return false
	}
	return ls.endAt.Before(at)
}

func (ls *Lifespan) In(at time.Time) bool {
	return ls.Started(at) && !ls.Ended(at)
}
