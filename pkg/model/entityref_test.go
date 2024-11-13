package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityRefString(t *testing.T) {
	type testCase struct {
		e           EntityRef
		expected    string
		description string
	}
	tcs := []testCase{
		{
			e: EntityRef{
				Kind:      "kind",
				Namespace: "namespace",
				Name:      "name",
			},
			expected:    "kind:namespace/name",
			description: "full entity ref",
		},
		{
			e: EntityRef{
				Kind:      "kind",
				Namespace: "",
				Name:      "name",
			},
			expected:    "kind:name",
			description: "no namespace",
		},
		{
			e: EntityRef{
				Kind:      "",
				Namespace: "namespace",
				Name:      "name",
			},
			expected:    "namespace/name",
			description: "no kind",
		},
		{
			e: EntityRef{
				Kind:      "",
				Namespace: "",
				Name:      "name",
			},
			expected:    "name",
			description: "name only",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			s := tc.e.String()
			assert.Equal(t, tc.expected, s)
		})
	}
}

func TestMakeEntityRefString(t *testing.T) {
	type testCase struct {
		s           string
		expected    EntityRef
		err         bool
		description string
	}
	tcs := []testCase{
		{
			s: "kind:namespace/name",
			expected: EntityRef{
				Kind:      "kind",
				Namespace: "namespace",
				Name:      "name",
			},
			err:         false,
			description: "full entity ref",
		},
		{
			s: "kind:name",
			expected: EntityRef{
				Kind:      "kind",
				Namespace: "",
				Name:      "name",
			},
			err:         false,
			description: "no namespace",
		},
		{
			s: "namespace/name",
			expected: EntityRef{
				Kind:      "",
				Namespace: "namespace",
				Name:      "name",
			},
			err:         false,
			description: "no kind",
		},
		{
			s: "name",
			expected: EntityRef{
				Kind:      "",
				Namespace: "",
				Name:      "name",
			},
			err:         false,
			description: "name only",
		},
		{
			s:           ":namespace/name",
			err:         true,
			description: "empty kind",
		},
		{
			s:           "kind:/name",
			err:         true,
			description: "empty namespace",
		},
		{
			s:           "kind:namespace/",
			err:         true,
			description: "empty name",
		},
		{
			s:           "",
			err:         true,
			description: "empty string",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			e, err := MakeEntityRef(tc.s)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, e)
			}
		})
	}
}
