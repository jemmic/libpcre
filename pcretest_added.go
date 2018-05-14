// +build pcretest

package libpcre

// #include <sys/resource.h>
// #include <locale.h>
// #include <time.h>
// #include <regex.h>
import "C"

type real_pcre_jit_stack interface{}
type real_pcre16_jit_stack interface{}
type real_pcre32_jit_stack interface{}

func setlocale(category int32, locale []byte) []byte {
	return []byte(C.GoString(C.setlocale(C.int(category), C.CString(CStringToString(locale)))))
}

func clock() clock_t {
	return clock_t(C.clock())
}

// CStringToString returns a string that contains all the bytes in the
// provided C string up until the first NULL character.
func CStringToString(s []byte) string {
	if s == nil {
		return ""
	}

	end := -1
	for i, b := range s {
		if b == 0 {
			end = i
			break
		}
	}

	if end == -1 {
		end = len(s)
	}

	newSlice := make([]byte, end)
	copy(newSlice, s)

	return string(newSlice)
}
