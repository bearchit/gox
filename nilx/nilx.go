package nilx

import "time"

func PtrInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func IntPtr(v int) *int {
	return &v
}

func PtrString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func StringPtr(v string) *string {
	return &v
}

func PtrBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func BoolPtr(v bool) *bool {
	return &v
}

func PtrTime(v *time.Time) time.Time {
	if v == nil {
		return time.Time{}
	}
	return *v
}

func TimePtr(v time.Time) *time.Time {
	return &v
}
