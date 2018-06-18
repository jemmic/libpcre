// +build darwin
// +build pcretest

package libpcre

// #include <sys/resource.h>
// #include <locale.h>
// #include <time.h>
// #include <regex.h>
import "C"
import "unsafe"
import "github.com/elliotchance/c2go/noarch"

func getrlimit(resource int32, rlim *rlimit) int32 {
	return int32(C.getrlimit(C.int(int(resource)), (*C.struct_rlimit)(unsafe.Pointer(rlim))))
}

func setrlimit(resource int32, rlim *rlimit) int32 {
	return int32(C.setrlimit(C.int(int(resource)), (*C.struct_rlimit)(unsafe.Pointer(rlim))))
}

func regcomp(preg *regex_t, pattern *byte, cflags int32) int32 {
	ppreg := (*C.struct___0)(unsafe.Pointer(preg))
	return int32(C.regcomp(ppreg, C.CString(noarch.CStringToString(pattern)), C.int(cflags)))
}

func regerror(errcode int32, preg *regex_t, errbuf *byte, errbuf_size size_t) size_t {
	ppreg := (*C.struct___0)(unsafe.Pointer(preg))
	perrbuf := (*C.char)(unsafe.Pointer(errbuf))
	return size_t(C.regerror(C.int(errcode), ppreg, perrbuf, C.size_t(errbuf_size)))
}

func regexec(preg *regex_t, str *byte, nmatch size_t, pmatch *regmatch_t, eflags int32) int32 {
	ppreg := (*C.struct___0)(unsafe.Pointer(preg))
	ppmatch := (*C.struct___1)(unsafe.Pointer(pmatch))
	cstr := C.CString(noarch.CStringToString(str))
	return int32(C.regexec(ppreg, cstr, C.size_t(nmatch), ppmatch, C.int(eflags)))
}

func regfree(preg *regex_t) {
	ppreg := (*C.struct___0)(unsafe.Pointer(preg))
	C.regfree(ppreg)
}
