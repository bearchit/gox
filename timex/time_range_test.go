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
		started bool
		ended   bool
		in      bool
	}

	tests := []struct {
		name  string
		given timex.TimeRange
		wants wants
	}{
		{"Infinite", timex.Build().Get(), wants{true, false, true}},
		{"Started 1 hour before", timex.Build().StartAt(now.Add(-time.Hour)).Get(), wants{true, false, true}},
		{"Start 1 hour later", timex.Build().StartAt(now.Add(time.Hour)).Get(), wants{false, false, false}},
		{"Ended 1 hour before", timex.Build().EndAt(now.Add(-time.Hour)).Get(), wants{true, true, false}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			t.Run("Started", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tt.wants.started, tt.given.Started(now))
			})
			t.Run("Ended", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tt.wants.ended, tt.given.Ended(now))
			})
			t.Run("In", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tt.wants.in, tt.given.In(now))
			})
		})
	}
}
