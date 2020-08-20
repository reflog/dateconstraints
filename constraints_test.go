package date_constraints

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConstraints(t *testing.T) {
	tests := []struct {
		name       string
		date       string
		constraint string
		valid      bool
		error      bool
	}{
		{
			name:       "valid full",
			date:       "2020-03-10T00:00:00Z",
			constraint: "> 2020-03-01T00:00:00Z <= 2020-04-01T00:00:00Z",
			valid:      true,
		},
		{
			name:       "invalid full",
			date:       "2020-03-01T00:00:00Z",
			constraint: "> 2020-03-01T00:00:00Z <= 2020-04-01T00:00:00Z",
			valid:      false,
		},
		{
			name:       "valid small",
			date:       "2020-03-10T00:00:00Z",
			constraint: "> 2020-03-01T00:00:00Z",
			valid:      true,
		},
		{
			name:       "valid minimal",
			date:       "2020-03-01T00:00:00Z",
			constraint: "= 2020-03-01T00:00:00Z",
			valid:      true,
		},
		{
			name:       "invalid minimal",
			date:       "2020-03-10T00:00:00Z",
			constraint: "= 2020-03-01T00:00:00Z",
			valid:      false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, e := NewConstraint(test.constraint)
			if test.error {
				require.Error(t, e)
			} else {
				require.NoError(t, e)
			}
			if test.valid {
				require.NotNil(t, c)
				d, e := time.Parse(time.RFC3339, test.date)
				require.NoError(t, e)
				require.True(t, c.Check(&d))
			}
		})
	}
}
