package libpcre

import (
	"testing"
)

type apiTest struct {
	pattern string
	matchers []string
	nomatchers []string
}

func TestApi(t *testing.T) {
	tests := map[string]apiTest{
		`simple1`: {
			pattern: "/a*abc?xyz+pqr{3}ab{2,}xy{4,5}pq{0,6}AB{0,}zz/",
			matchers: []string{"abxyzpqrrrabbxyyyypqAzz"},
			nomatchers: []string{"abxyzpqrrabbxyyyypqAzz"},
		},
		`simple2`: {
			pattern: "/(?<=(foo)a)bar/",
			matchers: []string{"fooabar"},
			nomatchers: []string{"foobbar"},
		},
	}
	for name, testit := range tests {
		test := testit
		t.Run(name, func(t *testing.T) {
			reg, err := Compile(test.pattern, UTF8)
			if err != nil {
				t.Errorf("Err should be nil: %v", err)
			}
			for _, m := range test.matchers {
				match := reg.MatcherString(m, 0)
				if !match.Matches() {
					t.Errorf("%s should match %s", m, test.pattern)
				}
			}
			for _, m := range test.nomatchers {
				match := reg.MatcherString(m, 0)
				if match.Matches() {
					t.Errorf("%s should not match %s", m, test.pattern)
				}
			}
		})
	}

}
