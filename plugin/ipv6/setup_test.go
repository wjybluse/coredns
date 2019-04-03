package ipv6

import (
	"testing"

	"github.com/mholt/caddy"
)

func TestDisable(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"ipv6 false", false},

		{"ipv6 true", true},
	}

	for _, tt := range tests {
		c := caddy.NewTestController("dns", tt.input)
		err := setup(c)
		r := ipv6Parse(c)
		if r != tt.expect {
			t.Fail()
		}
	}

}
