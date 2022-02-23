package timex

import (
	"errors"
	"time"
)

type TimeRange struct {
	startAt time.Time
	endAt   time.Time
}

func (ls TimeRange) StartAt() time.Time {
	return ls.startAt
}

func (ls TimeRange) EndAt() time.Time {
	return ls.endAt
}

func (ls *TimeRange) InProgress(at time.Time) bool {
	return !ls.Upcoming(at) && !ls.Ended(at)
}

func (ls *TimeRange) Upcoming(at time.Time) bool {
	return !ls.startAt.IsZero() && ls.startAt.After(at)
}

func (ls *TimeRange) Ended(at time.Time) bool {
	if ls.endAt.IsZero() {
		return false
	}
	return ls.endAt.Before(at)
}

func NewTimeRange(startAt, endAt time.Time) (*TimeRange, error) {
	if !endAt.IsZero() && endAt.Before(startAt) {
		return nil, errors.New("endAt is before startAt")
	}
	return &TimeRange{
		startAt: startAt,
		endAt:   endAt,
	}, nil
}

type TimeRangeBuilder struct {
	startAt time.Time
	endAt   time.Time
}

func (b *TimeRangeBuilder) StartAt(v time.Time) *TimeRangeBuilder {
	b.startAt = v
	return b
}

func (b *TimeRangeBuilder) EndAt(v time.Time) *TimeRangeBuilder {
	b.endAt = v
	return b
}

func (b *TimeRangeBuilder) Get() (*TimeRange, error) {
	return NewTimeRange(b.startAt, b.endAt)
}

func (b *TimeRangeBuilder) MustGet() *TimeRange {
	tr, err := b.Get()
	if err != nil {
		panic(err)
	}
	return tr
}

func Build() *TimeRangeBuilder {
	return &TimeRangeBuilder{}
}
