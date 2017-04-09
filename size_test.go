package humanize

import "testing"

func TestSize(t *testing.T) {
	var tests = []struct {
		Input  int64
		Output string
	}{
		{-1, "-1 byte"},
		{0, "0 byte"},
		{1, "1 byte"},
		{2047, "1 kB"},
		{2048, "2 kB"},
		{1 * 1024 * 1024, "1 MB"},
		{1 * 1024 * 1024 * 1024, "1 GB"},
	}

	for _, tt := range tests {
		if want, got := tt.Output, Size(tt.Input); want != got {
			t.Errorf("expected %q with %v, got %q", want, tt.Input, got)
		}
	}
}
