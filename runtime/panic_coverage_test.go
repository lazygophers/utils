package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCachePanicWithHandleNilError(t *testing.T) {
	t.Run("nil_error", func(t *testing.T) {
		var handled bool
		handler := func(err interface{}) {
			handled = true
		}
		defer func() {
			if r := recover(); r != nil {
				CachePanicWithHandle(handler)
			}
		}()
		panic(nil)
		assert.True(t, handled)
	})
}

func TestCachePanicWithHandleStructError(t *testing.T) {
	t.Run("struct_error", func(t *testing.T) {
		type customError struct {
			Code int
			Msg  string
		}
		err := customError{Code: 404, Msg: "not found"}

		var handledErr interface{}
		handler := func(e interface{}) {
			handledErr = e
		}
		defer func() {
			if r := recover(); r != nil {
				CachePanicWithHandle(handler)
			}
		}()
		panic(err)
		assert.Equal(t, err, handledErr)
	})
}

func TestCachePanicWithHandleSliceError(t *testing.T) {
	t.Run("slice_error", func(t *testing.T) {
		err := []string{"error1", "error2"}

		var handledErr interface{}
		handler := func(e interface{}) {
			handledErr = e
		}
		defer func() {
			if r := recover(); r != nil {
				CachePanicWithHandle(handler)
			}
		}()
		panic(err)
		assert.Equal(t, err, handledErr)
	})
}

func TestCachePanicWithHandleMapError(t *testing.T) {
	t.Run("map_error", func(t *testing.T) {
		err := map[string]string{"key": "value"}

		var handledErr interface{}
		handler := func(e interface{}) {
			handledErr = e
		}
		defer func() {
			if r := recover(); r != nil {
				CachePanicWithHandle(handler)
			}
		}()
		panic(err)
		assert.Equal(t, err, handledErr)
	})
}

func TestCachePanicWithHandleFloatError(t *testing.T) {
	t.Run("float_error", func(t *testing.T) {
		err := 3.14

		var handledErr interface{}
		handler := func(e interface{}) {
			handledErr = e
		}
		defer func() {
			if r := recover(); r != nil {
				CachePanicWithHandle(handler)
			}
		}()
		panic(err)
		assert.Equal(t, err, handledErr)
	})
}

func TestCachePanicWithHandleBoolError(t *testing.T) {
	t.Run("bool_error", func(t *testing.T) {
		err := true

		var handledErr interface{}
		handler := func(e interface{}) {
			handledErr = e
		}
		defer func() {
			if r := recover(); r != nil {
				CachePanicWithHandle(handler)
			}
		}()
		panic(err)
		assert.Equal(t, err, handledErr)
	})
}

func TestCachePanicWithHandleNoPanic(t *testing.T) {
	t.Run("no_panic", func(t *testing.T) {
		var handled bool
		handler := func(err interface{}) {
			handled = true
		}

		CachePanicWithHandle(handler)

		assert.False(t, handled)
	})
}
