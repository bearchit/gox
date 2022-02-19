package timex

import "time"

type TimeRange struct {
	startAt time.Time
	endAt   time.Time
}

func (ls *TimeRange) Started(at time.Time) bool {
	return ls.startAt.IsZero() || ls.startAt.Before(at)
}

func (ls *TimeRange) Ended(at time.Time) bool {
	if ls.endAt.IsZero() {
		return false
	}
	return ls.endAt.Before(at)
}

func (ls *TimeRange) In(at time.Time) bool {
	return ls.Started(at) && !ls.Ended(at)
}

type TimeRangeBuilder struct {
	TimeRange
}

func (b *TimeRangeBuilder) StartAt(v time.Time) *TimeRangeBuilder {
	b.startAt = v
	return b
}

func (b *TimeRangeBuilder) EndAt(v time.Time) *TimeRangeBuilder {
	b.endAt = v
	return b
}

func (b *TimeRangeBuilder) Get() TimeRange {
	return b.TimeRange
}

func Build() *TimeRangeBuilder {
	return &TimeRangeBuilder{}
}
