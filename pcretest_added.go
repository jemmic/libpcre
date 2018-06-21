// +build pcretest

package libpcre

// #include <sys/resource.h>
// #include <locale.h>
// #include <time.h>
// #include <regex.h>
import "C"
import "unsafe"
import "github.com/elliotchance/c2go/noarch"

type real_pcre_jit_stack unsafe.Pointer
type real_pcre16_jit_stack unsafe.Pointer
type real_pcre32_jit_stack unsafe.Pointer

func setlocale(category int32, locale *byte) *byte {
	ret := C.setlocale(C.int(category), C.CString(noarch.CStringToString(locale)))
	if ret == nil {
		return nil
	}
	return &[]byte(C.GoString(ret))[0]
}

func clock() clock_t {
	return clock_t(C.clock())
}

func pcre_assign_jit_stack(pcre_extra *pcre_extra, pcre_jit_callback func(arg unsafe.Pointer) *pcre_jit_stack, userdata unsafe.Pointer) {
	panic("not implemented")
}

func pcre_jit_stack_alloc(startsize int32, maxsize int32) *pcre_jit_stack {
	panic("not implemented")
}

func pcre_jit_stack_free(stack *pcre_jit_stack) {
	panic("not implemented")
}
