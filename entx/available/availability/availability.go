package availability

import (
	"fmt"
	"io"
	"strconv"
)

type Availability string

const (
	// Upcoming 예정됨
	Upcoming Availability = "UPCOMING"
	// InProgress 진행중
	InProgress Availability = "IN_PROGRESS"
	// Ended 종료됨
	Ended Availability = "ENDED"
	// Deactivated 비활성
	Deactivated Availability = "DEACTIVATED"
)

func (e Availability) IsValid() bool {
	switch e {
	case Upcoming, InProgress, Ended, Deactivated:
		return true
	}
	return false
}

func (e Availability) String() string {
	return string(e)
}

func (e *Availability) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Availability(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Availability", str)
	}
	return nil
}

func (e Availability) MarshalGQL(w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(e.String()))
}
