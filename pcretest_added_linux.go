// +build linux
// +build pcretest

package libpcre

// #include <sys/resource.h>
// #include <locale.h>
// #include <time.h>
// #include <regex.h>
import "C"
import "unsafe"

func getrlimit(resource int32, rlim []rlimit) int32 {
	return int32(C.getrlimit(C.__rlimit_resource_t(resource), (*C.struct_rlimit)(unsafe.Pointer(&rlim[0]))))
}

func setrlimit(resource int32, rlim []rlimit) int32 {
	return int32(C.setrlimit(C.__rlimit_resource_t(resource), (*C.struct_rlimit)(unsafe.Pointer(&rlim[0]))))
}

func regcomp(preg []regex_t, pattern []byte, cflags int32) int32 {
	ppreg := (*C.struct_re_pattern_buffer)(unsafe.Pointer(&preg[0]))
	return int32(C.regcomp(ppreg, C.CString(CStringToString(pattern)), C.int(cflags)))
}

func regerror(errcode int32, preg []regex_t, errbuf []byte, errbuf_size size_t) size_t {
	ppreg := (*C.struct_re_pattern_buffer)(unsafe.Pointer(&preg[0]))
	perrbuf := (*C.char)(unsafe.Pointer(&errbuf[0]))
	return size_t(C.regerror(C.int(errcode), ppreg, perrbuf, C.size_t(errbuf_size)))
}

func regexec(preg []regex_t, str []byte, nmatch size_t, pmatch []regmatch_t, eflags int32) int32 {
	ppreg := (*C.struct_re_pattern_buffer)(unsafe.Pointer(&preg[0]))
	ppmatch := (*C.struct___0)(unsafe.Pointer(&pmatch[0]))
	cstr := C.CString(CStringToString(str))
	return int32(C.regexec(ppreg, cstr, C.size_t(nmatch), ppmatch, C.int(eflags)))
}

func regfree(preg []regex_t) {
	ppreg := (*C.struct_re_pattern_buffer)(unsafe.Pointer(&preg[0]))
	C.regfree(ppreg)
}
