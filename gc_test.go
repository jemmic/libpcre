package libpcre

import (
	"runtime"
	"testing"
	"time"
)

func TestGCApiMatch(t *testing.T) {
	tests := map[string]apiMatchTest{
		`simple1`: {
			pattern: "a*abc?xyz+pqr{3}ab{2,}xy{4,5}pq{0,6}AB{0,}zz",
			matchers: []string{"abxyzpqrrrabbxyyyypqAzz"},
			nomatchers: []string{"abxyzpqrrabbxyyyypqAzz"},
		},
		`simple2`: {
			pattern: "(?<=(foo)a)bar",
			matchers: []string{"fooabar"},
			nomatchers: []string{"foobbar"},
		},
		`empty`: {
			pattern: "",
			matchers: []string{"foo"},
			nomatchers: []string{},
		},
		`nonword`: {
			pattern: "/\\*|[\"';)].*--",
			matchers: []string{"aaaaaaaaa/*aaaa", "aaaaa\"rr--zzzzz"},
			nomatchers: []string{"dd/rrr", "\"';)rrrr-ss"},
		},
		`defined`: {
			pattern: "(?(DEFINE)(?<defined>a|bb)(?<ff>fff))(?&defined)",
			matchers: []string{"weragghgh", "eihfbjbb"},
			nomatchers: []string{"cbefff"},
		},
	}
	for name, testit := range tests {
		test := testit
		t.Run(name, func(t *testing.T) {
			sig := make(chan bool)
			reg, err := Compile(test.pattern, UTF8)
			if err != nil {
				t.Errorf("Err should be nil: %v", err)
			}
			runtime.GC()
			for _, m := range test.matchers {
				go func(){
					for {
						select {
						case _ = <- sig:
							return
						default:
							match := reg.MatcherString(m, 0)
							//runtime.GC()
							if !match.Matches() {
								t.Errorf("%s should match %s", m, test.pattern)
							}
						}
					}
				}()
				for i := 0; i < 500; i++ {
					runtime.GC()
					time.Sleep(2 * time.Millisecond)
				}
				sig <- true
			}
			for _, m := range test.nomatchers {
				go func(){
					for {
						select {
						case _ = <- sig:
							return
						default:
							match := reg.MatcherString(m, 0)
							runtime.GC()
							if match.Matches() {
								t.Errorf("%s should not match %s", m, test.pattern)
							}
						}
					}
				}()
				for i := 0; i < 500; i++ {
					runtime.GC()
					time.Sleep(2 * time.Millisecond)
				}
				sig <- true
			}
		})
	}
}

func TestGCApiReplace(t *testing.T) {
	tests := map[string]apiReplaceTest{
		`simple`: {
			pattern: "a*abc?xyz+pqr{3}ab{2,}xy{4,5}pq{0,6}AB{0,}zz",
			instances: []apiReplaceInstance{
				{
					source: "xabxyzpqrrrabbxyyyypqAzzb",
					replacement: "jj",
					product: "xjjb",
				},
				{
					source: "xabxyzpqrrabbxyyyypqAzzb",
					replacement: "jj",
					product: "xabxyzpqrrabbxyyyypqAzzb",
				},
			},
		},
		`simple2`: {
			pattern: "(?<=(foo)a)bar",
			instances: []apiReplaceInstance{
				{
					source: "xfooabarbfooabarc",
					replacement: "jj",
					product: "xfooajjbfooajjc",
				},
				{
					source: "xfoobbarb",
					replacement: "jj",
					product: "xfoobbarb",
				},
			},
		},
		`nonword`: {
			pattern: "/\\*|[\"';)].*--",
			instances: []apiReplaceInstance{
				{
					source: "a/*aaa/*aaa/*aaaa",
					replacement: "j",
					product: "ajaaajaaajaaaa",
				},
				{
					source: "dd/rrr",
					replacement: "j",
					product: "dd/rrr",
				},
			},
		},
		`defined`: {
			pattern: "(?(DEFINE)(?<defined>a|bb)(?<ff>fff))(?&defined)",
			instances: []apiReplaceInstance{
				{
					source: "weragghbbgh",
					replacement: "j",
					product: "werjgghjgh",
				},
				{
					source: "cbefff",
					replacement: "j",
					product: "cbefff",
				},
			},
		},
	}
	for name, testit := range tests {
		test := testit
		t.Run(name, func(t *testing.T) {
			reg, err := Compile(test.pattern, UTF8)
			if err != nil {
				t.Errorf("Err should be nil: %v", err)
			}
			runtime.GC()
			for _, i := range test.instances {
				for j := 0; j < 10; j++ {
					result := string(reg.ReplaceAll([]byte(i.source),
						[]byte(i.replacement), 0))
					runtime.GC()
					if result != i.product {
						t.Errorf("Result %s should be %s", result, i.product)
					}
				}
			}
		})
	}
}
