package nilx_test

import (
	"testing"
	"time"

	"github.com/bearchit/gox/nilx"
	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 0, nilx.PtrInt(nil))
	})
	t.Run("value", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 1, nilx.PtrInt(nilx.IntPtr(1)))
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "", nilx.PtrString(nil))
	})
	t.Run("value", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "hello", nilx.PtrString(nilx.StringPtr("hello")))
	})
}

func TestBool(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, false, nilx.PtrBool(nil))
	})
	t.Run("value", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, true, nilx.PtrBool(nilx.BoolPtr(true)))
	})
}

func TestTime(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		assert.True(t, nilx.PtrTime(nil).IsZero())
	})
	t.Run("value", func(t *testing.T) {
		t.Parallel()
		now := time.Now()
		assert.Equal(t, now, nilx.PtrTime(nilx.TimePtr(now)))
	})
}
