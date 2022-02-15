package activation

import (
	"fmt"
	"io"
	"strconv"
)

type Activation string

const (
	Activated   Activation = "ACTIVATED"
	Deactivated Activation = "DEACTIVATED"
)

func (a Activation) String() string {
	return string(a)
}

func (a Activation) Values() []string {
	values := make([]string, 0)
	for _, x := range []Activation{Activated, Deactivated} {
		values = append(values, x.String())
	}
	return values
}

func Validator(a Activation) error {
	switch a {
	case Activated, Deactivated:
		return nil

	default:
		return fmt.Errorf("invalid enum value for activation: %q", a)
	}
}

func (a Activation) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(a.String()))
}

func (a *Activation) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*a = Activation(str)
	if err := Validator(*a); err != nil {
		return fmt.Errorf("%s is not a valid Activation", str)
	}
	return nil
}
