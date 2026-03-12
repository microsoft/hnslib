//go:build windows
// +build windows

package hcn

import (
	"fmt"
	"testing"

	"golang.org/x/sys/windows"
)

func TestHCNErrorHelpers(t *testing.T) {
	for _, tc := range []struct {
		name     string
		err      error
		checkFn  func(error) bool
		expected bool
	}{
		// IsElementNotFoundError
		{
			name:     "IsElementNotFoundError with matching error",
			err:      newHcnError(windows.Errno(windows.ERROR_NOT_FOUND), "test", ""),
			checkFn:  IsElementNotFoundError,
			expected: true,
		},
		{
			name:     "IsElementNotFoundError with non-matching error",
			err:      newHcnError(windows.Errno(windows.E_NOTIMPL), "test", ""),
			checkFn:  IsElementNotFoundError,
			expected: false,
		},

		// IsPortAlreadyExistsError
		{
			name:     "IsPortAlreadyExistsError with matching error",
			err:      newHcnError(windows.Errno(windows.HCN_E_PORT_ALREADY_EXISTS), "test", ""),
			checkFn:  IsPortAlreadyExistsError,
			expected: true,
		},
		{
			name:     "IsPortAlreadyExistsError with non-matching error",
			err:      newHcnError(windows.Errno(windows.ERROR_NOT_FOUND), "test", ""),
			checkFn:  IsPortAlreadyExistsError,
			expected: false,
		},

		// IsNotImplemented
		{
			name:     "IsNotImplemented with matching error",
			err:      newHcnError(windows.Errno(windows.E_NOTIMPL), "test", ""),
			checkFn:  IsNotImplemented,
			expected: true,
		},
		{
			name:     "IsNotImplemented with non-matching error",
			err:      newHcnError(windows.Errno(windows.ERROR_NOT_FOUND), "test", ""),
			checkFn:  IsNotImplemented,
			expected: false,
		},

		// IsNetworkNotFound
		{
			name:     "IsNetworkNotFoundError with matching error",
			err:      newHcnError(windows.Errno(windows.HCN_E_NETWORK_NOT_FOUND), "test", ""),
			checkFn:  IsNetworkNotFoundError,
			expected: true,
		},
		{
			name:     "IsNetworkNotFoundError with non-matching error",
			err:      newHcnError(windows.Errno(windows.ERROR_NOT_FOUND), "test", ""),
			checkFn:  IsNetworkNotFoundError,
			expected: false,
		},

		// IsEndpointNotFoundError
		{
			name:     "IsEndpointNotFoundError with matching error",
			err:      newHcnError(windows.Errno(windows.HCN_E_ENDPOINT_NOT_FOUND), "test", ""),
			checkFn:  IsEndpointNotFoundError,
			expected: true,
		},
		{
			name:     "IsEndpointNotFoundError with non-matching error",
			err:      newHcnError(windows.Errno(windows.ERROR_NOT_FOUND), "test", ""),
			checkFn:  IsEndpointNotFoundError,
			expected: false,
		},

		// IsPortNotFoundError
		{
			name:     "IsPortNotFoundError with matching error",
			err:      newHcnError(windows.Errno(windows.HCN_E_PORT_NOT_FOUND), "test", ""),
			checkFn:  IsPortNotFoundError,
			expected: true,
		},
		{
			name:     "IsPortNotFoundError with non-matching error",
			err:      newHcnError(windows.Errno(windows.ERROR_NOT_FOUND), "test", ""),
			checkFn:  IsPortNotFoundError,
			expected: false,
		},

		// IsInvalidIPError
		{
			name:     "IsInvalidIPError with matching error",
			err:      newHcnError(windows.Errno(windows.HCN_E_INVALID_IP), "test", ""),
			checkFn:  IsInvalidIPError,
			expected: true,
		},
		{
			name:     "IsInvalidIPError with non-matching error",
			err:      newHcnError(windows.Errno(windows.ERROR_NOT_FOUND), "test", ""),
			checkFn:  IsInvalidIPError,
			expected: false,
		},

		// IsAdapterNotFoundError
		{
			name:     "IsAdapterNotFoundError with matching error",
			err:      newHcnError(windows.Errno(windows.HCN_E_ADAPTER_NOT_FOUND), "test", ""),
			checkFn:  IsAdapterNotFoundError,
			expected: true,
		},
		{
			name:     "IsAdapterNotFoundError with non-matching error",
			err:      newHcnError(windows.Errno(windows.ERROR_NOT_FOUND), "test", ""),
			checkFn:  IsAdapterNotFoundError,
			expected: false,
		},

		// Non-HcnError
		{
			name:     "non-HcnError returns false",
			err:      fmt.Errorf("random error"),
			checkFn:  IsPortAlreadyExistsError,
			expected: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.checkFn(tc.err); got != tc.expected {
				t.Errorf("expected %t, got %t for error: %v", tc.expected, got, tc.err)
			}
			// Also test wrapped error
			wrapped := fmt.Errorf("wrapped: %w", tc.err)
			if got := tc.checkFn(wrapped); got != tc.expected {
				t.Errorf("expected %t for wrapped error, got %t for error: %v", tc.expected, got, wrapped)
			}
		})
	}
}
