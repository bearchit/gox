package timex

import "time"

type TimeRange struct {
	startAt time.Time
	endAt   time.Time
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
