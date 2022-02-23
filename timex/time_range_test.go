package timex_test

import (
	"testing"
	"time"

	"github.com/bearchit/gox/timex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeRange(t *testing.T) {
	t.Parallel()

	type wants struct {
		startAt time.Time
		endAt   time.Time
	}

	startAt := time.Now()
	endAt := startAt.Add(time.Hour)
	tests := []struct {
		given *timex.TimeRange
		wants wants
	}{
		{timex.Build().MustGet(), wants{}},
		{timex.Build().StartAt(startAt).MustGet(), wants{startAt: startAt}},
		{timex.Build().EndAt(endAt).MustGet(), wants{endAt: endAt}},
		{timex.Build().StartAt(startAt).EndAt(endAt).MustGet(), wants{startAt, endAt}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.given.StartAt(), tt.wants.startAt, "StartAt")
			assert.Equal(t, tt.given.EndAt(), tt.wants.endAt, "EndAt")
		})
	}
}

func TestTimeRange_Query(t *testing.T) {
	t.Parallel()

	type wants struct {
		inProgress bool
		upcoming   bool
		ended      bool
	}

	now := time.Now()
	tests := []struct {
		name  string
		given *timex.TimeRange
		wants wants
	}{
		{"Infinite", timex.Build().MustGet(), wants{inProgress: true}},
		{"Started 1 hour before", timex.Build().StartAt(now.Add(-time.Hour)).MustGet(), wants{inProgress: true}},
		{"Start 1 hour later", timex.Build().StartAt(now.Add(time.Hour)).MustGet(), wants{upcoming: true}},
		{"Ended 1 hour before", timex.Build().EndAt(now.Add(-time.Hour)).MustGet(), wants{ended: true}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			t.Run("In", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tt.wants.inProgress, tt.given.InProgress(now))
			})
			t.Run("Upcoming", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tt.wants.upcoming, tt.given.Upcoming(now))
			})
			t.Run("Ended", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tt.wants.ended, tt.given.Ended(now))
			})
		})
	}
}

func TestTimeRangeBuilder_Get(t *testing.T) {
	t.Parallel()

	now := time.Now()
	startAt := now.Add(-time.Hour)
	endAt := now.Add(time.Hour)

	tests := []struct {
		given   *timex.TimeRangeBuilder
		wantErr bool
	}{
		{timex.Build(), false},
		{timex.Build().StartAt(startAt), false},
		{timex.Build().EndAt(endAt), false},
		{timex.Build().StartAt(startAt).EndAt(endAt), false},
		{timex.Build().StartAt(startAt).EndAt(startAt), false},
		{timex.Build().StartAt(endAt).EndAt(startAt), true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()
			_, err := tt.given.Get()
			if tt.wantErr {
				require.Error(t, err)
				require.Panics(t, func() { tt.given.MustGet() })
				return
			}
			require.NoError(t, err)
			require.NotPanics(t, func() { tt.given.MustGet() })
		})
	}
}
