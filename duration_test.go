package relaxduration

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	cases := []struct {
		about    string
		expected time.Duration
		input    string
		err      string
	}{
		{
			about:    "default go parsing",
			expected: time.Hour * 3,
			input:    "3h",
		},
		{
			about:    "default go parsing complex",
			expected: time.Hour*3 + time.Minute*22 + 3*time.Second,
			input:    "3h22m3s",
		},
		{
			about:    "parse seconds",
			expected: time.Second * 2,
			input:    "2seconds",
		},
		{
			about:    "parse second",
			expected: time.Second * 3,
			input:    "3second",
		},
		{
			about:    "parse secs",
			expected: time.Second * 12,
			input:    "12secs",
		},
		{
			about:    "parse sec",
			expected: time.Second * 4,
			input:    "4sec",
		},
		{
			about:    "parse minutes",
			expected: time.Minute * 2,
			input:    "2minutes",
		},
		{
			about:    "parse minute",
			expected: time.Minute * 3,
			input:    "3minute",
		},
		{
			about:    "parse mins",
			expected: time.Minute * 6,
			input:    "6mins",
		},
		{
			about:    "parse min",
			expected: time.Minute * 4,
			input:    "4min",
		},
		{
			about:    "parse hours",
			expected: time.Hour * 2,
			input:    "2hours",
		},
		{
			about:    "parse hour",
			expected: time.Hour * 3,
			input:    "3 hour",
		},
		{
			about:    "parse days",
			expected: time.Hour * 24 * 4,
			input:    "4days",
		},
		{
			about:    "parse day",
			expected: time.Hour * 24,
			input:    "1day",
		},
		{
			about:    "parse d",
			expected: time.Hour * 24 * 3,
			input:    "3d",
		},
		{
			about:    "parse weeks",
			expected: time.Hour * 24 * 7 * 2,
			input:    "2weeks",
		},
		{
			about:    "parse week",
			expected: time.Hour * 24 * 7,
			input:    "1week",
		},
		{
			about:    "parse w",
			expected: time.Hour * 24 * 7 * 4,
			input:    "4w",
		},
		{
			about: "without magnitude it errors",
			err:   "invalid format",
			input: "week",
		},
		{
			about: "error in dias",
			err:   "invalid format",
			input: "5dias",
		},
		{
			about: "error in minuten",
			err:   "invalid format",
			input: "8minuten",
		},
		{
			about: "expects error",
			input: "abacate",
			err:   "invalid format",
		},
		{
			about:    "parse day and hours",
			expected: time.Hour*24 + time.Minute*30,
			input:    "1day30m",
		},
		{
			about:    "parse with spaces",
			expected: time.Hour*24*4 + time.Hour*3,
			input:    "4 days 3 hours",
		},
		{
			about:    "parse with spaces",
			expected: time.Hour*24*7*2 + time.Hour*24*3 + time.Hour*3 + time.Minute*12 + time.Second*15,
			input:    "2 weeks 3 day 3 h 12 mins 15 sec",
		},
		{
			about:    "parse with spaces and 'and's",
			expected: time.Hour*24*7*3 + time.Hour*24*3,
			input:    "3 weeks and 3 days",
		},
		{
			about:    "parse with spaces and 'with's",
			expected: time.Hour*3 + time.Minute*3,
			input:    "3 hours with 3 minutes",
		},
		{
			about:    "parse with spaces 'and's, and 'with's",
			expected: time.Hour*24*7*5 + time.Hour*3 + time.Minute*30 + time.Second*15,
			input:    "5 weeks with 3h30min and 15 secs",
		},
	}

	for _, c := range cases {
		t.Run(c.about, func(t *testing.T) {
			g, e := Parse(c.input)
			if c.err != "" {
				if e == nil {
					t.Fatalf("expected error '%v' got none", c.err)
				}
				if e.Error() != c.err {
					t.Fatalf("expected error '%v' got '%v'", c.err, e.Error())
				}
				return
			}
			if e != nil {
				t.Fatalf("unexpected error: %v", e)
			}
			if g != c.expected {
				t.Fatalf("expected duration %v got %v", c.expected, g)
			}
		})
	}

}
