package timex_test

import (
	"testing"
	"time"

	"github.com/bearchit/gox/timex"
	"github.com/stretchr/testify/assert"
)

func TestTimeRange_Query(t *testing.T) {
	t.Parallel()

	now := time.Now()
	type wants struct {
		inProgress bool
		upcoming   bool
		ended      bool
	}

	tests := []struct {
		name  string
		given timex.TimeRange
		wants wants
	}{
		{"Infinite", timex.Build().Get(), wants{inProgress: true}},
		{"Started 1 hour before", timex.Build().StartAt(now.Add(-time.Hour)).Get(), wants{inProgress: true}},
		{"Start 1 hour later", timex.Build().StartAt(now.Add(time.Hour)).Get(), wants{upcoming: true}},
		{"Ended 1 hour before", timex.Build().EndAt(now.Add(-time.Hour)).Get(), wants{ended: true}},
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
