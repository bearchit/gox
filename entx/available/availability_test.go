package available_test

import (
	"testing"
	"time"

	"github.com/bearchit/gox/entx/available"

	"github.com/bearchit/gox/entx/available/availability"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:funlen
func TestAvailability(t *testing.T) {
	t.Parallel()

	type given struct {
		startAt    time.Time
		endAt      time.Time
		activation available.Activation
	}

	now := time.Now()
	tests := []struct {
		given   given
		want    availability.Availability
		wantErr bool
	}{
		{
			given{now.Add(time.Hour), now.Add(2 * time.Hour), available.Deactivated},
			availability.Deactivated,
			false,
		},
		{
			given{now.Add(time.Hour), now.Add(2 * time.Hour), available.Activated},
			availability.Upcoming,
			false,
		},
		{
			given{now.Add(-2 * time.Hour), now.Add(-1 * time.Hour), available.Deactivated},
			availability.Deactivated,
			false,
		},
		{
			given{now.Add(-2 * time.Hour), now.Add(-1 * time.Hour), available.Activated},
			availability.Ended,
			false,
		},
		{
			given{now.Add(-time.Hour), now.Add(2 * time.Hour), available.Deactivated},
			availability.Deactivated,
			false,
		},
		{
			given{now.Add(-time.Hour), now.Add(2 * time.Hour), available.Activated},
			availability.InProgress,
			false,
		},
		{
			given{now.Add(-time.Hour), now.Add(-2 * time.Hour), available.Activated},
			availability.InProgress,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			got, err := available.Availability(tt.given.startAt, tt.given.endAt, tt.given.activation)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
