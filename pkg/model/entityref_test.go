package model

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
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

func TestEntityRefScan(t *testing.T) {
	type testCase struct {
		s           any
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
			description: "normal",
		},
		{
			s:           nil,
			expected:    EntityRef{},
			description: "nil",
		},
		{
			s:           "invalid:",
			err:         true,
			description: "invalid string",
		},
		{
			s:           42,
			err:         true,
			description: "not a string",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			e := EntityRef{}
			err := e.Scan(tc.s)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, e)
			}
		})
	}
}

func TestEntityRefValue(t *testing.T) {
	type testCase struct {
		e           EntityRef
		expected    driver.Value
		description string
	}
	tcs := []testCase{
		{
			e: EntityRef{
				Kind:      "kind",
				Namespace: "namespace",
				Name:      "name",
			},
			expected:    driver.Value("kind:namespace/name"),
			description: "normal",
		},
		{
			e:           EntityRef{},
			expected:    nil,
			description: "empty",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			value, err := tc.e.Value()
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, value)
		})
	}
}

func TestEntityRefsScan(t *testing.T) {
	type testCase struct {
		ss          any
		expected    []EntityRef
		err         bool
		description string
	}
	tcs := []testCase{
		{
			ss: "kind:namespace/name1 kind:namespace/name2",
			expected: []EntityRef{
				{
					Kind:      "kind",
					Namespace: "namespace",
					Name:      "name1",
				},
				{
					Kind:      "kind",
					Namespace: "namespace",
					Name:      "name2",
				},
			},
			description: "normal multiple",
		},
		{
			ss: "kind:namespace/name1",
			expected: []EntityRef{
				{
					Kind:      "kind",
					Namespace: "namespace",
					Name:      "name1",
				},
			},
			description: "normal single",
		},
		{
			ss:          nil,
			expected:    nil,
			description: "nil",
		},
		{
			ss:          "kind:namespace/name invalid:",
			err:         true,
			description: "invalid entity ref",
		},
		{
			ss:          42,
			err:         true,
			description: "not a string",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			es := EntityRefs{}
			err := es.Scan(tc.ss)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, es.items)
			}
		})
	}
}

func TestEntityRefsValue(t *testing.T) {
	type testCase struct {
		es          EntityRefs
		expected    driver.Value
		description string
	}
	tcs := []testCase{
		{
			es: EntityRefs{
				items: []EntityRef{
					{
						Kind:      "kind",
						Namespace: "namespace",
						Name:      "name1",
					},
					{
						Kind:      "kind",
						Namespace: "namespace",
						Name:      "name2",
					},
				},
			},
			expected:    driver.Value("kind:namespace/name1 kind:namespace/name2"),
			description: "normal multiple",
		},
		{
			es: EntityRefs{
				items: []EntityRef{
					{
						Kind:      "kind",
						Namespace: "namespace",
						Name:      "name1",
					},
				},
			},
			expected:    driver.Value("kind:namespace/name1"),
			description: "normal single",
		},
		{
			es: EntityRefs{
				items: []EntityRef{
					{
						Kind:      "kind",
						Namespace: "namespace",
						Name:      "name1",
					},
					{},
					{
						Kind:      "kind",
						Namespace: "namespace",
						Name:      "name2",
					},
				},
			},
			expected:    driver.Value("kind:namespace/name1 kind:namespace/name2"),
			description: "multiple with empty",
		},
		{
			es:          EntityRefs{},
			expected:    nil,
			description: "empty",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			value, err := tc.es.Value()
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, value)
		})
	}
}

func testEntityRefYAML(t *testing.T) {
	type testCase struct {
		e           EntityRef
		description string
	}
	tcs := []testCase{
		{
			e: EntityRef{
				Kind:      "kind",
				Namespace: "namespace",
				Name:      "name",
			},
			description: "full entity ref",
		},
		{
			e: EntityRef{
				Kind:      "kind",
				Namespace: "",
				Name:      "name",
			},
			description: "no namespace",
		},
		{
			e: EntityRef{
				Kind:      "",
				Namespace: "namespace",
				Name:      "name",
			},
			description: "no kind",
		},
		{
			e: EntityRef{
				Kind:      "",
				Namespace: "",
				Name:      "name",
			},
			description: "name only",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			out, err := yaml.Marshal(tc.e)
			assert.NoError(t, err)
			assert.Equal(t, tc.e.String(), string(out))

			var r EntityRef
			err = yaml.Unmarshal(out, &r)
			assert.NoError(t, err)
			assert.Equal(t, tc.e, r)
		})
	}
}
