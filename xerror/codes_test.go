package xerror

import (
	"testing"
)

type codeShortcutCase struct {
	name    string
	ctor    func(args ...any) *Error
	wantCode int
}

func TestCodeShortcutConstructors(t *testing.T) {
	cases := []codeShortcutCase{
		{"SystemError", NewSystemError, CodeSystem},
		{"InvalidParam", NewInvalidParam, CodeInvalidParam},
		{"NoAuth", NewNoAuth, CodeNoAuth},
		{"NoData", NewNoData, CodeNoData},
		{"Conflict", NewConflict, CodeConflict},
		{"NotLogin", NewNotLogin, CodeNotLogin},
		{"Timeout", NewTimeout, CodeTimeout},
		{"RateLimited", NewRateLimited, CodeRateLimited},
		{"Forbidden", NewForbidden, CodeForbidden},
		{"Unavailable", NewUnavailable, CodeUnavailable},
		{"DataCorrupted", NewDataCorrupted, CodeDataCorrupted},
	}
	for _, c := range cases {
		err := c.ctor("detail")
		if err.Code() != c.wantCode {
			t.Errorf("%s: code=%d, want %d", c.name, err.Code(), c.wantCode)
		}
		if err.Error() != "detail" {
			t.Errorf("%s: Error()=%q, want detail", c.name, err.Error())
		}
	}
}

func TestCodeConstantsUnique(t *testing.T) {
	codes := []int{
		CodeSystem, CodeSuccess,
		CodeInvalidParam, CodeNoAuth, CodeNoData, CodeConflict, CodeNotLogin,
		CodeTimeout, CodeRateLimited, CodeForbidden, CodeUnavailable, CodeDataCorrupted,
	}
	seen := map[int]bool{}
	for _, c := range codes {
		if seen[c] {
			t.Errorf("duplicate code: %d", c)
		}
		seen[c] = true
	}
}
