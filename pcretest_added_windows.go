// +build windows
// +build pcretest

package libpcre

// I couldn't find a regex C library for Windows, so for now these are stub functions to make it compilable.

func getrlimit(resource int32, rlim *rlimit) int32 {
	// nop
	return 0
}

func setrlimit(resource int32, rlim *rlimit) int32 {
	// nop
	return 0
}

func regcomp(preg *regex_t, pattern *byte, cflags int32) int32 {
	// nop
	return 0
}

func regerror(errcode int32, preg *regex_t, errbuf *byte, errbuf_size size_t) size_t {
	// nop
	return 0
}

func regexec(preg *regex_t, str *byte, nmatch size_t, pmatch *regmatch_t, eflags int32) int32 {
	// nop
	return 0
}

func regfree(preg *regex_t) {
	// nop
}
