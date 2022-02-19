package available

import (
	"errors"
	pkgactivation "github.com/bearchit/gox/entx/available/activation"
	"github.com/bearchit/gox/entx/available/availability"
	"github.com/bearchit/gox/timex"
	"time"
)

func Availability(startAt, endAt time.Time, activation pkgactivation.Activation) (availability.Availability, error) {
	if activation == pkgactivation.Deactivated {
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
