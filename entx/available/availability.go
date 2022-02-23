package available

import (
	"errors"
	"time"

	"github.com/bearchit/gox/entx/available/availability"
	"github.com/bearchit/gox/timex"
)

func Availability(startAt, endAt time.Time, activation Activation) (availability.Availability, error) {
	if activation == Deactivated {
		return availability.Deactivated, nil
	}

	tr, err := timex.NewTimeRange(startAt, endAt)
	if err != nil {
		return "", err
	}
	now := time.Now()
	switch {
	case tr.InProgress(now):
		return availability.InProgress, nil
	case tr.Upcoming(now):
		return availability.Upcoming, nil
	case tr.Ended(now):
		return availability.Ended, nil
	}

	return "", errors.New("unexpected availability")
}
