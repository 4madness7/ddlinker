package config

import "testing"

type testCase struct {
	input    Config
	expected testExpected
}

type testExpected struct {
	warns map[string]int
	errs  map[string]int
}

func TestValidate(t *testing.T) {
	cases := []testCase{

		// test for valid config
		{
			input: Config{
				[]destination{
					{
						Name:  "test",
						Path:  "~",
						Links: []string{"config.go"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"test": 0},
				errs:  map[string]int{"test": 0},
			},
		},

		// test for destination with no name
		{
			input: Config{
				[]destination{
					{
						Name:  "",
						Path:  "~",
						Links: []string{"config.go"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"": 0},
				errs:  map[string]int{"": 1},
			},
		},

		// test for destination with no path and non existant path
		{
			input: Config{
				[]destination{
					{
						Name:  "test",
						Path:  "",
						Links: []string{"config.go"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"test": 0},
				errs:  map[string]int{"test": 2},
			},
		},

		// test for destinations with same name
		{
			input: Config{
				[]destination{
					{
						Name:  "test",
						Path:  "~",
						Links: []string{"config.go"},
					},
					{
						Name:  "test",
						Path:  "~/bootdev",
						Links: []string{"config.go"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"test": 0},
				errs:  map[string]int{"test": 1},
			},
		},

		// test for destinations with same path
		{
			input: Config{
				[]destination{
					{
						Name:  "test",
						Path:  "~",
						Links: []string{"config.go"},
					},
					{
						Name:  "test1",
						Path:  "~",
						Links: []string{"config.go"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"test": 0, "test1": 0},
				errs:  map[string]int{"test": 0, "test1": 1},
			},
		},

		// test for destination with no absolute path and non existant path
		{
			input: Config{
				[]destination{
					{
						Name:  "test",
						Path:  "internal",
						Links: []string{"config.go"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"test": 0},
				errs:  map[string]int{"test": 2},
			},
		},

		// test for destination with consecutive '/' in path
		{
			input: Config{
				[]destination{
					{
						Name:  "test",
						Path:  "~////bootdev",
						Links: []string{"config.go"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"test": 1},
				errs:  map[string]int{"test": 0},
			},
		},

		// test for destination with no links
		{
			input: Config{
				[]destination{
					{
						Name:  "test",
						Path:  "~",
						Links: []string{},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"test": 1},
				errs:  map[string]int{"test": 0},
			},
		},

		// test for destination with non existant link
		{
			input: Config{
				[]destination{
					{
						Name:  "test",
						Path:  "~",
						Links: []string{"config"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"test": 0},
				errs:  map[string]int{"test": 1},
			},
		},

		// test for destination with non existant link and link contains '/'
		{
			input: Config{
				[]destination{
					{
						Name:  "test",
						Path:  "~",
						Links: []string{"internal/config"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"test": 1},
				errs:  map[string]int{"test": 1},
			},
		},

		// test for destination with no name, no path,
		// non existant path and link, link contains a '/'
		{
			input: Config{
				[]destination{
					{
						Name:  "",
						Path:  "",
						Links: []string{"internal/main.go"},
					},
				},
			},
			expected: testExpected{
				warns: map[string]int{"": 1},
				errs:  map[string]int{"": 4},
			},
		},
	}

	for _, c := range cases {
		warns, errs := c.input.Validate()
		for k, arr := range warns {
			if len(arr) != c.expected.warns[k] {
				t.Errorf("Warnings expected %d, got %d.", c.expected.warns[k], len(arr))
				t.Errorf("Warnings for '%s': %v", k, warns[k])
			}
		}
		for k, arr := range errs {
			if len(arr) != c.expected.errs[k] {
				t.Errorf("Errors expected %d, got %d.", c.expected.errs[k], len(arr))
				t.Errorf("Errors for '%s': %v", k, errs[k])
			}
		}
	}

}
