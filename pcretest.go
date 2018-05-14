// +build pcretest

package libpcre

import "os"
import "github.com/elliotchance/c2go/linux"
import "github.com/elliotchance/c2go/noarch"
import "unsafe"

type timex struct {
	modes     uint32
	offset    __syscall_slong_t
	freq      __syscall_slong_t
	maxerror  __syscall_slong_t
	esterror  __syscall_slong_t
	status    int32
	constant  __syscall_slong_t
	precision __syscall_slong_t
	tolerance __syscall_slong_t
	time      timeval
	tick      __syscall_slong_t
	ppsfreq   __syscall_slong_t
	jitter    __syscall_slong_t
	shift     int32
	stabil    __syscall_slong_t
	jitcnt    __syscall_slong_t
	calcnt    __syscall_slong_t
	errcnt    __syscall_slong_t
	stbcnt    __syscall_slong_t
	tai       int32
}
type tm struct {
	tm_sec    int32
	tm_min    int32
	tm_hour   int32
	tm_mday   int32
	tm_mon    int32
	tm_year   int32
	tm_wday   int32
	tm_yday   int32
	tm_isdst  int32
	tm_gmtoff int32
	tm_zone   *byte
}
type itimerspec struct {
	it_interval timespec
	it_value    timespec
}
type sigevent struct {
}
type lconv struct {
	decimal_point      *byte
	thousands_sep      *byte
	grouping           *byte
	int_curr_symbol    *byte
	currency_symbol    *byte
	mon_decimal_point  *byte
	mon_thousands_sep  *byte
	mon_grouping       *byte
	positive_sign      *byte
	negative_sign      *byte
	int_frac_digits    byte
	frac_digits        byte
	p_cs_precedes      byte
	p_sep_by_space     byte
	n_cs_precedes      byte
	n_sep_by_space     byte
	p_sign_posn        byte
	n_sign_posn        byte
	int_p_cs_precedes  byte
	int_p_sep_by_space byte
	int_n_cs_precedes  byte
	int_n_sep_by_space byte
	int_p_sign_posn    byte
	int_n_sign_posn    byte
}
type error_t int32
type timezone struct {
	tz_minuteswest int32
	tz_dsttime     int32
}
type __timezone_ptr_t *timezone
type __itimer_which int32

const (
	ITIMER_REAL    __itimer_which = 0
	ITIMER_VIRTUAL                = 1
	ITIMER_PROF                   = 2
)

type itimerval struct {
	it_interval timeval
	it_value    timeval
}
type __itimer_which_t int32
type __rlimit_resource int32

const (
	RLIMIT_CPU          __rlimit_resource = 0
	RLIMIT_FSIZE                          = 1
	RLIMIT_DATA                           = 2
	RLIMIT_STACK                          = 3
	RLIMIT_CORE                           = 4
	__RLIMIT_RSS                          = 5
	RLIMIT_NOFILE                         = 7
	__RLIMIT_OFILE                        = 8
	RLIMIT_AS                             = 9
	__RLIMIT_NPROC                        = 6
	__RLIMIT_MEMLOCK                      = 8
	__RLIMIT_LOCKS                        = 10
	__RLIMIT_SIGPENDING                   = 11
	__RLIMIT_MSGQUEUE                     = 12
	__RLIMIT_NICE                         = 13
	__RLIMIT_RTPRIO                       = 14
	__RLIMIT_RTTIME                       = 15
	__RLIMIT_NLIMITS                      = 16
	__RLIM_NLIMITS                        = 17
)

type rlim_t __rlim_t
type rlim64_t __rlim64_t
type rlimit struct {
	rlim_cur rlim_t
	rlim_max rlim_t
}
type rlimit64 struct {
	rlim_cur rlim64_t
	rlim_max rlim64_t
}
type __rusage_who int32

const (
	RUSAGE_SELF     __rusage_who = 0
	RUSAGE_CHILDREN              = 1
	RUSAGE_THREAD                = 1
)

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD194D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD194D19E) ru_maxrss() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD194D19E) __ru_maxrss_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD202D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD202D19E) ru_ixrss() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD202D19E) __ru_ixrss_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD208D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD208D19E) ru_idrss() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD208D19E) __ru_idrss_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD214D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD214D19E) ru_isrss() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD214D19E) __ru_isrss_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD221D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD221D19E) ru_minflt() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD221D19E) __ru_minflt_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD227D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD227D19E) ru_majflt() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD227D19E) __ru_majflt_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD233D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD233D19E) ru_nswap() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD233D19E) __ru_nswap_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD240D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD240D19E) ru_inblock() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD240D19E) __ru_inblock_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD246D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD246D19E) ru_oublock() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD246D19E) __ru_oublock_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD252D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD252D19E) ru_msgsnd() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD252D19E) __ru_msgsnd_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD258D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD258D19E) ru_msgrcv() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD258D19E) __ru_msgrcv_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD264D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD264D19E) ru_nsignals() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD264D19E) __ru_nsignals_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD272D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD272D19E) ru_nvcsw() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD272D19E) __ru_nvcsw_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD279D19E struct{ memory unsafe.Pointer }

func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD279D19E) ru_nivcsw() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *rusageDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSresourcePhD279D19E) __ru_nivcsw_word() *__syscall_slong_t {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__syscall_slong_t)(unionVar.memory)
}

type rusage struct {
	ru_utime timeval
	ru_stime timeval
}
type __priority_which int32

const (
	PRIO_PROCESS __priority_which = 0
	PRIO_PGRP                     = 1
	PRIO_USER                     = 2
)

type __rlimit_resource_t int32
type __rusage_who_t int32
type __priority_which_t int32

type pcre_jit_stack real_pcre_jit_stack
type pcre16_jit_stack real_pcre16_jit_stack
type pcre32_jit_stack real_pcre32_jit_stack

type pcre_jit_callback func(unsafe.Pointer) *pcre_jit_stack
type pcre16_jit_callback func(unsafe.Pointer) *pcre16_jit_stack
type pcre32_jit_callback func(unsafe.Pointer) *pcre32_jit_stack

var OP_lengths []pcre_uint8 = []pcre_uint8{pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((uint32(int32(1)) + (uint32(int32(32)) / 1))), pcre_uint8((uint32(int32(1)) + (uint32(int32(32)) / 1))), pcre_uint8(int32(0)), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(2) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8(((int32(1) + int32(2)) + int32(2))), pcre_uint8(((int32(1) + int32(2)) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8(((int32(1) + int32(2)) + int32(2))), pcre_uint8(((int32(1) + int32(2)) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8((int32(1) + int32(2))), pcre_uint8(int32(1))}
var hspace_list []pcre_uint32 = []pcre_uint32{pcre_uint32('\t'), pcre_uint32(' '), pcre_uint32((uint8('\u00a0'))), pcre_uint32(int32(5760)), pcre_uint32(int32(6158)), pcre_uint32(int32(8192)), pcre_uint32(int32(8193)), pcre_uint32(int32(8194)), pcre_uint32(int32(8195)), pcre_uint32(int32(8196)), pcre_uint32(int32(8197)), pcre_uint32(int32(8198)), pcre_uint32(int32(8199)), pcre_uint32(int32(8200)), pcre_uint32(int32(8201)), pcre_uint32(int32(8202)), pcre_uint32(int32(8239)), pcre_uint32(int32(8287)), pcre_uint32(int32(12288)), 4294967295}
var vspace_list []pcre_uint32 = []pcre_uint32{pcre_uint32('\n'), pcre_uint32('\v'), pcre_uint32('\f'), pcre_uint32('\r'), pcre_uint32((uint8('\u0085'))), pcre_uint32(int32(8232)), pcre_uint32(int32(8233)), 4294967295}
var utf8_table1 []int32 = []int32{int32(127), int32(2047), int32(65535), int32(2097151), int32(67108863), int32(2147483647)}
var utf8_table1_size int32 = int32((24 / 4))
var utf8_table2 []int32 = []int32{int32(0), int32(192), int32(224), int32(240), int32(248), int32(252)}
var utf8_table3 []int32 = []int32{int32(255), int32(31), int32(15), int32(7), int32(3), int32(1)}
var utf8_table4 []pcre_uint8 = []pcre_uint8{pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(3)), pcre_uint8(int32(3)), pcre_uint8(int32(3)), pcre_uint8(int32(3)), pcre_uint8(int32(3)), pcre_uint8(int32(3)), pcre_uint8(int32(3)), pcre_uint8(int32(3)), pcre_uint8(int32(4)), pcre_uint8(int32(4)), pcre_uint8(int32(4)), pcre_uint8(int32(4)), pcre_uint8(int32(5)), pcre_uint8(int32(5)), pcre_uint8(int32(5)), pcre_uint8(int32(5))}
var ucp_gentype []pcre_uint32 = []pcre_uint32{pcre_uint32(ucp_C), pcre_uint32(ucp_C), pcre_uint32(ucp_C), pcre_uint32(ucp_C), pcre_uint32(ucp_C), pcre_uint32(ucp_L), pcre_uint32(ucp_L), pcre_uint32(ucp_L), pcre_uint32(ucp_L), pcre_uint32(ucp_L), pcre_uint32(ucp_M), pcre_uint32(ucp_M), pcre_uint32(ucp_M), pcre_uint32(ucp_N), pcre_uint32(ucp_N), pcre_uint32(ucp_N), pcre_uint32(ucp_P), pcre_uint32(ucp_P), pcre_uint32(ucp_P), pcre_uint32(ucp_P), pcre_uint32(ucp_P), pcre_uint32(ucp_P), pcre_uint32(ucp_P), pcre_uint32(ucp_S), pcre_uint32(ucp_S), pcre_uint32(ucp_S), pcre_uint32(ucp_S), pcre_uint32(ucp_Z), pcre_uint32(ucp_Z), pcre_uint32(ucp_Z)}
var ucp_gbtable []pcre_uint32 = []pcre_uint32{pcre_uint32((int32(1) << uint64(ucp_gbLF))), pcre_uint32(int32(0)), pcre_uint32(int32(0)), pcre_uint32(((int32(1) << uint64(ucp_gbExtend)) | (int32(1) << uint64(ucp_gbSpacingMark)))), pcre_uint32((((((((((int32(1) << uint64(ucp_gbExtend)) | (int32(1) << uint64(ucp_gbPrepend))) | (int32(1) << uint64(ucp_gbSpacingMark))) | (int32(1) << uint64(ucp_gbL))) | (int32(1) << uint64(ucp_gbV))) | (int32(1) << uint64(ucp_gbT))) | (int32(1) << uint64(ucp_gbLV))) | (int32(1) << uint64(ucp_gbLVT))) | (int32(1) << uint64(ucp_gbOther)))), pcre_uint32(((int32(1) << uint64(ucp_gbExtend)) | (int32(1) << uint64(ucp_gbSpacingMark)))), pcre_uint32(((((((int32(1) << uint64(ucp_gbExtend)) | (int32(1) << uint64(ucp_gbSpacingMark))) | (int32(1) << uint64(ucp_gbL))) | (int32(1) << uint64(ucp_gbV))) | (int32(1) << uint64(ucp_gbLV))) | (int32(1) << uint64(ucp_gbLVT)))), pcre_uint32(((((int32(1) << uint64(ucp_gbExtend)) | (int32(1) << uint64(ucp_gbSpacingMark))) | (int32(1) << uint64(ucp_gbV))) | (int32(1) << uint64(ucp_gbT)))), pcre_uint32((((int32(1) << uint64(ucp_gbExtend)) | (int32(1) << uint64(ucp_gbSpacingMark))) | (int32(1) << uint64(ucp_gbT)))), pcre_uint32(((((int32(1) << uint64(ucp_gbExtend)) | (int32(1) << uint64(ucp_gbSpacingMark))) | (int32(1) << uint64(ucp_gbV))) | (int32(1) << uint64(ucp_gbT)))), pcre_uint32((((int32(1) << uint64(ucp_gbExtend)) | (int32(1) << uint64(ucp_gbSpacingMark))) | (int32(1) << uint64(ucp_gbT)))), pcre_uint32((int32(1) << uint64(ucp_gbRegionalIndicator))), pcre_uint32(((int32(1) << uint64(ucp_gbExtend)) | (int32(1) << uint64(ucp_gbSpacingMark))))}
var utt_names []byte = *(*[]byte)(unsafe.Pointer(noarch.UnsafeSliceToSlice([]byte("Any\x00Arabic\x00Armenian\x00Avestan\x00Balinese\x00Bamum\x00Bassa_Vah\x00Batak\x00Bengali\x00Bopomofo\x00Brahmi\x00Braille\x00Buginese\x00Buhid\x00C\x00Canadian_Aboriginal\x00Carian\x00Caucasian_Albanian\x00Cc\x00Cf\x00Chakma\x00Cham\x00Cherokee\x00Cn\x00Co\x00Common\x00Coptic\x00Cs\x00Cuneiform\x00Cypriot\x00Cyrillic\x00Deseret\x00Devanagari\x00Duployan\x00Egyptian_Hieroglyphs\x00Elbasan\x00Ethiopic\x00Georgian\x00Glagolitic\x00Gothic\x00Grantha\x00Greek\x00Gujarati\x00Gurmukhi\x00Han\x00Hangul\x00Hanunoo\x00Hebrew\x00Hiragana\x00Imperial_Aramaic\x00Inherited\x00Inscriptional_Pahlavi\x00Inscriptional_Parthian\x00Javanese\x00Kaithi\x00Kannada\x00Katakana\x00Kayah_Li\x00Kharoshthi\x00Khmer\x00Khojki\x00Khudawadi\x00L\x00L&\x00Lao\x00Latin\x00Lepcha\x00Limbu\x00Linear_A\x00Linear_B\x00Lisu\x00Ll\x00Lm\x00Lo\x00Lt\x00Lu\x00Lycian\x00Lydian\x00M\x00Mahajani\x00Malayalam\x00Mandaic\x00Manichaean\x00Mc\x00Me\x00Meetei_Mayek\x00Mende_Kikakui\x00Meroitic_Cursive\x00Meroitic_Hieroglyphs\x00Miao\x00Mn\x00Modi\x00Mongolian\x00Mro\x00Myanmar\x00N\x00Nabataean\x00Nd\x00New_Tai_Lue\x00Nko\x00Nl\x00No\x00Ogham\x00Ol_Chiki\x00Old_Italic\x00Old_North_Arabian\x00Old_Permic\x00Old_Persian\x00Old_South_Arabian\x00Old_Turkic\x00Oriya\x00Osmanya\x00P\x00Pahawh_Hmong\x00Palmyrene\x00Pau_Cin_Hau\x00Pc\x00Pd\x00Pe\x00Pf\x00Phags_Pa\x00Phoenician\x00Pi\x00Po\x00Ps\x00Psalter_Pahlavi\x00Rejang\x00Runic\x00S\x00Samaritan\x00Saurashtra\x00Sc\x00Sharada\x00Shavian\x00Siddham\x00Sinhala\x00Sk\x00Sm\x00So\x00Sora_Sompeng\x00Sundanese\x00Syloti_Nagri\x00Syriac\x00Tagalog\x00Tagbanwa\x00Tai_Le\x00Tai_Tham\x00Tai_Viet\x00Takri\x00Tamil\x00Telugu\x00Thaana\x00Thai\x00Tibetan\x00Tifinagh\x00Tirhuta\x00Ugaritic\x00Vai\x00Warang_Citi\x00Xan\x00Xps\x00Xsp\x00Xuc\x00Xwd\x00Yi\x00Z\x00Zl\x00Zp\x00Zs\x00\x00"), 1, 1)))
var utt []ucp_type_table = []ucp_type_table{ucp_type_table{pcre_uint16(int32(0)), pcre_uint16(int32(0)), pcre_uint16(int32(0))}, ucp_type_table{pcre_uint16(int32(4)), pcre_uint16(int32(4)), pcre_uint16(ucp_Arabic)}, ucp_type_table{pcre_uint16(int32(11)), pcre_uint16(int32(4)), pcre_uint16(ucp_Armenian)}, ucp_type_table{pcre_uint16(int32(20)), pcre_uint16(int32(4)), pcre_uint16(ucp_Avestan)}, ucp_type_table{pcre_uint16(int32(28)), pcre_uint16(int32(4)), pcre_uint16(ucp_Balinese)}, ucp_type_table{pcre_uint16(int32(37)), pcre_uint16(int32(4)), pcre_uint16(ucp_Bamum)}, ucp_type_table{pcre_uint16(int32(43)), pcre_uint16(int32(4)), pcre_uint16(ucp_Bassa_Vah)}, ucp_type_table{pcre_uint16(int32(53)), pcre_uint16(int32(4)), pcre_uint16(ucp_Batak)}, ucp_type_table{pcre_uint16(int32(59)), pcre_uint16(int32(4)), pcre_uint16(ucp_Bengali)}, ucp_type_table{pcre_uint16(int32(67)), pcre_uint16(int32(4)), pcre_uint16(ucp_Bopomofo)}, ucp_type_table{pcre_uint16(int32(76)), pcre_uint16(int32(4)), pcre_uint16(ucp_Brahmi)}, ucp_type_table{pcre_uint16(int32(83)), pcre_uint16(int32(4)), pcre_uint16(ucp_Braille)}, ucp_type_table{pcre_uint16(int32(91)), pcre_uint16(int32(4)), pcre_uint16(ucp_Buginese)}, ucp_type_table{pcre_uint16(int32(100)), pcre_uint16(int32(4)), pcre_uint16(ucp_Buhid)}, ucp_type_table{pcre_uint16(int32(106)), pcre_uint16(int32(2)), pcre_uint16(ucp_C)}, ucp_type_table{pcre_uint16(int32(108)), pcre_uint16(int32(4)), pcre_uint16(ucp_Canadian_Aboriginal)}, ucp_type_table{pcre_uint16(int32(128)), pcre_uint16(int32(4)), pcre_uint16(ucp_Carian)}, ucp_type_table{pcre_uint16(int32(135)), pcre_uint16(int32(4)), pcre_uint16(ucp_Caucasian_Albanian)}, ucp_type_table{pcre_uint16(int32(154)), pcre_uint16(int32(3)), pcre_uint16(ucp_Cc)}, ucp_type_table{pcre_uint16(int32(157)), pcre_uint16(int32(3)), pcre_uint16(ucp_Cf)}, ucp_type_table{pcre_uint16(int32(160)), pcre_uint16(int32(4)), pcre_uint16(ucp_Chakma)}, ucp_type_table{pcre_uint16(int32(167)), pcre_uint16(int32(4)), pcre_uint16(ucp_Cham)}, ucp_type_table{pcre_uint16(int32(172)), pcre_uint16(int32(4)), pcre_uint16(ucp_Cherokee)}, ucp_type_table{pcre_uint16(int32(181)), pcre_uint16(int32(3)), pcre_uint16(ucp_Cn)}, ucp_type_table{pcre_uint16(int32(184)), pcre_uint16(int32(3)), pcre_uint16(ucp_Co)}, ucp_type_table{pcre_uint16(int32(187)), pcre_uint16(int32(4)), pcre_uint16(ucp_Common)}, ucp_type_table{pcre_uint16(int32(194)), pcre_uint16(int32(4)), pcre_uint16(ucp_Coptic)}, ucp_type_table{pcre_uint16(int32(201)), pcre_uint16(int32(3)), pcre_uint16(ucp_Cs)}, ucp_type_table{pcre_uint16(int32(204)), pcre_uint16(int32(4)), pcre_uint16(ucp_Cuneiform)}, ucp_type_table{pcre_uint16(int32(214)), pcre_uint16(int32(4)), pcre_uint16(ucp_Cypriot)}, ucp_type_table{pcre_uint16(int32(222)), pcre_uint16(int32(4)), pcre_uint16(ucp_Cyrillic)}, ucp_type_table{pcre_uint16(int32(231)), pcre_uint16(int32(4)), pcre_uint16(ucp_Deseret)}, ucp_type_table{pcre_uint16(int32(239)), pcre_uint16(int32(4)), pcre_uint16(ucp_Devanagari)}, ucp_type_table{pcre_uint16(int32(250)), pcre_uint16(int32(4)), pcre_uint16(ucp_Duployan)}, ucp_type_table{pcre_uint16(int32(259)), pcre_uint16(int32(4)), pcre_uint16(ucp_Egyptian_Hieroglyphs)}, ucp_type_table{pcre_uint16(int32(280)), pcre_uint16(int32(4)), pcre_uint16(ucp_Elbasan)}, ucp_type_table{pcre_uint16(int32(288)), pcre_uint16(int32(4)), pcre_uint16(ucp_Ethiopic)}, ucp_type_table{pcre_uint16(int32(297)), pcre_uint16(int32(4)), pcre_uint16(ucp_Georgian)}, ucp_type_table{pcre_uint16(int32(306)), pcre_uint16(int32(4)), pcre_uint16(ucp_Glagolitic)}, ucp_type_table{pcre_uint16(int32(317)), pcre_uint16(int32(4)), pcre_uint16(ucp_Gothic)}, ucp_type_table{pcre_uint16(int32(324)), pcre_uint16(int32(4)), pcre_uint16(ucp_Grantha)}, ucp_type_table{pcre_uint16(int32(332)), pcre_uint16(int32(4)), pcre_uint16(ucp_Greek)}, ucp_type_table{pcre_uint16(int32(338)), pcre_uint16(int32(4)), pcre_uint16(ucp_Gujarati)}, ucp_type_table{pcre_uint16(int32(347)), pcre_uint16(int32(4)), pcre_uint16(ucp_Gurmukhi)}, ucp_type_table{pcre_uint16(int32(356)), pcre_uint16(int32(4)), pcre_uint16(ucp_Han)}, ucp_type_table{pcre_uint16(int32(360)), pcre_uint16(int32(4)), pcre_uint16(ucp_Hangul)}, ucp_type_table{pcre_uint16(int32(367)), pcre_uint16(int32(4)), pcre_uint16(ucp_Hanunoo)}, ucp_type_table{pcre_uint16(int32(375)), pcre_uint16(int32(4)), pcre_uint16(ucp_Hebrew)}, ucp_type_table{pcre_uint16(int32(382)), pcre_uint16(int32(4)), pcre_uint16(ucp_Hiragana)}, ucp_type_table{pcre_uint16(int32(391)), pcre_uint16(int32(4)), pcre_uint16(ucp_Imperial_Aramaic)}, ucp_type_table{pcre_uint16(int32(408)), pcre_uint16(int32(4)), pcre_uint16(ucp_Inherited)}, ucp_type_table{pcre_uint16(int32(418)), pcre_uint16(int32(4)), pcre_uint16(ucp_Inscriptional_Pahlavi)}, ucp_type_table{pcre_uint16(int32(440)), pcre_uint16(int32(4)), pcre_uint16(ucp_Inscriptional_Parthian)}, ucp_type_table{pcre_uint16(int32(463)), pcre_uint16(int32(4)), pcre_uint16(ucp_Javanese)}, ucp_type_table{pcre_uint16(int32(472)), pcre_uint16(int32(4)), pcre_uint16(ucp_Kaithi)}, ucp_type_table{pcre_uint16(int32(479)), pcre_uint16(int32(4)), pcre_uint16(ucp_Kannada)}, ucp_type_table{pcre_uint16(int32(487)), pcre_uint16(int32(4)), pcre_uint16(ucp_Katakana)}, ucp_type_table{pcre_uint16(int32(496)), pcre_uint16(int32(4)), pcre_uint16(ucp_Kayah_Li)}, ucp_type_table{pcre_uint16(int32(505)), pcre_uint16(int32(4)), pcre_uint16(ucp_Kharoshthi)}, ucp_type_table{pcre_uint16(int32(516)), pcre_uint16(int32(4)), pcre_uint16(ucp_Khmer)}, ucp_type_table{pcre_uint16(int32(522)), pcre_uint16(int32(4)), pcre_uint16(ucp_Khojki)}, ucp_type_table{pcre_uint16(int32(529)), pcre_uint16(int32(4)), pcre_uint16(ucp_Khudawadi)}, ucp_type_table{pcre_uint16(int32(539)), pcre_uint16(int32(2)), pcre_uint16(ucp_L)}, ucp_type_table{pcre_uint16(int32(541)), pcre_uint16(int32(1)), pcre_uint16(int32(0))}, ucp_type_table{pcre_uint16(int32(544)), pcre_uint16(int32(4)), pcre_uint16(ucp_Lao)}, ucp_type_table{pcre_uint16(int32(548)), pcre_uint16(int32(4)), pcre_uint16(ucp_Latin)}, ucp_type_table{pcre_uint16(int32(554)), pcre_uint16(int32(4)), pcre_uint16(ucp_Lepcha)}, ucp_type_table{pcre_uint16(int32(561)), pcre_uint16(int32(4)), pcre_uint16(ucp_Limbu)}, ucp_type_table{pcre_uint16(int32(567)), pcre_uint16(int32(4)), pcre_uint16(ucp_Linear_A)}, ucp_type_table{pcre_uint16(int32(576)), pcre_uint16(int32(4)), pcre_uint16(ucp_Linear_B)}, ucp_type_table{pcre_uint16(int32(585)), pcre_uint16(int32(4)), pcre_uint16(ucp_Lisu)}, ucp_type_table{pcre_uint16(int32(590)), pcre_uint16(int32(3)), pcre_uint16(ucp_Ll)}, ucp_type_table{pcre_uint16(int32(593)), pcre_uint16(int32(3)), pcre_uint16(ucp_Lm)}, ucp_type_table{pcre_uint16(int32(596)), pcre_uint16(int32(3)), pcre_uint16(ucp_Lo)}, ucp_type_table{pcre_uint16(int32(599)), pcre_uint16(int32(3)), pcre_uint16(ucp_Lt)}, ucp_type_table{pcre_uint16(int32(602)), pcre_uint16(int32(3)), pcre_uint16(ucp_Lu)}, ucp_type_table{pcre_uint16(int32(605)), pcre_uint16(int32(4)), pcre_uint16(ucp_Lycian)}, ucp_type_table{pcre_uint16(int32(612)), pcre_uint16(int32(4)), pcre_uint16(ucp_Lydian)}, ucp_type_table{pcre_uint16(int32(619)), pcre_uint16(int32(2)), pcre_uint16(ucp_M)}, ucp_type_table{pcre_uint16(int32(621)), pcre_uint16(int32(4)), pcre_uint16(ucp_Mahajani)}, ucp_type_table{pcre_uint16(int32(630)), pcre_uint16(int32(4)), pcre_uint16(ucp_Malayalam)}, ucp_type_table{pcre_uint16(int32(640)), pcre_uint16(int32(4)), pcre_uint16(ucp_Mandaic)}, ucp_type_table{pcre_uint16(int32(648)), pcre_uint16(int32(4)), pcre_uint16(ucp_Manichaean)}, ucp_type_table{pcre_uint16(int32(659)), pcre_uint16(int32(3)), pcre_uint16(ucp_Mc)}, ucp_type_table{pcre_uint16(int32(662)), pcre_uint16(int32(3)), pcre_uint16(ucp_Me)}, ucp_type_table{pcre_uint16(int32(665)), pcre_uint16(int32(4)), pcre_uint16(ucp_Meetei_Mayek)}, ucp_type_table{pcre_uint16(int32(678)), pcre_uint16(int32(4)), pcre_uint16(ucp_Mende_Kikakui)}, ucp_type_table{pcre_uint16(int32(692)), pcre_uint16(int32(4)), pcre_uint16(ucp_Meroitic_Cursive)}, ucp_type_table{pcre_uint16(int32(709)), pcre_uint16(int32(4)), pcre_uint16(ucp_Meroitic_Hieroglyphs)}, ucp_type_table{pcre_uint16(int32(730)), pcre_uint16(int32(4)), pcre_uint16(ucp_Miao)}, ucp_type_table{pcre_uint16(int32(735)), pcre_uint16(int32(3)), pcre_uint16(ucp_Mn)}, ucp_type_table{pcre_uint16(int32(738)), pcre_uint16(int32(4)), pcre_uint16(ucp_Modi)}, ucp_type_table{pcre_uint16(int32(743)), pcre_uint16(int32(4)), pcre_uint16(ucp_Mongolian)}, ucp_type_table{pcre_uint16(int32(753)), pcre_uint16(int32(4)), pcre_uint16(ucp_Mro)}, ucp_type_table{pcre_uint16(int32(757)), pcre_uint16(int32(4)), pcre_uint16(ucp_Myanmar)}, ucp_type_table{pcre_uint16(int32(765)), pcre_uint16(int32(2)), pcre_uint16(ucp_N)}, ucp_type_table{pcre_uint16(int32(767)), pcre_uint16(int32(4)), pcre_uint16(ucp_Nabataean)}, ucp_type_table{pcre_uint16(int32(777)), pcre_uint16(int32(3)), pcre_uint16(ucp_Nd)}, ucp_type_table{pcre_uint16(int32(780)), pcre_uint16(int32(4)), pcre_uint16(ucp_New_Tai_Lue)}, ucp_type_table{pcre_uint16(int32(792)), pcre_uint16(int32(4)), pcre_uint16(ucp_Nko)}, ucp_type_table{pcre_uint16(int32(796)), pcre_uint16(int32(3)), pcre_uint16(ucp_Nl)}, ucp_type_table{pcre_uint16(int32(799)), pcre_uint16(int32(3)), pcre_uint16(ucp_No)}, ucp_type_table{pcre_uint16(int32(802)), pcre_uint16(int32(4)), pcre_uint16(ucp_Ogham)}, ucp_type_table{pcre_uint16(int32(808)), pcre_uint16(int32(4)), pcre_uint16(ucp_Ol_Chiki)}, ucp_type_table{pcre_uint16(int32(817)), pcre_uint16(int32(4)), pcre_uint16(ucp_Old_Italic)}, ucp_type_table{pcre_uint16(int32(828)), pcre_uint16(int32(4)), pcre_uint16(ucp_Old_North_Arabian)}, ucp_type_table{pcre_uint16(int32(846)), pcre_uint16(int32(4)), pcre_uint16(ucp_Old_Permic)}, ucp_type_table{pcre_uint16(int32(857)), pcre_uint16(int32(4)), pcre_uint16(ucp_Old_Persian)}, ucp_type_table{pcre_uint16(int32(869)), pcre_uint16(int32(4)), pcre_uint16(ucp_Old_South_Arabian)}, ucp_type_table{pcre_uint16(int32(887)), pcre_uint16(int32(4)), pcre_uint16(ucp_Old_Turkic)}, ucp_type_table{pcre_uint16(int32(898)), pcre_uint16(int32(4)), pcre_uint16(ucp_Oriya)}, ucp_type_table{pcre_uint16(int32(904)), pcre_uint16(int32(4)), pcre_uint16(ucp_Osmanya)}, ucp_type_table{pcre_uint16(int32(912)), pcre_uint16(int32(2)), pcre_uint16(ucp_P)}, ucp_type_table{pcre_uint16(int32(914)), pcre_uint16(int32(4)), pcre_uint16(ucp_Pahawh_Hmong)}, ucp_type_table{pcre_uint16(int32(927)), pcre_uint16(int32(4)), pcre_uint16(ucp_Palmyrene)}, ucp_type_table{pcre_uint16(int32(937)), pcre_uint16(int32(4)), pcre_uint16(ucp_Pau_Cin_Hau)}, ucp_type_table{pcre_uint16(int32(949)), pcre_uint16(int32(3)), pcre_uint16(ucp_Pc)}, ucp_type_table{pcre_uint16(int32(952)), pcre_uint16(int32(3)), pcre_uint16(ucp_Pd)}, ucp_type_table{pcre_uint16(int32(955)), pcre_uint16(int32(3)), pcre_uint16(ucp_Pe)}, ucp_type_table{pcre_uint16(int32(958)), pcre_uint16(int32(3)), pcre_uint16(ucp_Pf)}, ucp_type_table{pcre_uint16(int32(961)), pcre_uint16(int32(4)), pcre_uint16(ucp_Phags_Pa)}, ucp_type_table{pcre_uint16(int32(970)), pcre_uint16(int32(4)), pcre_uint16(ucp_Phoenician)}, ucp_type_table{pcre_uint16(int32(981)), pcre_uint16(int32(3)), pcre_uint16(ucp_Pi)}, ucp_type_table{pcre_uint16(int32(984)), pcre_uint16(int32(3)), pcre_uint16(ucp_Po)}, ucp_type_table{pcre_uint16(int32(987)), pcre_uint16(int32(3)), pcre_uint16(ucp_Ps)}, ucp_type_table{pcre_uint16(int32(990)), pcre_uint16(int32(4)), pcre_uint16(ucp_Psalter_Pahlavi)}, ucp_type_table{pcre_uint16(int32(1006)), pcre_uint16(int32(4)), pcre_uint16(ucp_Rejang)}, ucp_type_table{pcre_uint16(int32(1013)), pcre_uint16(int32(4)), pcre_uint16(ucp_Runic)}, ucp_type_table{pcre_uint16(int32(1019)), pcre_uint16(int32(2)), pcre_uint16(ucp_S)}, ucp_type_table{pcre_uint16(int32(1021)), pcre_uint16(int32(4)), pcre_uint16(ucp_Samaritan)}, ucp_type_table{pcre_uint16(int32(1031)), pcre_uint16(int32(4)), pcre_uint16(ucp_Saurashtra)}, ucp_type_table{pcre_uint16(int32(1042)), pcre_uint16(int32(3)), pcre_uint16(ucp_Sc)}, ucp_type_table{pcre_uint16(int32(1045)), pcre_uint16(int32(4)), pcre_uint16(ucp_Sharada)}, ucp_type_table{pcre_uint16(int32(1053)), pcre_uint16(int32(4)), pcre_uint16(ucp_Shavian)}, ucp_type_table{pcre_uint16(int32(1061)), pcre_uint16(int32(4)), pcre_uint16(ucp_Siddham)}, ucp_type_table{pcre_uint16(int32(1069)), pcre_uint16(int32(4)), pcre_uint16(ucp_Sinhala)}, ucp_type_table{pcre_uint16(int32(1077)), pcre_uint16(int32(3)), pcre_uint16(ucp_Sk)}, ucp_type_table{pcre_uint16(int32(1080)), pcre_uint16(int32(3)), pcre_uint16(ucp_Sm)}, ucp_type_table{pcre_uint16(int32(1083)), pcre_uint16(int32(3)), pcre_uint16(ucp_So)}, ucp_type_table{pcre_uint16(int32(1086)), pcre_uint16(int32(4)), pcre_uint16(ucp_Sora_Sompeng)}, ucp_type_table{pcre_uint16(int32(1099)), pcre_uint16(int32(4)), pcre_uint16(ucp_Sundanese)}, ucp_type_table{pcre_uint16(int32(1109)), pcre_uint16(int32(4)), pcre_uint16(ucp_Syloti_Nagri)}, ucp_type_table{pcre_uint16(int32(1122)), pcre_uint16(int32(4)), pcre_uint16(ucp_Syriac)}, ucp_type_table{pcre_uint16(int32(1129)), pcre_uint16(int32(4)), pcre_uint16(ucp_Tagalog)}, ucp_type_table{pcre_uint16(int32(1137)), pcre_uint16(int32(4)), pcre_uint16(ucp_Tagbanwa)}, ucp_type_table{pcre_uint16(int32(1146)), pcre_uint16(int32(4)), pcre_uint16(ucp_Tai_Le)}, ucp_type_table{pcre_uint16(int32(1153)), pcre_uint16(int32(4)), pcre_uint16(ucp_Tai_Tham)}, ucp_type_table{pcre_uint16(int32(1162)), pcre_uint16(int32(4)), pcre_uint16(ucp_Tai_Viet)}, ucp_type_table{pcre_uint16(int32(1171)), pcre_uint16(int32(4)), pcre_uint16(ucp_Takri)}, ucp_type_table{pcre_uint16(int32(1177)), pcre_uint16(int32(4)), pcre_uint16(ucp_Tamil)}, ucp_type_table{pcre_uint16(int32(1183)), pcre_uint16(int32(4)), pcre_uint16(ucp_Telugu)}, ucp_type_table{pcre_uint16(int32(1190)), pcre_uint16(int32(4)), pcre_uint16(ucp_Thaana)}, ucp_type_table{pcre_uint16(int32(1197)), pcre_uint16(int32(4)), pcre_uint16(ucp_Thai)}, ucp_type_table{pcre_uint16(int32(1202)), pcre_uint16(int32(4)), pcre_uint16(ucp_Tibetan)}, ucp_type_table{pcre_uint16(int32(1210)), pcre_uint16(int32(4)), pcre_uint16(ucp_Tifinagh)}, ucp_type_table{pcre_uint16(int32(1219)), pcre_uint16(int32(4)), pcre_uint16(ucp_Tirhuta)}, ucp_type_table{pcre_uint16(int32(1227)), pcre_uint16(int32(4)), pcre_uint16(ucp_Ugaritic)}, ucp_type_table{pcre_uint16(int32(1236)), pcre_uint16(int32(4)), pcre_uint16(ucp_Vai)}, ucp_type_table{pcre_uint16(int32(1240)), pcre_uint16(int32(4)), pcre_uint16(ucp_Warang_Citi)}, ucp_type_table{pcre_uint16(int32(1252)), pcre_uint16(int32(5)), pcre_uint16(int32(0))}, ucp_type_table{pcre_uint16(int32(1256)), pcre_uint16(int32(7)), pcre_uint16(int32(0))}, ucp_type_table{pcre_uint16(int32(1260)), pcre_uint16(int32(6)), pcre_uint16(int32(0))}, ucp_type_table{pcre_uint16(int32(1264)), pcre_uint16(int32(10)), pcre_uint16(int32(0))}, ucp_type_table{pcre_uint16(int32(1268)), pcre_uint16(int32(8)), pcre_uint16(int32(0))}, ucp_type_table{pcre_uint16(int32(1272)), pcre_uint16(int32(4)), pcre_uint16(ucp_Yi)}, ucp_type_table{pcre_uint16(int32(1275)), pcre_uint16(int32(2)), pcre_uint16(ucp_Z)}, ucp_type_table{pcre_uint16(int32(1277)), pcre_uint16(int32(3)), pcre_uint16(ucp_Zl)}, ucp_type_table{pcre_uint16(int32(1280)), pcre_uint16(int32(3)), pcre_uint16(ucp_Zp)}, ucp_type_table{pcre_uint16(int32(1283)), pcre_uint16(int32(3)), pcre_uint16(ucp_Zs)}}
var utt_size int32 = int32((1352 / 8))
var ucd_caseless_sets []pcre_uint32 = []pcre_uint32{4294967295, pcre_uint32(int32(83)), pcre_uint32(int32(115)), pcre_uint32(int32(383)), 4294967295, pcre_uint32(int32(452)), pcre_uint32(int32(453)), pcre_uint32(int32(454)), 4294967295, pcre_uint32(int32(455)), pcre_uint32(int32(456)), pcre_uint32(int32(457)), 4294967295, pcre_uint32(int32(458)), pcre_uint32(int32(459)), pcre_uint32(int32(460)), 4294967295, pcre_uint32(int32(497)), pcre_uint32(int32(498)), pcre_uint32(int32(499)), 4294967295, pcre_uint32(int32(837)), pcre_uint32(int32(921)), pcre_uint32(int32(953)), pcre_uint32(int32(8126)), 4294967295, pcre_uint32(int32(181)), pcre_uint32(int32(924)), pcre_uint32(int32(956)), 4294967295, pcre_uint32(int32(931)), pcre_uint32(int32(962)), pcre_uint32(int32(963)), 4294967295, pcre_uint32(int32(914)), pcre_uint32(int32(946)), pcre_uint32(int32(976)), 4294967295, pcre_uint32(int32(920)), pcre_uint32(int32(952)), pcre_uint32(int32(977)), pcre_uint32(int32(1012)), 4294967295, pcre_uint32(int32(934)), pcre_uint32(int32(966)), pcre_uint32(int32(981)), 4294967295, pcre_uint32(int32(928)), pcre_uint32(int32(960)), pcre_uint32(int32(982)), 4294967295, pcre_uint32(int32(922)), pcre_uint32(int32(954)), pcre_uint32(int32(1008)), 4294967295, pcre_uint32(int32(929)), pcre_uint32(int32(961)), pcre_uint32(int32(1009)), 4294967295, pcre_uint32(int32(917)), pcre_uint32(int32(949)), pcre_uint32(int32(1013)), 4294967295, pcre_uint32(int32(7776)), pcre_uint32(int32(7777)), pcre_uint32(int32(7835)), 4294967295, pcre_uint32(int32(937)), pcre_uint32(int32(969)), pcre_uint32(int32(8486)), 4294967295, pcre_uint32(int32(75)), pcre_uint32(int32(107)), pcre_uint32(int32(8490)), 4294967295, pcre_uint32(int32(197)), pcre_uint32(int32(229)), pcre_uint32(int32(8491)), 4294967295}

const REG_ASSERT int32 = 1
const REG_BADBR int32 = 2
const REG_BADPAT int32 = 3
const REG_BADRPT int32 = 4
const REG_EBRACE int32 = 5
const REG_EBRACK int32 = 6
const REG_ECOLLATE int32 = 7
const REG_ECTYPE int32 = 8
const REG_EESCAPE int32 = 9
const REG_EMPTY int32 = 10
const REG_EPAREN int32 = 11
const REG_ERANGE int32 = 12
const REG_ESIZE int32 = 13
const REG_ESPACE int32 = 14
const REG_ESUBREG int32 = 15
const REG_INVARG int32 = 16
const REG_NOMATCH int32 = 17

type regex_t struct {
	re_pcre      unsafe.Pointer
	re_nsub      size_t
	re_erroffset size_t
}
type regoff_t int32
type regmatch_t struct {
	rm_so regoff_t
	rm_eo regoff_t
}

const PCRE8_MODE int32 = 0
const PCRE16_MODE int32 = 1
const PCRE32_MODE int32 = 2

var outfile *noarch.File
var log_store int32 = int32(0)
var callout_count int32
var callout_extra int32
var callout_fail_count int32
var callout_fail_id int32
var debug_lengths int32
var first_callout int32
var jit_was_used int32
var locale_set int32 = int32(0)
var show_malloc int32
var stack_guard_return int32
var use_utf int32
var last_callout_mark *uint8 = nil
var buffer_size int32 = int32(50000)
var buffer *pcre_uint8 = nil
var pbuffer *pcre_uint8 = nil
var pcre_mode int32 = PCRE8_MODE
var jit_study_bits []int32 = []int32{int32(1), int32(2), (int32(1) + int32(2)), int32(4), (int32(1) + int32(4)), (int32(2) + int32(4)), ((int32(1) + int32(2)) + int32(4))}
var errtexts []*byte = []*byte{nil, nil, (&[]byte("NULL argument passed\x00")[0]), (&[]byte("bad option value\x00")[0]), (&[]byte("magic number missing\x00")[0]), (&[]byte("unknown opcode - pattern overwritten?\x00")[0]), (&[]byte("no more memory\x00")[0]), nil, (&[]byte("match limit exceeded\x00")[0]), (&[]byte("callout error code\x00")[0]), nil, nil, nil, (&[]byte("not used - internal error\x00")[0]), (&[]byte("internal error - pattern overwritten?\x00")[0]), (&[]byte("bad count value\x00")[0]), (&[]byte("item unsupported for DFA matching\x00")[0]), (&[]byte("backreference condition or recursion test not supported for DFA matching\x00")[0]), (&[]byte("match limit not supported for DFA matching\x00")[0]), (&[]byte("workspace size exceeded in DFA matching\x00")[0]), (&[]byte("too much recursion for DFA matching\x00")[0]), (&[]byte("recursion limit exceeded\x00")[0]), (&[]byte("not used - internal error\x00")[0]), (&[]byte("invalid combination of newline options\x00")[0]), (&[]byte("bad offset value\x00")[0]), nil, (&[]byte("nested recursion at the same subject position\x00")[0]), (&[]byte("JIT stack limit reached\x00")[0]), (&[]byte("pattern compiled in wrong mode: 8-bit/16-bit error\x00")[0]), (&[]byte("pattern compiled with other endianness\x00")[0]), (&[]byte("invalid data in workspace for DFA restart\x00")[0]), (&[]byte("bad JIT option\x00")[0]), (&[]byte("bad length\x00")[0])}
var tables0 []pcre_uint8 = []pcre_uint8{pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(2)), pcre_uint8(int32(3)), pcre_uint8(int32(4)), pcre_uint8(int32(5)), pcre_uint8(int32(6)), pcre_uint8(int32(7)), pcre_uint8(int32(8)), pcre_uint8(int32(9)), pcre_uint8(int32(10)), pcre_uint8(int32(11)), pcre_uint8(int32(12)), pcre_uint8(int32(13)), pcre_uint8(int32(14)), pcre_uint8(int32(15)), pcre_uint8(int32(16)), pcre_uint8(int32(17)), pcre_uint8(int32(18)), pcre_uint8(int32(19)), pcre_uint8(int32(20)), pcre_uint8(int32(21)), pcre_uint8(int32(22)), pcre_uint8(int32(23)), pcre_uint8(int32(24)), pcre_uint8(int32(25)), pcre_uint8(int32(26)), pcre_uint8(int32(27)), pcre_uint8(int32(28)), pcre_uint8(int32(29)), pcre_uint8(int32(30)), pcre_uint8(int32(31)), pcre_uint8(int32(32)), pcre_uint8(int32(33)), pcre_uint8(int32(34)), pcre_uint8(int32(35)), pcre_uint8(int32(36)), pcre_uint8(int32(37)), pcre_uint8(int32(38)), pcre_uint8(int32(39)), pcre_uint8(int32(40)), pcre_uint8(int32(41)), pcre_uint8(int32(42)), pcre_uint8(int32(43)), pcre_uint8(int32(44)), pcre_uint8(int32(45)), pcre_uint8(int32(46)), pcre_uint8(int32(47)), pcre_uint8(int32(48)), pcre_uint8(int32(49)), pcre_uint8(int32(50)), pcre_uint8(int32(51)), pcre_uint8(int32(52)), pcre_uint8(int32(53)), pcre_uint8(int32(54)), pcre_uint8(int32(55)), pcre_uint8(int32(56)), pcre_uint8(int32(57)), pcre_uint8(int32(58)), pcre_uint8(int32(59)), pcre_uint8(int32(60)), pcre_uint8(int32(61)), pcre_uint8(int32(62)), pcre_uint8(int32(63)), pcre_uint8(int32(64)), pcre_uint8(int32(97)), pcre_uint8(int32(98)), pcre_uint8(int32(99)), pcre_uint8(int32(100)), pcre_uint8(int32(101)), pcre_uint8(int32(102)), pcre_uint8(int32(103)), pcre_uint8(int32(104)), pcre_uint8(int32(105)), pcre_uint8(int32(106)), pcre_uint8(int32(107)), pcre_uint8(int32(108)), pcre_uint8(int32(109)), pcre_uint8(int32(110)), pcre_uint8(int32(111)), pcre_uint8(int32(112)), pcre_uint8(int32(113)), pcre_uint8(int32(114)), pcre_uint8(int32(115)), pcre_uint8(int32(116)), pcre_uint8(int32(117)), pcre_uint8(int32(118)), pcre_uint8(int32(119)), pcre_uint8(int32(120)), pcre_uint8(int32(121)), pcre_uint8(int32(122)), pcre_uint8(int32(91)), pcre_uint8(int32(92)), pcre_uint8(int32(93)), pcre_uint8(int32(94)), pcre_uint8(int32(95)), pcre_uint8(int32(96)), pcre_uint8(int32(97)), pcre_uint8(int32(98)), pcre_uint8(int32(99)), pcre_uint8(int32(100)), pcre_uint8(int32(101)), pcre_uint8(int32(102)), pcre_uint8(int32(103)), pcre_uint8(int32(104)), pcre_uint8(int32(105)), pcre_uint8(int32(106)), pcre_uint8(int32(107)), pcre_uint8(int32(108)), pcre_uint8(int32(109)), pcre_uint8(int32(110)), pcre_uint8(int32(111)), pcre_uint8(int32(112)), pcre_uint8(int32(113)), pcre_uint8(int32(114)), pcre_uint8(int32(115)), pcre_uint8(int32(116)), pcre_uint8(int32(117)), pcre_uint8(int32(118)), pcre_uint8(int32(119)), pcre_uint8(int32(120)), pcre_uint8(int32(121)), pcre_uint8(int32(122)), pcre_uint8(int32(123)), pcre_uint8(int32(124)), pcre_uint8(int32(125)), pcre_uint8(int32(126)), pcre_uint8(int32(127)), pcre_uint8(int32(128)), pcre_uint8(int32(129)), pcre_uint8(int32(130)), pcre_uint8(int32(131)), pcre_uint8(int32(132)), pcre_uint8(int32(133)), pcre_uint8(int32(134)), pcre_uint8(int32(135)), pcre_uint8(int32(136)), pcre_uint8(int32(137)), pcre_uint8(int32(138)), pcre_uint8(int32(139)), pcre_uint8(int32(140)), pcre_uint8(int32(141)), pcre_uint8(int32(142)), pcre_uint8(int32(143)), pcre_uint8(int32(144)), pcre_uint8(int32(145)), pcre_uint8(int32(146)), pcre_uint8(int32(147)), pcre_uint8(int32(148)), pcre_uint8(int32(149)), pcre_uint8(int32(150)), pcre_uint8(int32(151)), pcre_uint8(int32(152)), pcre_uint8(int32(153)), pcre_uint8(int32(154)), pcre_uint8(int32(155)), pcre_uint8(int32(156)), pcre_uint8(int32(157)), pcre_uint8(int32(158)), pcre_uint8(int32(159)), pcre_uint8(int32(160)), pcre_uint8(int32(161)), pcre_uint8(int32(162)), pcre_uint8(int32(163)), pcre_uint8(int32(164)), pcre_uint8(int32(165)), pcre_uint8(int32(166)), pcre_uint8(int32(167)), pcre_uint8(int32(168)), pcre_uint8(int32(169)), pcre_uint8(int32(170)), pcre_uint8(int32(171)), pcre_uint8(int32(172)), pcre_uint8(int32(173)), pcre_uint8(int32(174)), pcre_uint8(int32(175)), pcre_uint8(int32(176)), pcre_uint8(int32(177)), pcre_uint8(int32(178)), pcre_uint8(int32(179)), pcre_uint8(int32(180)), pcre_uint8(int32(181)), pcre_uint8(int32(182)), pcre_uint8(int32(183)), pcre_uint8(int32(184)), pcre_uint8(int32(185)), pcre_uint8(int32(186)), pcre_uint8(int32(187)), pcre_uint8(int32(188)), pcre_uint8(int32(189)), pcre_uint8(int32(190)), pcre_uint8(int32(191)), pcre_uint8(int32(192)), pcre_uint8(int32(193)), pcre_uint8(int32(194)), pcre_uint8(int32(195)), pcre_uint8(int32(196)), pcre_uint8(int32(197)), pcre_uint8(int32(198)), pcre_uint8(int32(199)), pcre_uint8(int32(200)), pcre_uint8(int32(201)), pcre_uint8(int32(202)), pcre_uint8(int32(203)), pcre_uint8(int32(204)), pcre_uint8(int32(205)), pcre_uint8(int32(206)), pcre_uint8(int32(207)), pcre_uint8(int32(208)), pcre_uint8(int32(209)), pcre_uint8(int32(210)), pcre_uint8(int32(211)), pcre_uint8(int32(212)), pcre_uint8(int32(213)), pcre_uint8(int32(214)), pcre_uint8(int32(215)), pcre_uint8(int32(216)), pcre_uint8(int32(217)), pcre_uint8(int32(218)), pcre_uint8(int32(219)), pcre_uint8(int32(220)), pcre_uint8(int32(221)), pcre_uint8(int32(222)), pcre_uint8(int32(223)), pcre_uint8(int32(224)), pcre_uint8(int32(225)), pcre_uint8(int32(226)), pcre_uint8(int32(227)), pcre_uint8(int32(228)), pcre_uint8(int32(229)), pcre_uint8(int32(230)), pcre_uint8(int32(231)), pcre_uint8(int32(232)), pcre_uint8(int32(233)), pcre_uint8(int32(234)), pcre_uint8(int32(235)), pcre_uint8(int32(236)), pcre_uint8(int32(237)), pcre_uint8(int32(238)), pcre_uint8(int32(239)), pcre_uint8(int32(240)), pcre_uint8(int32(241)), pcre_uint8(int32(242)), pcre_uint8(int32(243)), pcre_uint8(int32(244)), pcre_uint8(int32(245)), pcre_uint8(int32(246)), pcre_uint8(int32(247)), pcre_uint8(int32(248)), pcre_uint8(int32(249)), pcre_uint8(int32(250)), pcre_uint8(int32(251)), pcre_uint8(int32(252)), pcre_uint8(int32(253)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(2)), pcre_uint8(int32(3)), pcre_uint8(int32(4)), pcre_uint8(int32(5)), pcre_uint8(int32(6)), pcre_uint8(int32(7)), pcre_uint8(int32(8)), pcre_uint8(int32(9)), pcre_uint8(int32(10)), pcre_uint8(int32(11)), pcre_uint8(int32(12)), pcre_uint8(int32(13)), pcre_uint8(int32(14)), pcre_uint8(int32(15)), pcre_uint8(int32(16)), pcre_uint8(int32(17)), pcre_uint8(int32(18)), pcre_uint8(int32(19)), pcre_uint8(int32(20)), pcre_uint8(int32(21)), pcre_uint8(int32(22)), pcre_uint8(int32(23)), pcre_uint8(int32(24)), pcre_uint8(int32(25)), pcre_uint8(int32(26)), pcre_uint8(int32(27)), pcre_uint8(int32(28)), pcre_uint8(int32(29)), pcre_uint8(int32(30)), pcre_uint8(int32(31)), pcre_uint8(int32(32)), pcre_uint8(int32(33)), pcre_uint8(int32(34)), pcre_uint8(int32(35)), pcre_uint8(int32(36)), pcre_uint8(int32(37)), pcre_uint8(int32(38)), pcre_uint8(int32(39)), pcre_uint8(int32(40)), pcre_uint8(int32(41)), pcre_uint8(int32(42)), pcre_uint8(int32(43)), pcre_uint8(int32(44)), pcre_uint8(int32(45)), pcre_uint8(int32(46)), pcre_uint8(int32(47)), pcre_uint8(int32(48)), pcre_uint8(int32(49)), pcre_uint8(int32(50)), pcre_uint8(int32(51)), pcre_uint8(int32(52)), pcre_uint8(int32(53)), pcre_uint8(int32(54)), pcre_uint8(int32(55)), pcre_uint8(int32(56)), pcre_uint8(int32(57)), pcre_uint8(int32(58)), pcre_uint8(int32(59)), pcre_uint8(int32(60)), pcre_uint8(int32(61)), pcre_uint8(int32(62)), pcre_uint8(int32(63)), pcre_uint8(int32(64)), pcre_uint8(int32(97)), pcre_uint8(int32(98)), pcre_uint8(int32(99)), pcre_uint8(int32(100)), pcre_uint8(int32(101)), pcre_uint8(int32(102)), pcre_uint8(int32(103)), pcre_uint8(int32(104)), pcre_uint8(int32(105)), pcre_uint8(int32(106)), pcre_uint8(int32(107)), pcre_uint8(int32(108)), pcre_uint8(int32(109)), pcre_uint8(int32(110)), pcre_uint8(int32(111)), pcre_uint8(int32(112)), pcre_uint8(int32(113)), pcre_uint8(int32(114)), pcre_uint8(int32(115)), pcre_uint8(int32(116)), pcre_uint8(int32(117)), pcre_uint8(int32(118)), pcre_uint8(int32(119)), pcre_uint8(int32(120)), pcre_uint8(int32(121)), pcre_uint8(int32(122)), pcre_uint8(int32(91)), pcre_uint8(int32(92)), pcre_uint8(int32(93)), pcre_uint8(int32(94)), pcre_uint8(int32(95)), pcre_uint8(int32(96)), pcre_uint8(int32(65)), pcre_uint8(int32(66)), pcre_uint8(int32(67)), pcre_uint8(int32(68)), pcre_uint8(int32(69)), pcre_uint8(int32(70)), pcre_uint8(int32(71)), pcre_uint8(int32(72)), pcre_uint8(int32(73)), pcre_uint8(int32(74)), pcre_uint8(int32(75)), pcre_uint8(int32(76)), pcre_uint8(int32(77)), pcre_uint8(int32(78)), pcre_uint8(int32(79)), pcre_uint8(int32(80)), pcre_uint8(int32(81)), pcre_uint8(int32(82)), pcre_uint8(int32(83)), pcre_uint8(int32(84)), pcre_uint8(int32(85)), pcre_uint8(int32(86)), pcre_uint8(int32(87)), pcre_uint8(int32(88)), pcre_uint8(int32(89)), pcre_uint8(int32(90)), pcre_uint8(int32(123)), pcre_uint8(int32(124)), pcre_uint8(int32(125)), pcre_uint8(int32(126)), pcre_uint8(int32(127)), pcre_uint8(int32(128)), pcre_uint8(int32(129)), pcre_uint8(int32(130)), pcre_uint8(int32(131)), pcre_uint8(int32(132)), pcre_uint8(int32(133)), pcre_uint8(int32(134)), pcre_uint8(int32(135)), pcre_uint8(int32(136)), pcre_uint8(int32(137)), pcre_uint8(int32(138)), pcre_uint8(int32(139)), pcre_uint8(int32(140)), pcre_uint8(int32(141)), pcre_uint8(int32(142)), pcre_uint8(int32(143)), pcre_uint8(int32(144)), pcre_uint8(int32(145)), pcre_uint8(int32(146)), pcre_uint8(int32(147)), pcre_uint8(int32(148)), pcre_uint8(int32(149)), pcre_uint8(int32(150)), pcre_uint8(int32(151)), pcre_uint8(int32(152)), pcre_uint8(int32(153)), pcre_uint8(int32(154)), pcre_uint8(int32(155)), pcre_uint8(int32(156)), pcre_uint8(int32(157)), pcre_uint8(int32(158)), pcre_uint8(int32(159)), pcre_uint8(int32(160)), pcre_uint8(int32(161)), pcre_uint8(int32(162)), pcre_uint8(int32(163)), pcre_uint8(int32(164)), pcre_uint8(int32(165)), pcre_uint8(int32(166)), pcre_uint8(int32(167)), pcre_uint8(int32(168)), pcre_uint8(int32(169)), pcre_uint8(int32(170)), pcre_uint8(int32(171)), pcre_uint8(int32(172)), pcre_uint8(int32(173)), pcre_uint8(int32(174)), pcre_uint8(int32(175)), pcre_uint8(int32(176)), pcre_uint8(int32(177)), pcre_uint8(int32(178)), pcre_uint8(int32(179)), pcre_uint8(int32(180)), pcre_uint8(int32(181)), pcre_uint8(int32(182)), pcre_uint8(int32(183)), pcre_uint8(int32(184)), pcre_uint8(int32(185)), pcre_uint8(int32(186)), pcre_uint8(int32(187)), pcre_uint8(int32(188)), pcre_uint8(int32(189)), pcre_uint8(int32(190)), pcre_uint8(int32(191)), pcre_uint8(int32(192)), pcre_uint8(int32(193)), pcre_uint8(int32(194)), pcre_uint8(int32(195)), pcre_uint8(int32(196)), pcre_uint8(int32(197)), pcre_uint8(int32(198)), pcre_uint8(int32(199)), pcre_uint8(int32(200)), pcre_uint8(int32(201)), pcre_uint8(int32(202)), pcre_uint8(int32(203)), pcre_uint8(int32(204)), pcre_uint8(int32(205)), pcre_uint8(int32(206)), pcre_uint8(int32(207)), pcre_uint8(int32(208)), pcre_uint8(int32(209)), pcre_uint8(int32(210)), pcre_uint8(int32(211)), pcre_uint8(int32(212)), pcre_uint8(int32(213)), pcre_uint8(int32(214)), pcre_uint8(int32(215)), pcre_uint8(int32(216)), pcre_uint8(int32(217)), pcre_uint8(int32(218)), pcre_uint8(int32(219)), pcre_uint8(int32(220)), pcre_uint8(int32(221)), pcre_uint8(int32(222)), pcre_uint8(int32(223)), pcre_uint8(int32(224)), pcre_uint8(int32(225)), pcre_uint8(int32(226)), pcre_uint8(int32(227)), pcre_uint8(int32(228)), pcre_uint8(int32(229)), pcre_uint8(int32(230)), pcre_uint8(int32(231)), pcre_uint8(int32(232)), pcre_uint8(int32(233)), pcre_uint8(int32(234)), pcre_uint8(int32(235)), pcre_uint8(int32(236)), pcre_uint8(int32(237)), pcre_uint8(int32(238)), pcre_uint8(int32(239)), pcre_uint8(int32(240)), pcre_uint8(int32(241)), pcre_uint8(int32(242)), pcre_uint8(int32(243)), pcre_uint8(int32(244)), pcre_uint8(int32(245)), pcre_uint8(int32(246)), pcre_uint8(int32(247)), pcre_uint8(int32(248)), pcre_uint8(int32(249)), pcre_uint8(int32(250)), pcre_uint8(int32(251)), pcre_uint8(int32(252)), pcre_uint8(int32(253)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(62)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(3)), pcre_uint8(int32(126)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(126)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(3)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(7)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(7)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(3)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(135)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(7)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(127)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(127)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(252)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(248)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(120)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(16)), pcre_uint8(int32(0)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0))}
var tables1 []pcre_uint8 = []pcre_uint8{pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(2)), pcre_uint8(int32(3)), pcre_uint8(int32(4)), pcre_uint8(int32(5)), pcre_uint8(int32(6)), pcre_uint8(int32(7)), pcre_uint8(int32(8)), pcre_uint8(int32(9)), pcre_uint8(int32(10)), pcre_uint8(int32(11)), pcre_uint8(int32(12)), pcre_uint8(int32(13)), pcre_uint8(int32(14)), pcre_uint8(int32(15)), pcre_uint8(int32(16)), pcre_uint8(int32(17)), pcre_uint8(int32(18)), pcre_uint8(int32(19)), pcre_uint8(int32(20)), pcre_uint8(int32(21)), pcre_uint8(int32(22)), pcre_uint8(int32(23)), pcre_uint8(int32(24)), pcre_uint8(int32(25)), pcre_uint8(int32(26)), pcre_uint8(int32(27)), pcre_uint8(int32(28)), pcre_uint8(int32(29)), pcre_uint8(int32(30)), pcre_uint8(int32(31)), pcre_uint8(int32(32)), pcre_uint8(int32(33)), pcre_uint8(int32(34)), pcre_uint8(int32(35)), pcre_uint8(int32(36)), pcre_uint8(int32(37)), pcre_uint8(int32(38)), pcre_uint8(int32(39)), pcre_uint8(int32(40)), pcre_uint8(int32(41)), pcre_uint8(int32(42)), pcre_uint8(int32(43)), pcre_uint8(int32(44)), pcre_uint8(int32(45)), pcre_uint8(int32(46)), pcre_uint8(int32(47)), pcre_uint8(int32(48)), pcre_uint8(int32(49)), pcre_uint8(int32(50)), pcre_uint8(int32(51)), pcre_uint8(int32(52)), pcre_uint8(int32(53)), pcre_uint8(int32(54)), pcre_uint8(int32(55)), pcre_uint8(int32(56)), pcre_uint8(int32(57)), pcre_uint8(int32(58)), pcre_uint8(int32(59)), pcre_uint8(int32(60)), pcre_uint8(int32(61)), pcre_uint8(int32(62)), pcre_uint8(int32(63)), pcre_uint8(int32(64)), pcre_uint8(int32(97)), pcre_uint8(int32(98)), pcre_uint8(int32(99)), pcre_uint8(int32(100)), pcre_uint8(int32(101)), pcre_uint8(int32(102)), pcre_uint8(int32(103)), pcre_uint8(int32(104)), pcre_uint8(int32(105)), pcre_uint8(int32(106)), pcre_uint8(int32(107)), pcre_uint8(int32(108)), pcre_uint8(int32(109)), pcre_uint8(int32(110)), pcre_uint8(int32(111)), pcre_uint8(int32(112)), pcre_uint8(int32(113)), pcre_uint8(int32(114)), pcre_uint8(int32(115)), pcre_uint8(int32(116)), pcre_uint8(int32(117)), pcre_uint8(int32(118)), pcre_uint8(int32(119)), pcre_uint8(int32(120)), pcre_uint8(int32(121)), pcre_uint8(int32(122)), pcre_uint8(int32(91)), pcre_uint8(int32(92)), pcre_uint8(int32(93)), pcre_uint8(int32(94)), pcre_uint8(int32(95)), pcre_uint8(int32(96)), pcre_uint8(int32(97)), pcre_uint8(int32(98)), pcre_uint8(int32(99)), pcre_uint8(int32(100)), pcre_uint8(int32(101)), pcre_uint8(int32(102)), pcre_uint8(int32(103)), pcre_uint8(int32(104)), pcre_uint8(int32(105)), pcre_uint8(int32(106)), pcre_uint8(int32(107)), pcre_uint8(int32(108)), pcre_uint8(int32(109)), pcre_uint8(int32(110)), pcre_uint8(int32(111)), pcre_uint8(int32(112)), pcre_uint8(int32(113)), pcre_uint8(int32(114)), pcre_uint8(int32(115)), pcre_uint8(int32(116)), pcre_uint8(int32(117)), pcre_uint8(int32(118)), pcre_uint8(int32(119)), pcre_uint8(int32(120)), pcre_uint8(int32(121)), pcre_uint8(int32(122)), pcre_uint8(int32(123)), pcre_uint8(int32(124)), pcre_uint8(int32(125)), pcre_uint8(int32(126)), pcre_uint8(int32(127)), pcre_uint8(int32(128)), pcre_uint8(int32(129)), pcre_uint8(int32(130)), pcre_uint8(int32(131)), pcre_uint8(int32(132)), pcre_uint8(int32(133)), pcre_uint8(int32(134)), pcre_uint8(int32(135)), pcre_uint8(int32(136)), pcre_uint8(int32(137)), pcre_uint8(int32(138)), pcre_uint8(int32(139)), pcre_uint8(int32(140)), pcre_uint8(int32(141)), pcre_uint8(int32(142)), pcre_uint8(int32(143)), pcre_uint8(int32(144)), pcre_uint8(int32(145)), pcre_uint8(int32(146)), pcre_uint8(int32(147)), pcre_uint8(int32(148)), pcre_uint8(int32(149)), pcre_uint8(int32(150)), pcre_uint8(int32(151)), pcre_uint8(int32(152)), pcre_uint8(int32(153)), pcre_uint8(int32(154)), pcre_uint8(int32(155)), pcre_uint8(int32(156)), pcre_uint8(int32(157)), pcre_uint8(int32(158)), pcre_uint8(int32(159)), pcre_uint8(int32(160)), pcre_uint8(int32(161)), pcre_uint8(int32(162)), pcre_uint8(int32(163)), pcre_uint8(int32(164)), pcre_uint8(int32(165)), pcre_uint8(int32(166)), pcre_uint8(int32(167)), pcre_uint8(int32(168)), pcre_uint8(int32(169)), pcre_uint8(int32(170)), pcre_uint8(int32(171)), pcre_uint8(int32(172)), pcre_uint8(int32(173)), pcre_uint8(int32(174)), pcre_uint8(int32(175)), pcre_uint8(int32(176)), pcre_uint8(int32(177)), pcre_uint8(int32(178)), pcre_uint8(int32(179)), pcre_uint8(int32(180)), pcre_uint8(int32(181)), pcre_uint8(int32(182)), pcre_uint8(int32(183)), pcre_uint8(int32(184)), pcre_uint8(int32(185)), pcre_uint8(int32(186)), pcre_uint8(int32(187)), pcre_uint8(int32(188)), pcre_uint8(int32(189)), pcre_uint8(int32(190)), pcre_uint8(int32(191)), pcre_uint8(int32(224)), pcre_uint8(int32(225)), pcre_uint8(int32(226)), pcre_uint8(int32(227)), pcre_uint8(int32(228)), pcre_uint8(int32(229)), pcre_uint8(int32(230)), pcre_uint8(int32(231)), pcre_uint8(int32(232)), pcre_uint8(int32(233)), pcre_uint8(int32(234)), pcre_uint8(int32(235)), pcre_uint8(int32(236)), pcre_uint8(int32(237)), pcre_uint8(int32(238)), pcre_uint8(int32(239)), pcre_uint8(int32(240)), pcre_uint8(int32(241)), pcre_uint8(int32(242)), pcre_uint8(int32(243)), pcre_uint8(int32(244)), pcre_uint8(int32(245)), pcre_uint8(int32(246)), pcre_uint8(int32(215)), pcre_uint8(int32(248)), pcre_uint8(int32(249)), pcre_uint8(int32(250)), pcre_uint8(int32(251)), pcre_uint8(int32(252)), pcre_uint8(int32(253)), pcre_uint8(int32(254)), pcre_uint8(int32(223)), pcre_uint8(int32(224)), pcre_uint8(int32(225)), pcre_uint8(int32(226)), pcre_uint8(int32(227)), pcre_uint8(int32(228)), pcre_uint8(int32(229)), pcre_uint8(int32(230)), pcre_uint8(int32(231)), pcre_uint8(int32(232)), pcre_uint8(int32(233)), pcre_uint8(int32(234)), pcre_uint8(int32(235)), pcre_uint8(int32(236)), pcre_uint8(int32(237)), pcre_uint8(int32(238)), pcre_uint8(int32(239)), pcre_uint8(int32(240)), pcre_uint8(int32(241)), pcre_uint8(int32(242)), pcre_uint8(int32(243)), pcre_uint8(int32(244)), pcre_uint8(int32(245)), pcre_uint8(int32(246)), pcre_uint8(int32(247)), pcre_uint8(int32(248)), pcre_uint8(int32(249)), pcre_uint8(int32(250)), pcre_uint8(int32(251)), pcre_uint8(int32(252)), pcre_uint8(int32(253)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(2)), pcre_uint8(int32(3)), pcre_uint8(int32(4)), pcre_uint8(int32(5)), pcre_uint8(int32(6)), pcre_uint8(int32(7)), pcre_uint8(int32(8)), pcre_uint8(int32(9)), pcre_uint8(int32(10)), pcre_uint8(int32(11)), pcre_uint8(int32(12)), pcre_uint8(int32(13)), pcre_uint8(int32(14)), pcre_uint8(int32(15)), pcre_uint8(int32(16)), pcre_uint8(int32(17)), pcre_uint8(int32(18)), pcre_uint8(int32(19)), pcre_uint8(int32(20)), pcre_uint8(int32(21)), pcre_uint8(int32(22)), pcre_uint8(int32(23)), pcre_uint8(int32(24)), pcre_uint8(int32(25)), pcre_uint8(int32(26)), pcre_uint8(int32(27)), pcre_uint8(int32(28)), pcre_uint8(int32(29)), pcre_uint8(int32(30)), pcre_uint8(int32(31)), pcre_uint8(int32(32)), pcre_uint8(int32(33)), pcre_uint8(int32(34)), pcre_uint8(int32(35)), pcre_uint8(int32(36)), pcre_uint8(int32(37)), pcre_uint8(int32(38)), pcre_uint8(int32(39)), pcre_uint8(int32(40)), pcre_uint8(int32(41)), pcre_uint8(int32(42)), pcre_uint8(int32(43)), pcre_uint8(int32(44)), pcre_uint8(int32(45)), pcre_uint8(int32(46)), pcre_uint8(int32(47)), pcre_uint8(int32(48)), pcre_uint8(int32(49)), pcre_uint8(int32(50)), pcre_uint8(int32(51)), pcre_uint8(int32(52)), pcre_uint8(int32(53)), pcre_uint8(int32(54)), pcre_uint8(int32(55)), pcre_uint8(int32(56)), pcre_uint8(int32(57)), pcre_uint8(int32(58)), pcre_uint8(int32(59)), pcre_uint8(int32(60)), pcre_uint8(int32(61)), pcre_uint8(int32(62)), pcre_uint8(int32(63)), pcre_uint8(int32(64)), pcre_uint8(int32(97)), pcre_uint8(int32(98)), pcre_uint8(int32(99)), pcre_uint8(int32(100)), pcre_uint8(int32(101)), pcre_uint8(int32(102)), pcre_uint8(int32(103)), pcre_uint8(int32(104)), pcre_uint8(int32(105)), pcre_uint8(int32(106)), pcre_uint8(int32(107)), pcre_uint8(int32(108)), pcre_uint8(int32(109)), pcre_uint8(int32(110)), pcre_uint8(int32(111)), pcre_uint8(int32(112)), pcre_uint8(int32(113)), pcre_uint8(int32(114)), pcre_uint8(int32(115)), pcre_uint8(int32(116)), pcre_uint8(int32(117)), pcre_uint8(int32(118)), pcre_uint8(int32(119)), pcre_uint8(int32(120)), pcre_uint8(int32(121)), pcre_uint8(int32(122)), pcre_uint8(int32(91)), pcre_uint8(int32(92)), pcre_uint8(int32(93)), pcre_uint8(int32(94)), pcre_uint8(int32(95)), pcre_uint8(int32(96)), pcre_uint8(int32(65)), pcre_uint8(int32(66)), pcre_uint8(int32(67)), pcre_uint8(int32(68)), pcre_uint8(int32(69)), pcre_uint8(int32(70)), pcre_uint8(int32(71)), pcre_uint8(int32(72)), pcre_uint8(int32(73)), pcre_uint8(int32(74)), pcre_uint8(int32(75)), pcre_uint8(int32(76)), pcre_uint8(int32(77)), pcre_uint8(int32(78)), pcre_uint8(int32(79)), pcre_uint8(int32(80)), pcre_uint8(int32(81)), pcre_uint8(int32(82)), pcre_uint8(int32(83)), pcre_uint8(int32(84)), pcre_uint8(int32(85)), pcre_uint8(int32(86)), pcre_uint8(int32(87)), pcre_uint8(int32(88)), pcre_uint8(int32(89)), pcre_uint8(int32(90)), pcre_uint8(int32(123)), pcre_uint8(int32(124)), pcre_uint8(int32(125)), pcre_uint8(int32(126)), pcre_uint8(int32(127)), pcre_uint8(int32(128)), pcre_uint8(int32(129)), pcre_uint8(int32(130)), pcre_uint8(int32(131)), pcre_uint8(int32(132)), pcre_uint8(int32(133)), pcre_uint8(int32(134)), pcre_uint8(int32(135)), pcre_uint8(int32(136)), pcre_uint8(int32(137)), pcre_uint8(int32(138)), pcre_uint8(int32(139)), pcre_uint8(int32(140)), pcre_uint8(int32(141)), pcre_uint8(int32(142)), pcre_uint8(int32(143)), pcre_uint8(int32(144)), pcre_uint8(int32(145)), pcre_uint8(int32(146)), pcre_uint8(int32(147)), pcre_uint8(int32(148)), pcre_uint8(int32(149)), pcre_uint8(int32(150)), pcre_uint8(int32(151)), pcre_uint8(int32(152)), pcre_uint8(int32(153)), pcre_uint8(int32(154)), pcre_uint8(int32(155)), pcre_uint8(int32(156)), pcre_uint8(int32(157)), pcre_uint8(int32(158)), pcre_uint8(int32(159)), pcre_uint8(int32(160)), pcre_uint8(int32(161)), pcre_uint8(int32(162)), pcre_uint8(int32(163)), pcre_uint8(int32(164)), pcre_uint8(int32(165)), pcre_uint8(int32(166)), pcre_uint8(int32(167)), pcre_uint8(int32(168)), pcre_uint8(int32(169)), pcre_uint8(int32(170)), pcre_uint8(int32(171)), pcre_uint8(int32(172)), pcre_uint8(int32(173)), pcre_uint8(int32(174)), pcre_uint8(int32(175)), pcre_uint8(int32(176)), pcre_uint8(int32(177)), pcre_uint8(int32(178)), pcre_uint8(int32(179)), pcre_uint8(int32(180)), pcre_uint8(int32(181)), pcre_uint8(int32(182)), pcre_uint8(int32(183)), pcre_uint8(int32(184)), pcre_uint8(int32(185)), pcre_uint8(int32(186)), pcre_uint8(int32(187)), pcre_uint8(int32(188)), pcre_uint8(int32(189)), pcre_uint8(int32(190)), pcre_uint8(int32(191)), pcre_uint8(int32(224)), pcre_uint8(int32(225)), pcre_uint8(int32(226)), pcre_uint8(int32(227)), pcre_uint8(int32(228)), pcre_uint8(int32(229)), pcre_uint8(int32(230)), pcre_uint8(int32(231)), pcre_uint8(int32(232)), pcre_uint8(int32(233)), pcre_uint8(int32(234)), pcre_uint8(int32(235)), pcre_uint8(int32(236)), pcre_uint8(int32(237)), pcre_uint8(int32(238)), pcre_uint8(int32(239)), pcre_uint8(int32(240)), pcre_uint8(int32(241)), pcre_uint8(int32(242)), pcre_uint8(int32(243)), pcre_uint8(int32(244)), pcre_uint8(int32(245)), pcre_uint8(int32(246)), pcre_uint8(int32(215)), pcre_uint8(int32(248)), pcre_uint8(int32(249)), pcre_uint8(int32(250)), pcre_uint8(int32(251)), pcre_uint8(int32(252)), pcre_uint8(int32(253)), pcre_uint8(int32(254)), pcre_uint8(int32(223)), pcre_uint8(int32(192)), pcre_uint8(int32(193)), pcre_uint8(int32(194)), pcre_uint8(int32(195)), pcre_uint8(int32(196)), pcre_uint8(int32(197)), pcre_uint8(int32(198)), pcre_uint8(int32(199)), pcre_uint8(int32(200)), pcre_uint8(int32(201)), pcre_uint8(int32(202)), pcre_uint8(int32(203)), pcre_uint8(int32(204)), pcre_uint8(int32(205)), pcre_uint8(int32(206)), pcre_uint8(int32(207)), pcre_uint8(int32(208)), pcre_uint8(int32(209)), pcre_uint8(int32(210)), pcre_uint8(int32(211)), pcre_uint8(int32(212)), pcre_uint8(int32(213)), pcre_uint8(int32(214)), pcre_uint8(int32(247)), pcre_uint8(int32(216)), pcre_uint8(int32(217)), pcre_uint8(int32(218)), pcre_uint8(int32(219)), pcre_uint8(int32(220)), pcre_uint8(int32(221)), pcre_uint8(int32(222)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(62)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(32)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(3)), pcre_uint8(int32(126)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(126)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(3)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(12)), pcre_uint8(int32(2)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(7)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(127)), pcre_uint8(int32(127)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(7)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(4)), pcre_uint8(int32(32)), pcre_uint8(int32(4)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(127)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(3)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(135)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(7)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(4)), pcre_uint8(int32(44)), pcre_uint8(int32(6)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(127)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(127)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(127)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(2)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(127)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(252)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(248)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(120)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(254)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(255)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(28)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(128)), pcre_uint8(int32(16)), pcre_uint8(int32(0)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(26)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(128)), pcre_uint8(int32(128)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(1)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(18)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(20)), pcre_uint8(int32(20)), pcre_uint8(int32(0)), pcre_uint8(int32(18)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(20)), pcre_uint8(int32(18)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(0)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(0)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(0)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18)), pcre_uint8(int32(18))}

func print_newline_config(rc int32, isc BOOL) {
	var s *byte = nil
	if int32((NotBOOL(BOOL(isc)))) != 0 {
		noarch.Printf((&[]byte("  Newline sequence is \x00")[0]))
	}
	switch rc {
	case '\r':
		{
			s = (&[]byte("CR\x00")[0])
		}
	case '\n':
		{
			s = (&[]byte("LF\x00")[0])
		}
	case (int32((('\r' << uint64(int32(8))) | '\n'))):
		{
			s = (&[]byte("CRLF\x00")[0])
		}
	case -int32(1):
		{
			s = (&[]byte("ANY\x00")[0])
		}
	case -int32(2):
		{
			s = (&[]byte("ANYCRLF\x00")[0])
		}
	default:
		{
			noarch.Printf((&[]byte("a non-standard value: 0x%04x\n\x00")[0]), rc)
			return
		}
	}
	noarch.Printf((&[]byte("%s\n\x00")[0]), s)
}

func jit_callback(arg unsafe.Pointer) *pcre_jit_stack {
	jit_was_used = int32(1)
	return (*pcre_jit_stack)(arg)
}

func utf82ord(utf8bytes *pcre_uint8, vptr *pcre_uint32) int32 {
	var c pcre_uint32 = pcre_uint32((uint32(uint8((*func() *pcre_uint8 {
		defer func() {
			utf8bytes = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(utf8bytes)) + (uintptr)(1)*unsafe.Sizeof(*utf8bytes))))
		}()
		return utf8bytes
	}())))))
	var d pcre_uint32 = c
	var i int32
	var j int32
	var s int32
	for i = -int32(1); i < int32(6); i++ {
		if (d & pcre_uint32((uint32(int32(128))))) == pcre_uint32((uint32(int32(0)))) {
			break
		}
		d <<= pcre_uint32((uint32(uint64(int32(1)))))
	}
	if i == -int32(1) {
		*vptr = c
		return int32(1)
	}
	if (i == int32(0)) || (i == int32(6)) {
		return int32(0)
	}
	s = (int32(6) * i)
	d = ((c & pcre_uint32((uint32(*((*int32)(func() unsafe.Pointer {
		tempVar := &utf8_table3[0]
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
	}())))))) << uint64(s))
	for j = int32(0); j < i; j++ {
		c = pcre_uint32((uint32(uint8((*func() *pcre_uint8 {
			defer func() {
				utf8bytes = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(utf8bytes)) + (uintptr)(1)*unsafe.Sizeof(*utf8bytes))))
			}()
			return utf8bytes
		}())))))
		if (c & pcre_uint32((uint32(int32(192))))) != pcre_uint32((uint32(int32(128)))) {
			return -(j + int32(1))
		}
		s -= int32(6)
		d |= ((c & pcre_uint32((uint32(int32(63))))) << uint64(s))
	}
	for j = int32(0); j < utf8_table1_size; j++ {
		if d <= pcre_uint32(*((*int32)(func() unsafe.Pointer {
			tempVar := &utf8_table1[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(j)*unsafe.Sizeof(*tempVar))
		}()))) {
			break
		}
	}
	if j != i {
		return -(i + int32(1))
	}
	*vptr = d
	return (i + int32(1))
}

func ord2utf8(cvalue pcre_uint32, utf8bytes *pcre_uint8) int32 {
	var i int32
	var j int32
	if cvalue > pcre_uint32((uint32(2147483647))) {
		return -int32(1)
	}
	for i = int32(0); i < utf8_table1_size; i++ {
		if cvalue <= pcre_uint32(*((*int32)(func() unsafe.Pointer {
			tempVar := &utf8_table1[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
		}()))) {
			break
		}
	}
	utf8bytes = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(utf8bytes)) + (uintptr)(i)*unsafe.Sizeof(*utf8bytes))))
	for j = i; j > int32(0); j-- {
		*func() *pcre_uint8 {
			defer func() {
				utf8bytes = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(utf8bytes)) - (uintptr)(1)*unsafe.Sizeof(*utf8bytes))))
			}()
			return utf8bytes
		}() = pcre_uint8((uint8((uint32(int32(128)) | uint32((cvalue & pcre_uint32((uint32(int32(63))))))))))
		cvalue >>= pcre_uint32((uint32(uint64(int32(6)))))
	}
	*utf8bytes = pcre_uint8((uint8((uint32(*((*int32)(func() unsafe.Pointer {
		tempVar := &utf8_table2[0]
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
	}()))) | uint32((cvalue))))))
	return (i + int32(1))
}

func extend_inputline(f *noarch.File, start *pcre_uint8, prompt *byte) (c2goDefaultReturn *pcre_uint8) {
	var here *pcre_uint8 = start
	for {
		var rlen size_t = size_t((buffer_size - int32((int64(uintptr(unsafe.Pointer(here))) - int64(uintptr(unsafe.Pointer(buffer)))))))
		if rlen > size_t((uint32(int32(1000)))) {
			var dlen int32
			{
				if int64(uintptr(unsafe.Pointer(f))) == int64(uintptr(unsafe.Pointer(stdin))) {
					noarch.Printf((&[]byte("%s\x00")[0]), prompt)
				}
				if noarch.Fgets((*byte)(unsafe.Pointer(here)), int32(uint32((size_t(rlen)))), f) == nil {
					return func() *pcre_uint8 {
						if (map[bool]int32{false: 0, true: 1}[(int64(uintptr(unsafe.Pointer(here))) == int64(uintptr(unsafe.Pointer(start))))]) != 0 {
							return nil
						} else {
							return start
						}
					}()
				}
			}
			dlen = noarch.Strlen((*byte)(unsafe.Pointer(here)))
			if (dlen > int32(0)) && (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(here)) + (uintptr)((dlen-int32(1)))*unsafe.Sizeof(*here))))))) == int32('\n')) {
				return start
			}
			here = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(here)) + (uintptr)(dlen)*unsafe.Sizeof(*here))))
		} else {
			var new_buffer_size int32 = (int32(2) * buffer_size)
			var new_buffer *pcre_uint8 = (*pcre_uint8)(noarch.Malloc(int32(uint32(new_buffer_size))))
			var new_pbuffer *pcre_uint8 = (*pcre_uint8)(noarch.Malloc(int32(uint32(new_buffer_size))))
			if (new_buffer == nil) || (new_pbuffer == nil) {
				noarch.Fprintf(stderr, (&[]byte("pcretest: malloc(%d) failed\n\x00")[0]), new_buffer_size)
				noarch.Exit(int32(1))
			}
			noarch.Memcpy(unsafe.Pointer(new_buffer), unsafe.Pointer(buffer), int32(uint32(buffer_size)))
			noarch.Memcpy(unsafe.Pointer(new_pbuffer), unsafe.Pointer(pbuffer), int32(uint32(buffer_size)))
			buffer_size = new_buffer_size
			start = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(new_buffer)) + (uintptr)(int32((int64(uintptr(unsafe.Pointer(start)))-int64(uintptr(unsafe.Pointer(buffer))))))*unsafe.Sizeof(*new_buffer))))
			here = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(new_buffer)) + (uintptr)(int32((int64(uintptr(unsafe.Pointer(here)))-int64(uintptr(unsafe.Pointer(buffer))))))*unsafe.Sizeof(*new_buffer))))
			noarch.Free(unsafe.Pointer(buffer))
			noarch.Free(unsafe.Pointer(pbuffer))
			buffer = new_buffer
			pbuffer = new_pbuffer
		}
	}
	return
}

func get_value(str *pcre_uint8, endptr **pcre_uint8) int32 {
	var result int32 = int32(0)
	for (int32(uint8((*str))) != int32(0)) && ((int32(*((*uint16)(func() unsafe.Pointer {
		tempVar := (*linux.CtypeLoc())
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*str))))))*unsafe.Sizeof(*tempVar))
	}()))) & int32(uint16(_ISspace))) != 0) {
		str = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + (uintptr)(1)*unsafe.Sizeof(*str))))
	}
	for (int32(*((*uint16)(func() unsafe.Pointer {
		tempVar := (*linux.CtypeLoc())
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*str))))))*unsafe.Sizeof(*tempVar))
	}()))) & int32(uint16(_ISdigit))) != 0 {
		result = ((result * int32(10)) + (int32(uint8((*func() *pcre_uint8 {
			defer func() {
				str = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + (uintptr)(1)*unsafe.Sizeof(*str))))
			}()
			return str
		}()))) - int32('0')))
	}
	*endptr = str
	return (result)
}

func pchar(c pcre_uint32, f *noarch.File) int32 {
	var n int32 = int32(0)
	var tempbuffer []byte = make([]byte, 16, 16)
	if (func() int32 {
		if locale_set != 0 {
			return (map[bool]int32{false: 0, true: 1}[((c < pcre_uint32((uint32(int32(256))))) && ((int32(*((*uint16)(func() unsafe.Pointer {
				tempVar := (*linux.CtypeLoc())
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((pcre_uint32((c))))))*unsafe.Sizeof(*tempVar))
			}()))) & int32(uint16(_ISprint))) != 0))])
		} else {
			return (map[bool]int32{false: 0, true: 1}[((c >= pcre_uint32((uint32(int32(32))))) && (c < pcre_uint32((uint32(int32(127))))))])
		}
	}()) != 0 {
		if f != nil {
			noarch.Fprintf(f, (&[]byte("%c\x00")[0]), pcre_uint32(c))
		}
		return int32(1)
	}
	if c < pcre_uint32((uint32(int32(256)))) {
		if use_utf != 0 {
			if f != nil {
				noarch.Fprintf(f, (&[]byte("\\x{%02x}\x00")[0]), pcre_uint32(c))
			}
			return int32(6)
		} else {
			if f != nil {
				noarch.Fprintf(f, (&[]byte("\\x%02x\x00")[0]), pcre_uint32(c))
			}
			return int32(4)
		}
	}
	if f != nil {
		n = noarch.Fprintf(f, (&[]byte("\\x{%02x}\x00")[0]), pcre_uint32(c))
	} else {
		n = noarch.Sprintf(&tempbuffer[0], (&[]byte("\\x{%02x}\x00")[0]), pcre_uint32(c))
	}
	return func() int32 {
		if n >= int32(0) {
			return n
		} else {
			return int32(0)
		}
	}()
}

func pchars(p *pcre_uint8, length int32, f *noarch.File) int32 {
	var c pcre_uint32 = pcre_uint32(int32(0))
	var yield int32 = int32(0)
	if length < int32(0) {
		length = noarch.Strlen((*byte)(unsafe.Pointer(p)))
	}
	for func() int32 {
		defer func() {
			length -= 1
		}()
		return length
	}() > int32(0) {
		if use_utf != 0 {
			var rc int32 = utf82ord(p, &c)
			if (rc > int32(0)) && (rc <= (length + int32(1))) {
				length -= (rc - int32(1))
				p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(rc)*unsafe.Sizeof(*p))))
				yield += pchar(pcre_uint32(c), f)
				continue
			}
		}
		c = pcre_uint32((uint32(uint8((*func() *pcre_uint8 {
			defer func() {
				p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
			}()
			return p
		}())))))
		yield += pchar(pcre_uint32(c), f)
	}
	return yield
}

func read_capture_name8(p *pcre_uint8, pp **pcre_uint8, re *pcre) *pcre_uint8 {
	var npp *pcre_uint8 = *pp
	for (int32(*((*uint16)(func() unsafe.Pointer {
		tempVar := (*linux.CtypeLoc())
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
	}()))) & int32(uint16(_ISalnum))) != 0 {
		*func() *pcre_uint8 {
			defer func() {
				npp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(npp)) + (uintptr)(1)*unsafe.Sizeof(*npp))))
			}()
			return npp
		}() = *func() *pcre_uint8 {
			defer func() {
				p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
			}()
			return p
		}()
	}
	*func() *pcre_uint8 {
		defer func() {
			npp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(npp)) + (uintptr)(1)*unsafe.Sizeof(*npp))))
		}()
		return npp
	}() = pcre_uint8(int32(0))
	*npp = pcre_uint8(int32(0))
	if pcre_get_stringnumber(re, (*byte)(unsafe.Pointer((*pp)))) < int32(0) {
		noarch.Fprintf(outfile, (&[]byte("no parentheses with name \"\x00")[0]))
		_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
			tempVar := (*pp)
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
		}())), -int32(1), outfile)
		noarch.Fprintf(outfile, (&[]byte("\"\n\x00")[0]))
	}
	*pp = npp
	return p
}

func stack_guard() int32 {
	return stack_guard_return
}

func callout(cb *pcre_callout_block) int32 {
	var f *noarch.File = func() *noarch.File {
		if (first_callout | callout_extra) != 0 {
			return outfile
		} else {
			return nil
		}
	}()
	var i int32
	var current_position int32
	var pre_start int32
	var post_start int32
	var subject_length int32
	if callout_extra != 0 {
		noarch.Fprintf(f, (&[]byte("Callout %d: last capture = %d\n\x00")[0]), (*cb).callout_number, (*cb).capture_last)
		if (*cb).offset_vector != nil {
			for i = int32(0); i < ((*cb).capture_top * int32(2)); i += int32(2) {
				if *((*int32)(func() unsafe.Pointer {
					tempVar := (*cb).offset_vector
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
				}())) < int32(0) {
					noarch.Fprintf(f, (&[]byte("%2d: <unset>\n\x00")[0]), (i / int32(2)))
				} else {
					noarch.Fprintf(f, (&[]byte("%2d: \x00")[0]), (i / int32(2)))
					_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := (*pcre_uint8)(unsafe.Pointer(((*cb).subject)))
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(*((*int32)(func() unsafe.Pointer {
							tempVar := (*cb).offset_vector
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
						}())))*unsafe.Sizeof(*tempVar))
					}())), (*((*int32)(func() unsafe.Pointer {
						tempVar := (*cb).offset_vector
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)((i+int32(1)))*unsafe.Sizeof(*tempVar))
					}())) - *((*int32)(func() unsafe.Pointer {
						tempVar := (*cb).offset_vector
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
					}()))), f)
					noarch.Fprintf(f, (&[]byte("\n\x00")[0]))
				}
			}
		}
	}
	if f != nil {
		noarch.Fprintf(f, (&[]byte("--->\x00")[0]))
	}
	current_position = func() int32 {
		if (map[bool]int32{false: 0, true: 1}[((*cb).current_position >= (*cb).start_match)]) != 0 {
			return (*cb).current_position
		} else {
			return (*cb).start_match
		}
	}()
	pre_start = pchars(((*pcre_uint8)(func() unsafe.Pointer {
		tempVar := (*pcre_uint8)(unsafe.Pointer(((*cb).subject)))
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
	}())), (*cb).start_match, f)
	post_start = pchars(((*pcre_uint8)(func() unsafe.Pointer {
		tempVar := (*pcre_uint8)(unsafe.Pointer(((*cb).subject)))
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)((*cb).start_match)*unsafe.Sizeof(*tempVar))
	}())), (current_position - (*cb).start_match), f)
	subject_length = pchars(((*pcre_uint8)(func() unsafe.Pointer {
		tempVar := (*pcre_uint8)(unsafe.Pointer(((*cb).subject)))
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
	}())), (*cb).subject_length, nil)
	_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
		tempVar := (*pcre_uint8)(unsafe.Pointer(((*cb).subject)))
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(current_position)*unsafe.Sizeof(*tempVar))
	}())), ((*cb).subject_length - current_position), f)
	if f != nil {
		noarch.Fprintf(f, (&[]byte("\n\x00")[0]))
	}
	if (*cb).callout_number == int32(255) {
		noarch.Fprintf(outfile, (&[]byte("%+3d \x00")[0]), (*cb).pattern_position)
		if (*cb).pattern_position > int32(99) {
			noarch.Fprintf(outfile, (&[]byte("\n    \x00")[0]))
		}
	} else {
		if callout_extra != 0 {
			noarch.Fprintf(outfile, (&[]byte("    \x00")[0]))
		} else {
			noarch.Fprintf(outfile, (&[]byte("%3d \x00")[0]), (*cb).callout_number)
		}
	}
	for i = int32(0); i < pre_start; i++ {
		noarch.Fprintf(outfile, (&[]byte(" \x00")[0]))
	}
	noarch.Fprintf(outfile, (&[]byte("^\x00")[0]))
	if post_start > int32(0) {
		for i = int32(0); i < (post_start - int32(1)); i++ {
			noarch.Fprintf(outfile, (&[]byte(" \x00")[0]))
		}
		noarch.Fprintf(outfile, (&[]byte("^\x00")[0]))
	}
	for i = int32(0); i < (((subject_length - pre_start) - post_start) + int32(4)); i++ {
		noarch.Fprintf(outfile, (&[]byte(" \x00")[0]))
	}
	noarch.Fprintf(outfile, (&[]byte("%.*s\x00")[0]), func() int32 {
		if (map[bool]int32{false: 0, true: 1}[((*cb).next_item_length == int32(0))]) != 0 {
			return int32(1)
		} else {
			return (*cb).next_item_length
		}
	}(), ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pbuffer)) + (uintptr)((*cb).pattern_position)*unsafe.Sizeof(*pbuffer)))))
	noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
	first_callout = int32(0)
	if int64(uintptr(unsafe.Pointer((*cb).mark))) != int64(uintptr(unsafe.Pointer(last_callout_mark))) {
		if (*cb).mark == nil {
			noarch.Fprintf(outfile, (&[]byte("Latest Mark: <unset>\n\x00")[0]))
		} else {
			noarch.Fprintf(outfile, (&[]byte("Latest Mark: \x00")[0]))
			_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
				tempVar := (*pcre_uint8)(unsafe.Pointer(((*cb).mark)))
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
			}())), -int32(1), outfile)
			noarch.Fputc(int32('\n'), outfile)
		}
		last_callout_mark = (*cb).mark
	}
	if (*cb).callout_data != nil {
		var callout_data int32 = *(*int32)(((*cb).callout_data))
		if callout_data != int32(0) {
			noarch.Fprintf(outfile, (&[]byte("Callout data = %d\n\x00")[0]), callout_data)
			return callout_data
		}
	}
	return func() int32 {
		if (map[bool]int32{false: 0, true: 1}[((*cb).callout_number != callout_fail_id)]) != 0 {
			return int32(0)
		} else {
			return func() int32 {
				if (map[bool]int32{false: 0, true: 1}[(func() int32 {
					callout_count += 1
					return callout_count
				}() >= callout_fail_count)]) != 0 {
					return int32(1)
				} else {
					return int32(0)
				}
			}()
		}
	}()
}

func new_malloc(size size_t) unsafe.Pointer {
	var block unsafe.Pointer = noarch.Malloc(int32(uint32((size_t(size)))))
	if show_malloc != 0 {
		noarch.Fprintf(outfile, (&[]byte("malloc       %3d %p\n\x00")[0]), int32(uint32((size_t(size)))), block)
	}
	return block
}

func new_free(block unsafe.Pointer) {
	if show_malloc != 0 {
		noarch.Fprintf(outfile, (&[]byte("free             %p\n\x00")[0]), block)
	}
	noarch.Free(block)
}

func stack_malloc(size size_t) unsafe.Pointer {
	var block unsafe.Pointer = noarch.Malloc(int32(uint32((size_t(size)))))
	if show_malloc != 0 {
		noarch.Fprintf(outfile, (&[]byte("stack_malloc %3d %p\n\x00")[0]), int32(uint32((size_t(size)))), block)
	}
	return block
}

func stack_free(block unsafe.Pointer) {
	if show_malloc != 0 {
		noarch.Fprintf(outfile, (&[]byte("stack_free       %p\n\x00")[0]), block)
	}
	noarch.Free(block)
}

func new_info(re *pcre, study *pcre_extra, option int32, ptr unsafe.Pointer) int32 {
	var rc int32
	if pcre_mode == PCRE32_MODE {
		rc = -int32(28)
	} else if pcre_mode == PCRE16_MODE {
		rc = -int32(28)
	} else {
		rc = pcre_fullinfo(re, study, option, ptr)
	}
	if (rc < int32(0)) && (rc != -int32(33)) {
		noarch.Fprintf(outfile, (&[]byte("Error %d from pcre%s_fullinfo(%d)\n\x00")[0]), rc, func() *byte {
			if pcre_mode == PCRE32_MODE {
				return (&[]byte("32\x00")[0])
			} else {
				return func() *byte {
					if pcre_mode == PCRE16_MODE {
						return (&[]byte("16\x00")[0])
					} else {
						return (&[]byte("\x00")[0])
					}
				}()
			}
		}(), option)
		if rc == -int32(28) {
			noarch.Fprintf(outfile, (&[]byte("Running in %d-bit mode but pattern was compiled in %d-bit mode\n\x00")[0]), (int32(8) * int32(1)), (uint32(int32(8)) * uint32(((*((*real_pcre)(unsafe.Pointer(re)))).flags & pcre_uint32((uint32(((int32(1) | int32(2)) | int32(4)))))))))
		}
	}
	return rc
}

func regexflip8_or_16(ere *pcre, extra *pcre_extra) {
	var re *real_pcre8_or_16 = (*real_pcre8_or_16)(unsafe.Pointer(ere))
	(*re).magic_number = pcre_uint32(1163019088)
	(*re).size = swap_uint32(pcre_uint32((*re).size))
	(*re).options = swap_uint32(pcre_uint32((*re).options))
	(*re).flags = swap_uint32(pcre_uint32((*re).flags))
	(*re).limit_match = swap_uint32(pcre_uint32((*re).limit_match))
	(*re).limit_recursion = swap_uint32(pcre_uint32((*re).limit_recursion))
	(*re).first_char = swap_uint16(pcre_uint16((*re).first_char))
	(*re).req_char = swap_uint16(pcre_uint16((*re).req_char))
	(*re).max_lookbehind = swap_uint16(pcre_uint16((*re).max_lookbehind))
	(*re).top_bracket = swap_uint16(pcre_uint16((*re).top_bracket))
	(*re).top_backref = swap_uint16(pcre_uint16((*re).top_backref))
	(*re).name_table_offset = swap_uint16(pcre_uint16((*re).name_table_offset))
	(*re).name_entry_size = swap_uint16(pcre_uint16((*re).name_entry_size))
	(*re).name_count = swap_uint16(pcre_uint16((*re).name_count))
	(*re).ref_count = swap_uint16(pcre_uint16((*re).ref_count))
	if (extra != nil) && (((*extra).flags & uint32(int32(1))) != uint32(int32(0))) {
		var rsd *pcre_study_data = (*pcre_study_data)(((*extra).study_data))
		(*rsd).size = swap_uint32(pcre_uint32((*rsd).size))
		(*rsd).flags = swap_uint32(pcre_uint32((*rsd).flags))
		(*rsd).minlength = swap_uint32(pcre_uint32((*rsd).minlength))
	}
}

func regexflip(ere *pcre, extra *pcre_extra) {
	if uint32(((*((*real_pcre)(unsafe.Pointer(ere)))).flags & pcre_uint32((uint32((int32(1) | int32(2))))))) != 0 {
		regexflip8_or_16(ere, extra)
	}
}

func check_match_limit(re *pcre, extra *pcre_extra, bptr *pcre_uint8, len int32, start_offset int32, options int32, use_offsets *int32, use_size_offsets int32, flag int32, limit *uint32, errnumber int32, msg *byte) int32 {
	var count int32
	var min int32 = int32(0)
	var mid int32 = int32(64)
	var max int32 = -int32(1)
	(*extra).flags |= uint32(flag)
	for {
		*limit = uint32(mid)
		count = pcre_exec(re, extra, (*byte)(unsafe.Pointer(bptr)), len, start_offset, options, use_offsets, use_size_offsets)
		if count == errnumber {
			min = mid
			mid = func() int32 {
				if (map[bool]int32{false: 0, true: 1}[(mid == (max - int32(1)))]) != 0 {
					return max
				} else {
					return func() int32 {
						if (map[bool]int32{false: 0, true: 1}[(max > int32(0))]) != 0 {
							return ((min + max) / int32(2))
						} else {
							return (mid * int32(2))
						}
					}()
				}
			}()
		} else if ((count >= int32(0)) || (count == -int32(1))) || (count == -int32(12)) {
			if mid == (min + int32(1)) {
				noarch.Fprintf(outfile, (&[]byte("Minimum %s limit = %d\n\x00")[0]), msg, mid)
				break
			}
			max = mid
			mid = ((min + mid) / int32(2))
		} else {
			break
		}
	}
	(*extra).flags &= ^uint32(flag)
	return count
}

func strncmpic(s *pcre_uint8, t *pcre_uint8, n int32) int32 {
	for func() int32 {
		defer func() {
			n -= 1
		}()
		return n
	}() != 0 {
		var c int32 = (linux.ToLower(int32(uint8((pcre_uint8(*func() *pcre_uint8 {
			defer func() {
				s = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + (uintptr)(1)*unsafe.Sizeof(*s))))
			}()
			return s
		}()))))) - linux.ToLower(int32(uint8((pcre_uint8(*func() *pcre_uint8 {
			defer func() {
				t = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(t)) + (uintptr)(1)*unsafe.Sizeof(*t))))
			}()
			return t
		}()))))))
		if c != 0 {
			return c
		}
	}
	return int32(0)
}

func check_mc_option(p *pcre_uint8, f *noarch.File, nl BOOL, stype *byte) int32 {
	if strncmpic(p, (*pcre_uint8)(unsafe.Pointer((&[]byte("cr>\x00")[0]))), int32(3)) == int32(0) {
		return int32(1048576)
	}
	if strncmpic(p, (*pcre_uint8)(unsafe.Pointer((&[]byte("lf>\x00")[0]))), int32(3)) == int32(0) {
		return int32(2097152)
	}
	if strncmpic(p, (*pcre_uint8)(unsafe.Pointer((&[]byte("crlf>\x00")[0]))), int32(5)) == int32(0) {
		return int32(3145728)
	}
	if strncmpic(p, (*pcre_uint8)(unsafe.Pointer((&[]byte("anycrlf>\x00")[0]))), int32(8)) == int32(0) {
		return int32(5242880)
	}
	if strncmpic(p, (*pcre_uint8)(unsafe.Pointer((&[]byte("any>\x00")[0]))), int32(4)) == int32(0) {
		return int32(4194304)
	}
	if strncmpic(p, (*pcre_uint8)(unsafe.Pointer((&[]byte("bsr_anycrlf>\x00")[0]))), int32(12)) == int32(0) {
		return int32(8388608)
	}
	if strncmpic(p, (*pcre_uint8)(unsafe.Pointer((&[]byte("bsr_unicode>\x00")[0]))), int32(12)) == int32(0) {
		return int32(16777216)
	}
	if int32((NotBOOL(BOOL(nl)))) != 0 {
		if strncmpic(p, (*pcre_uint8)(unsafe.Pointer((&[]byte("JS>\x00")[0]))), int32(3)) == int32(0) {
			return int32(33554432)
		}
	}
	noarch.Fprintf(f, (&[]byte("Unknown %s at: <%s\n\x00")[0]), stype, p)
	return int32(0)
}

func usage() {
	noarch.Printf((&[]byte("Usage:     pcretest [options] [<input file> [<output file>]]\n\n\x00")[0]))
	noarch.Printf((&[]byte("Input and output default to stdin and stdout.\n\x00")[0]))
	noarch.Printf((&[]byte("This version of pcretest is not linked with readline().\n\x00")[0]))
	noarch.Printf((&[]byte("\nOptions:\n\x00")[0]))
	noarch.Printf((&[]byte("  -b       show compiled code\n\x00")[0]))
	noarch.Printf((&[]byte("  -C       show PCRE compile-time options and exit\n\x00")[0]))
	noarch.Printf((&[]byte("  -C arg   show a specific compile-time option and exit\n\x00")[0]))
	noarch.Printf((&[]byte("           with its value if numeric (else 0). The arg can be:\n\x00")[0]))
	noarch.Printf((&[]byte("     linksize     internal link size [2, 3, 4]\n\x00")[0]))
	noarch.Printf((&[]byte("     pcre8        8 bit library support enabled [0, 1]\n\x00")[0]))
	noarch.Printf((&[]byte("     pcre16       16 bit library support enabled [0, 1]\n\x00")[0]))
	noarch.Printf((&[]byte("     pcre32       32 bit library support enabled [0, 1]\n\x00")[0]))
	noarch.Printf((&[]byte("     utf          Unicode Transformation Format supported [0, 1]\n\x00")[0]))
	noarch.Printf((&[]byte("     ucp          Unicode Properties supported [0, 1]\n\x00")[0]))
	noarch.Printf((&[]byte("     jit          Just-in-time compiler supported [0, 1]\n\x00")[0]))
	noarch.Printf((&[]byte("     newline      Newline type [CR, LF, CRLF, ANYCRLF, ANY]\n\x00")[0]))
	noarch.Printf((&[]byte("     bsr          \\R type [ANYCRLF, ANY]\n\x00")[0]))
	noarch.Printf((&[]byte("  -d       debug: show compiled code and information (-b and -i)\n\x00")[0]))
	noarch.Printf((&[]byte("  -dfa     force DFA matching for all subjects\n\x00")[0]))
	noarch.Printf((&[]byte("  -help    show usage information\n\x00")[0]))
	noarch.Printf((&[]byte("  -i       show information about compiled patterns\n  -M       find MATCH_LIMIT minimum for each subject\n  -m       output memory used information\n  -O       set PCRE_NO_AUTO_POSSESS on each pattern\n  -o <n>   set size of offsets vector to <n>\n\x00")[0]))
	noarch.Printf((&[]byte("  -p       use POSIX interface\n\x00")[0]))
	noarch.Printf((&[]byte("  -q       quiet: do not output PCRE version number at start\n\x00")[0]))
	noarch.Printf((&[]byte("  -S <n>   set stack size to <n> megabytes\n\x00")[0]))
	noarch.Printf((&[]byte("  -s       force each pattern to be studied at basic level\n  -s+      force each pattern to be studied, using JIT if available\n  -s++     ditto, verifying when JIT was actually used\n  -s+n     force each pattern to be studied, using JIT if available,\n             where 1 <= n <= 7 selects JIT options\n  -s++n    ditto, verifying when JIT was actually used\n  -t       time compilation and execution\n\x00")[0]))
	noarch.Printf((&[]byte("  -t <n>   time compilation and execution, repeating <n> times\n\x00")[0]))
	noarch.Printf((&[]byte("  -tm      time execution (matching) only\n\x00")[0]))
	noarch.Printf((&[]byte("  -tm <n>  time execution (matching) only, repeating <n> times\n\x00")[0]))
	noarch.Printf((&[]byte("  -T       same as -t, but show total times at the end\n\x00")[0]))
	noarch.Printf((&[]byte("  -TM      same as -tm, but show total time at the end\n\x00")[0]))
}

func main() {
	argc := int32(len(os.Args))
	argv__multiarray := [][]byte{}
	argv__array := []*byte{}
	for _, argvSingle := range os.Args {
		argv__multiarray = append(argv__multiarray, append([]byte(argvSingle), 0))
	}
	for _, argvSingle := range argv__multiarray {
		argv__array = append(argv__array, &argvSingle[0])
	}
	argv := *(***byte)(unsafe.Pointer(&argv__array))
	var infile *noarch.File = stdin
	var version *byte
	var options int32 = int32(0)
	var study_options int32 = int32(0)
	var default_find_match_limit int32 = int32(0)
	var default_options pcre_uint32 = pcre_uint32(int32(0))
	var op int32 = int32(1)
	var timeit int32 = int32(0)
	var timeitm int32 = int32(0)
	var showtotaltimes int32 = int32(0)
	var showinfo int32 = int32(0)
	var showstore int32 = int32(0)
	var force_study int32 = -int32(1)
	var force_study_options int32 = int32(0)
	var quiet int32 = int32(0)
	var size_offsets int32 = int32(45)
	var size_offsets_max int32
	var offsets *int32 = nil
	var debug int32 = int32(0)
	var done int32 = int32(0)
	var all_use_dfa int32 = int32(0)
	var verify_jit int32 = int32(0)
	var yield int32 = int32(0)
	var stack_size int32
	var dbuffer *pcre_uint8 = nil
	var lockout []pcre_uint8 = (&[24]pcre_uint8{pcre_uint8(int32(0))})[:]
	var dbuffer_size size_t = size_t((uint32((1 << uint64(int32(14))))))
	var total_compile_time clock_t = clock_t(int32(0))
	var total_study_time clock_t = clock_t(int32(0))
	var total_match_time clock_t = clock_t(int32(0))
	var posix int32 = int32(0)
	var dfa_workspace *int32 = nil
	var jit_stack *pcre_jit_stack = nil
	var copynames []pcre_uint32 = make([]pcre_uint32, 1024, 1024)
	var getnames []pcre_uint32 = make([]pcre_uint32, 1024, 1024)
	var copynames8 *pcre_uint8 = (*pcre_uint8)(unsafe.Pointer(&copynames[0]))
	var getnames8 *pcre_uint8 = (*pcre_uint8)(unsafe.Pointer(&getnames[0]))
	var cn8ptr *pcre_uint8
	var gn8ptr *pcre_uint8
	buffer = (*pcre_uint8)(noarch.Malloc(int32(uint32(buffer_size))))
	pbuffer = (*pcre_uint8)(noarch.Malloc(int32(uint32(buffer_size))))
	outfile = stdout
	version = pcre_version()
	for (argc > int32(1)) && (int32(**((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)(op)*unsafe.Sizeof(*argv))))) == int32('-')) {
		var endptr *pcre_uint8
		var arg *byte = *((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)(op)*unsafe.Sizeof(*argv))))
		var goto_BAD_ARG bool
	BAD_ARG_CONTAINER:
		if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-m\x00")[0])) == int32(0) {
			showstore = int32(1)
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-s\x00")[0])) == int32(0) {
			force_study = int32(0)
		} else if !goto_BAD_ARG && noarch.Strncmp(arg, (&[]byte("-s+\x00")[0]), int32(uint32(int32(3)))) == int32(0) {
			arg = ((*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(arg)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg))))
			if int32(*arg) == int32('+') {
				arg = ((*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(arg)) + (uintptr)(1)*unsafe.Sizeof(*arg))))
				verify_jit = int32(1)
			}
			force_study = int32(1)
			if int32(*arg) == int32(0) {
				force_study_options = *((*int32)(func() unsafe.Pointer {
					tempVar := &jit_study_bits[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}()))
			} else if (int32(*arg) >= int32('1')) && (int32(*arg) <= int32('7')) {
				force_study_options = *((*int32)(func() unsafe.Pointer {
					tempVar := &jit_study_bits[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)((int32(*arg)-int32('1')))*unsafe.Sizeof(*tempVar))
				}()))
			} else {
				goto_BAD_ARG = true
				goto BAD_ARG_CONTAINER
			}
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-8\x00")[0])) == int32(0) {
			pcre_mode = PCRE8_MODE
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-16\x00")[0])) == int32(0) {
			noarch.Printf((&[]byte("** This version of PCRE was built without 16-bit support\n\x00")[0]))
			noarch.Exit(int32(1))
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-32\x00")[0])) == int32(0) {
			noarch.Printf((&[]byte("** This version of PCRE was built without 32-bit support\n\x00")[0]))
			noarch.Exit(int32(1))
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-q\x00")[0])) == int32(0) {
			quiet = int32(1)
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-b\x00")[0])) == int32(0) {
			debug = int32(1)
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-i\x00")[0])) == int32(0) {
			showinfo = int32(1)
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-d\x00")[0])) == int32(0) {
			debug = int32(1)
			showinfo = debug
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-M\x00")[0])) == int32(0) {
			default_find_match_limit = int32(1)
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-O\x00")[0])) == int32(0) {
			default_options |= pcre_uint32((uint32(int32(131072))))
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-dfa\x00")[0])) == int32(0) {
			all_use_dfa = int32(1)
		} else if !goto_BAD_ARG && ((noarch.Strcmp(arg, (&[]byte("-o\x00")[0])) == int32(0)) && (argc > int32(2))) && func() bool {
			(func() int32 {
				size_offsets = get_value((*pcre_uint8)(unsafe.Pointer(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))))), &endptr)
				return size_offsets
			}())
			return int32(uint8((*endptr))) == int32(0)
		}() {
			op += 1
			argc -= 1
		} else if !goto_BAD_ARG && ((((noarch.Strcmp(arg, (&[]byte("-t\x00")[0])) == int32(0)) || (noarch.Strcmp(arg, (&[]byte("-tm\x00")[0])) == int32(0))) || (noarch.Strcmp(arg, (&[]byte("-T\x00")[0])) == int32(0))) || (noarch.Strcmp(arg, (&[]byte("-TM\x00")[0])) == int32(0))) {
			var temp int32
			var both int32 = map[bool]int32{false: 0, true: 1}[(int32(*((*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(arg)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg))))) == int32(0))]
			showtotaltimes = map[bool]int32{false: 0, true: 1}[(int32(*((*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(arg)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg))))) == int32('T'))]
			if (argc > int32(2)) && func() bool {
				temp = get_value((*pcre_uint8)(unsafe.Pointer(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))))), &endptr)
				return int32(uint8((*endptr))) == int32(0)
			}() {
				timeitm = temp
				op += 1
				argc -= 1
			} else {
				timeitm = int32(500000)
			}
			if both != 0 {
				timeit = timeitm
			}
		} else if !goto_BAD_ARG && ((noarch.Strcmp(arg, (&[]byte("-S\x00")[0])) == int32(0)) && (argc > int32(2))) && func() bool {
			(func() int32 {
				stack_size = get_value((*pcre_uint8)(unsafe.Pointer(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))))), &endptr)
				return stack_size
			}())
			return int32(uint8((*endptr))) == int32(0)
		}() {
			var rc int32
			var rlim rlimit
			getrlimit(int32(RLIMIT_STACK), &rlim)
			rlim.rlim_cur = rlim_t((__rlim_t((uint32(((stack_size * int32(1024)) * int32(1024)))))))
			rc = setrlimit(int32(RLIMIT_STACK), &rlim)
			if rc != int32(0) {
				noarch.Printf((&[]byte("PCRE: setrlimit() failed with error %d\n\x00")[0]), rc)
				noarch.Exit(int32(1))
			}
			op += 1
			argc -= 1
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-p\x00")[0])) == int32(0) {
			posix = int32(1)
		} else if !goto_BAD_ARG && noarch.Strcmp(arg, (&[]byte("-C\x00")[0])) == int32(0) {
			var rc int32
			var lrc uint32
			if argc > int32(2) {
				if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("linksize\x00")[0])) == int32(0) {
					_ = pcre_config(int32(2), unsafe.Pointer(&rc))
					noarch.Printf((&[]byte("%d\n\x00")[0]), rc)
					yield = rc
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("pcre8\x00")[0])) == int32(0) {
					noarch.Printf((&[]byte("1\n\x00")[0]))
					yield = int32(1)
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("pcre16\x00")[0])) == int32(0) {
					noarch.Printf((&[]byte("0\n\x00")[0]))
					yield = int32(0)
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("pcre32\x00")[0])) == int32(0) {
					noarch.Printf((&[]byte("0\n\x00")[0]))
					yield = int32(0)
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("utf\x00")[0])) == int32(0) {
					if pcre_mode == PCRE8_MODE {
						_ = pcre_config(int32(0), unsafe.Pointer(&rc))
					}
					noarch.Printf((&[]byte("%d\n\x00")[0]), rc)
					yield = rc
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("ucp\x00")[0])) == int32(0) {
					_ = pcre_config(int32(6), unsafe.Pointer(&rc))
					noarch.Printf((&[]byte("%d\n\x00")[0]), rc)
					yield = rc
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("jit\x00")[0])) == int32(0) {
					_ = pcre_config(int32(9), unsafe.Pointer(&rc))
					noarch.Printf((&[]byte("%d\n\x00")[0]), rc)
					yield = rc
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("newline\x00")[0])) == int32(0) {
					_ = pcre_config(int32(1), unsafe.Pointer(&rc))
					print_newline_config(rc, BOOL((int32(1))))
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("bsr\x00")[0])) == int32(0) {
					_ = pcre_config(int32(8), unsafe.Pointer(&rc))
					noarch.Printf((&[]byte("%s\n\x00")[0]), func() *byte {
						if rc != 0 {
							return (&[]byte("ANYCRLF\x00")[0])
						} else {
							return (&[]byte("ANY\x00")[0])
						}
					}())
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("ebcdic\x00")[0])) == int32(0) {
					noarch.Printf((&[]byte("0\n\x00")[0]))
				} else if noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("ebcdic-nl\x00")[0])) == int32(0) {
					noarch.Printf((&[]byte("0\n\x00")[0]))
				} else {
					noarch.Printf((&[]byte("Unknown -C option: %s\n\x00")[0]), *((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))))
				}
				goto EXIT
			}
			noarch.Printf((&[]byte("PCRE version %s\n\x00")[0]), version)
			noarch.Printf((&[]byte("Compiled with\n\x00")[0]))
			noarch.Printf((&[]byte("  8-bit support\n\x00")[0]))
			_ = pcre_config(int32(0), unsafe.Pointer(&rc))
			noarch.Printf((&[]byte("  %sUTF-8 support\n\x00")[0]), func() *byte {
				if rc != 0 {
					return (&[]byte("\x00")[0])
				} else {
					return (&[]byte("No \x00")[0])
				}
			}())
			_ = pcre_config(int32(6), unsafe.Pointer(&rc))
			noarch.Printf((&[]byte("  %sUnicode properties support\n\x00")[0]), func() *byte {
				if rc != 0 {
					return (&[]byte("\x00")[0])
				} else {
					return (&[]byte("No \x00")[0])
				}
			}())
			_ = pcre_config(int32(9), unsafe.Pointer(&rc))
			if rc != 0 {
				var arch *byte
				_ = pcre_config(int32(11), unsafe.Pointer((&arch)))
				noarch.Printf((&[]byte("  Just-in-time compiler support: %s\n\x00")[0]), arch)
			} else {
				noarch.Printf((&[]byte("  No just-in-time compiler support\n\x00")[0]))
			}
			_ = pcre_config(int32(1), unsafe.Pointer(&rc))
			print_newline_config(rc, BOOL((int32(0))))
			_ = pcre_config(int32(8), unsafe.Pointer(&rc))
			noarch.Printf((&[]byte("  \\R matches %s\n\x00")[0]), func() *byte {
				if rc != 0 {
					return (&[]byte("CR, LF, or CRLF only\x00")[0])
				} else {
					return (&[]byte("all Unicode newlines\x00")[0])
				}
			}())
			_ = pcre_config(int32(2), unsafe.Pointer(&rc))
			noarch.Printf((&[]byte("  Internal link size = %d\n\x00")[0]), rc)
			_ = pcre_config(int32(3), unsafe.Pointer(&rc))
			noarch.Printf((&[]byte("  POSIX malloc threshold = %d\n\x00")[0]), rc)
			_ = pcre_config(int32(13), unsafe.Pointer(&lrc))
			noarch.Printf((&[]byte("  Parentheses nest limit = %ld\n\x00")[0]), lrc)
			_ = pcre_config(int32(4), unsafe.Pointer(&lrc))
			noarch.Printf((&[]byte("  Default match limit = %ld\n\x00")[0]), lrc)
			_ = pcre_config(int32(7), unsafe.Pointer(&lrc))
			noarch.Printf((&[]byte("  Default recursion depth limit = %ld\n\x00")[0]), lrc)
			_ = pcre_config(int32(5), unsafe.Pointer(&rc))
			noarch.Printf((&[]byte("  Match recursion uses %s\x00")[0]), func() *byte {
				if rc != 0 {
					return (&[]byte("stack\x00")[0])
				} else {
					return (&[]byte("heap\x00")[0])
				}
			}())
			if showstore != 0 {
				stack_size = pcre_exec(nil, nil, nil, -int32(999), -int32(999), int32(0), nil, int32(0))
				noarch.Printf((&[]byte(": %sframe size = %d bytes\x00")[0]), func() *byte {
					if rc != 0 {
						return (&[]byte("approximate \x00")[0])
					} else {
						return (&[]byte("\x00")[0])
					}
				}(), -stack_size)
			}
			noarch.Printf((&[]byte("\n\x00")[0]))
			goto EXIT
		} else if !goto_BAD_ARG && ((noarch.Strcmp(arg, (&[]byte("-help\x00")[0])) == int32(0)) || (noarch.Strcmp(arg, (&[]byte("--help\x00")[0])) == int32(0))) {
			usage()
			goto EXIT
		} else {
			if goto_BAD_ARG {
				goto BAD_ARG
			}
		BAD_ARG:
			goto_BAD_ARG = false
			noarch.Printf((&[]byte("** Unknown or malformed option %s\n\x00")[0]), arg)
			usage()
			yield = int32(1)
			goto EXIT
		}
		op += 1
		argc -= 1
	}
	size_offsets_max = size_offsets
	offsets = (*int32)(noarch.Malloc(int32((uint32(size_offsets_max) * 4))))
	if offsets == nil {
		noarch.Printf((&[]byte("** Failed to get %d bytes of memory for offsets vector\n\x00")[0]), int32((uint32(size_offsets_max) * 4)))
		yield = int32(1)
		goto EXIT
	}
	if argc > int32(1) {
		infile = noarch.Fopen(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)(op)*unsafe.Sizeof(*argv)))), (&[]byte("rb\x00")[0]))
		if infile == nil {
			noarch.Printf((&[]byte("** Failed to open %s\n\x00")[0]), *((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)(op)*unsafe.Sizeof(*argv)))))
			yield = int32(1)
			goto EXIT
		}
	}
	if argc > int32(2) {
		outfile = noarch.Fopen(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))), (&[]byte("wb\x00")[0]))
		if outfile == nil {
			noarch.Printf((&[]byte("** Failed to open %s\n\x00")[0]), *((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)((op+int32(1)))*unsafe.Sizeof(*argv)))))
			yield = int32(1)
			goto EXIT
		}
	}
	pcre_malloc = new_malloc
	pcre_free = new_free
	pcre_stack_malloc = stack_malloc
	pcre_stack_free = stack_free
	if noarch.NotInt32(quiet) != 0 {
		noarch.Fprintf(outfile, (&[]byte("PCRE version %s\n\n\x00")[0]), version)
	}
	for noarch.NotInt32(done) != 0 {
		var re *pcre = nil
		var extra *pcre_extra = nil
		var preg regex_t = regex_t{(nil), size_t(int32(0)), size_t(int32(0))}
		var do_posix int32 = int32(0)
		var error *byte
		var markptr *pcre_uint8
		var p *pcre_uint8
		var pp *pcre_uint8
		var ppp *pcre_uint8
		var to_file *pcre_uint8 = nil
		var tables *pcre_uint8 = nil
		var get_options uint32
		var true_size uint32
		var true_study_size uint32 = uint32(int32(0))
		var size size_t
		var do_allcaps int32 = int32(0)
		var do_mark int32 = int32(0)
		var do_study int32 = int32(0)
		var no_force_study int32 = int32(0)
		var do_debug int32 = debug
		var do_G int32 = int32(0)
		var do_g int32 = int32(0)
		var do_showinfo int32 = showinfo
		var do_showrest int32 = int32(0)
		var do_showcaprest int32 = int32(0)
		var do_flip int32 = int32(0)
		var erroroffset int32
		var len int32
		var delimiter int32
		var poffset int32
		var dfa_matched int32 = int32(0)
		var goto_SKIP_DATA bool
		var goto_SHOW_INFO bool
		use_utf = int32(0)
		debug_lengths = int32(1)
		pcre_stack_guard = nil
		if extend_inputline(infile, buffer, (&[]byte("  re> \x00")[0])) == nil {
			break
		}
		if int64(uintptr(unsafe.Pointer(infile))) != int64(uintptr(unsafe.Pointer(stdin))) {
			noarch.Fprintf(outfile, (&[]byte("%s\x00")[0]), (*byte)(unsafe.Pointer(buffer)))
		}
		noarch.Fflush(outfile)
		p = buffer
		for (int32(*((*uint16)(func() unsafe.Pointer {
			tempVar := (*linux.CtypeLoc())
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
		}()))) & int32(uint16(_ISspace))) != 0 {
			p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
		}
		if int32(uint8((*p))) == int32(0) {
			continue
		}
		if (int32(uint8((*p))) == int32('<')) && (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(1))*unsafe.Sizeof(*p))))))) == int32(' ')) {
			p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(2))*unsafe.Sizeof(*p))))
			for (int32(*((*uint16)(func() unsafe.Pointer {
				tempVar := (*linux.CtypeLoc())
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
			}()))) & int32(uint16(_ISspace))) != 0 {
				p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
			}
			if noarch.Strncmp((*byte)(unsafe.Pointer(p)), (&[]byte("forbid \x00")[0]), int32(uint32(int32(7)))) == int32(0) {
				p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(7))*unsafe.Sizeof(*p))))
				for (int32(*((*uint16)(func() unsafe.Pointer {
					tempVar := (*linux.CtypeLoc())
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
				}()))) & int32(uint16(_ISspace))) != 0 {
					p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
				}
				pp = &lockout[0]
				for (noarch.NotInt32((int32(*((*uint16)(func() unsafe.Pointer {
					tempVar := (*linux.CtypeLoc())
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
				}()))) & int32(uint16(_ISspace)))) != 0) && (int64(uintptr(unsafe.Pointer(pp))) < int64(uintptr(unsafe.Pointer(((*pcre_uint8)(func() unsafe.Pointer {
					tempVar := ((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &lockout[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(24))*unsafe.Sizeof(*tempVar))
					}()))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) - (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}())))))) {
					*func() *pcre_uint8 {
						defer func() {
							pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
						}()
						return pp
					}() = *func() *pcre_uint8 {
						defer func() {
							p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
						}()
						return p
					}()
				}
				*pp = pcre_uint8(int32(0))
			} else {
				noarch.Printf((&[]byte("** Unrecognized special command '%s'\n\x00")[0]), p)
				yield = int32(1)
				goto EXIT
			}
			continue
		}
		if (int32(uint8((*p))) == int32('<')) && (noarch.Strchr((*byte)(unsafe.Pointer(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p))+(uintptr)(int32(1))*unsafe.Sizeof(*p)))))), int32('<')) == nil) {
			var magic pcre_uint32
			var sbuf []pcre_uint8 = make([]pcre_uint8, 8, 8)
			var f *noarch.File
			var goto_FAIL_READ bool
			p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
			if int32(uint8((*p))) == int32('!') {
				do_debug = int32(1)
				do_showinfo = int32(1)
				p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
			}
			pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(noarch.Strlen((*byte)(unsafe.Pointer(p))))*unsafe.Sizeof(*p))))
			for (int32(*((*uint16)(func() unsafe.Pointer {
				tempVar := (*linux.CtypeLoc())
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) - (uintptr)(1)*unsafe.Sizeof(*pp))))))))))*unsafe.Sizeof(*tempVar))
			}()))) & int32(uint16(_ISspace))) != 0 {
				pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) - (uintptr)(1)*unsafe.Sizeof(*pp))))
			}
			*pp = pcre_uint8(int32(0))
			f = noarch.Fopen((*byte)(unsafe.Pointer(p)), (&[]byte("rb\x00")[0]))
			if f == nil {
				noarch.Fprintf(outfile, (&[]byte("Failed to open %s: %s\n\x00")[0]), p, noarch.Strerror((*noarch.Errno())))
				continue
			}
			if noarch.Fread(unsafe.Pointer(&sbuf[0]), int32(int32(1)), int32(int32(8)), f) != int32(uint32(int32(8))) {
				goto_FAIL_READ = true
				goto FAIL_READ_CONTAINER_OUTER
			}
			true_size = uint32(((((int32(uint8((*&sbuf[0]))) << uint64(int32(24))) | (int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
				tempVar := &sbuf[0]
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
			}()))))) << uint64(int32(16)))) | (int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
				tempVar := &sbuf[0]
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
			}()))))) << uint64(int32(8)))) | int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
				tempVar := &sbuf[0]
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
			}())))))))
			true_study_size = uint32(((((int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
				tempVar := &sbuf[0]
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
			}()))))) << uint64(int32(24))) | (int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
				tempVar := &sbuf[0]
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
			}()))))) << uint64(int32(16)))) | (int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
				tempVar := &sbuf[0]
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
			}()))))) << uint64(int32(8)))) | int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
				tempVar := &sbuf[0]
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
			}())))))))
			re = (*pcre)(new_malloc(size_t((true_size))))
			if re == nil {
				noarch.Printf((&[]byte("** Failed to get %d bytes of memory for pcre object\n\x00")[0]), int32(true_size))
				yield = int32(1)
				goto EXIT
			}
			if noarch.Fread(unsafe.Pointer(re), int32(int32(1)), int32(true_size), f) != int32(true_size) {
				goto_FAIL_READ = true
				goto FAIL_READ_CONTAINER_OUTER
			}
			magic = (*((*real_pcre)(unsafe.Pointer(re)))).magic_number
			if uint32((magic)) != uint32(1346589253) {
				if uint32((swap_uint32(pcre_uint32(magic)))) == uint32(1346589253) {
					do_flip = int32(1)
				} else {
					noarch.Fprintf(outfile, (&[]byte("Data in %s is not a compiled PCRE regex\n\x00")[0]), p)
					new_free(unsafe.Pointer(re))
					noarch.Fclose(f)
					continue
				}
			}
			noarch.Fprintf(outfile, (&[]byte("Compiled pattern%s loaded from %s\n\x00")[0]), func() *byte {
				if (do_flip != 0) && (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) - (uintptr)(1)*unsafe.Sizeof(*p))))))) == int32('<')) {
					return (&[]byte(" (byte-inverted)\x00")[0])
				} else {
					return (&[]byte("\x00")[0])
				}
			}(), p)
		FAIL_READ_CONTAINER_OUTER:
			if goto_FAIL_READ || true_study_size != uint32(int32(0)) {
				var psd *pcre_study_data
				if goto_FAIL_READ {
					goto FAIL_READ_CONTAINER
				}
				extra = (*pcre_extra)(new_malloc(size_t((64 + true_study_size))))
				(*extra).flags = uint32(int32(1))
				psd = (*pcre_study_data)(unsafe.Pointer(((*byte)(func() unsafe.Pointer {
					tempVar := (*byte)(unsafe.Pointer(extra))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(64))*unsafe.Sizeof(*tempVar))
				}()))))
				(*extra).study_data = unsafe.Pointer(psd)
			FAIL_READ_CONTAINER:
				if goto_FAIL_READ || noarch.Fread(unsafe.Pointer(psd), int32(int32(1)), int32(true_study_size), f) != int32(true_study_size) {
					if goto_FAIL_READ {
						goto FAIL_READ
					}
				FAIL_READ:
					goto_FAIL_READ = false
					noarch.Fprintf(outfile, (&[]byte("Failed to read data from %s\n\x00")[0]), p)

					if extra != nil {
						pcre_free_study(extra)
					}
					new_free(unsafe.Pointer(re))
					noarch.Fclose(f)
					continue
				}
				noarch.Fprintf(outfile, (&[]byte("Study data loaded from %s\n\x00")[0]), p)
				do_study = int32(1)
			} else {
				noarch.Fprintf(outfile, (&[]byte("No study data\n\x00")[0]))
			}
			if do_flip != 0 {
				var rc int32
				rc = pcre_pattern_to_host_byte_order(re, extra, nil)
				if rc == -int32(28) {
					var flags_in_host_byte_order pcre_uint32
					if uint32(((*((*real_pcre)(unsafe.Pointer(re)))).magic_number)) == uint32(1346589253) {
						flags_in_host_byte_order = (*((*real_pcre)(unsafe.Pointer(re)))).flags
					} else {
						flags_in_host_byte_order = swap_uint32(pcre_uint32(((*((*real_pcre)(unsafe.Pointer(re)))).flags)))
					}
					noarch.Fprintf(outfile, (&[]byte("Error %d from pcre%s_fullinfo(%d)\n\x00")[0]), rc, func() *byte {
						if pcre_mode == PCRE32_MODE {
							return (&[]byte("32\x00")[0])
						} else {
							return func() *byte {
								if pcre_mode == PCRE16_MODE {
									return (&[]byte("16\x00")[0])
								} else {
									return (&[]byte("\x00")[0])
								}
							}()
						}
					}(), int32(0))
					noarch.Fprintf(outfile, (&[]byte("Running in %d-bit mode but pattern was compiled in %d-bit mode\n\x00")[0]), (int32(8) * int32(1)), (uint32(int32(8)) * uint32((flags_in_host_byte_order & pcre_uint32((uint32(((int32(1) | int32(2)) | int32(4)))))))))
					new_free(unsafe.Pointer(re))
					noarch.Fclose(f)
					continue
				}
			}
			if new_info(re, nil, int32(0), unsafe.Pointer(&get_options)) < int32(0) {
				new_free(unsafe.Pointer(re))
				noarch.Fclose(f)
				continue
			}
			use_utf = map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(2048))) != uint32(int32(0)))]
			noarch.Fclose(f)
			goto_SHOW_INFO = true
			goto SHOW_INFO_CONTAINER
		}
		delimiter = int32(uint8((*func() *pcre_uint8 {
			defer func() {
				p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
			}()
			return p
		}())))
		if ((int32(*((*uint16)(func() unsafe.Pointer {
			tempVar := (*linux.CtypeLoc())
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)((delimiter))*unsafe.Sizeof(*tempVar))
		}()))) & int32(uint16(_ISalnum))) != 0) || (delimiter == int32('\\')) {
			noarch.Fprintf(outfile, (&[]byte("** Delimiter must not be alphanumeric or \\\n\x00")[0]))
			goto_SKIP_DATA = true
			goto SKIP_DATA_CONTAINER_OUTER
		}
		pp = p
		poffset = (int32((int64(uintptr(unsafe.Pointer(p))) - int64(uintptr(unsafe.Pointer(buffer))))))
		for {
			for int32(uint8((*pp))) != int32(0) {
				if (int32(uint8((*pp))) == int32('\\')) && (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(int32(1))*unsafe.Sizeof(*pp))))))) != int32(0)) {
					pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
				} else if int32(uint8((*pp))) == delimiter {
					break
				}
				pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
			}
			if int32(uint8((*pp))) != int32(0) {
				break
			}
			if (func() *pcre_uint8 {
				tempVar := extend_inputline(infile, pp, (&[]byte("    > \x00")[0]))
				pp = tempVar
				return tempVar
			}()) == nil {
				noarch.Fprintf(outfile, (&[]byte("** Unexpected EOF\n\x00")[0]))
				done = int32(1)
				goto CONTINUE
			}
			if int64(uintptr(unsafe.Pointer(infile))) != int64(uintptr(unsafe.Pointer(stdin))) {
				noarch.Fprintf(outfile, (&[]byte("%s\x00")[0]), (*byte)(unsafe.Pointer(pp)))
			}
		}
		p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buffer)) + (uintptr)(poffset)*unsafe.Sizeof(*buffer))))
		if int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(int32(1))*unsafe.Sizeof(*pp))))))) == int32('\\') {
			*func() *pcre_uint8 {
				defer func() {
					pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
				}()
				return pp
			}() = pcre_uint8((uint8('\\')))
		}
		*func() *pcre_uint8 {
			defer func() {
				pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
			}()
			return pp
		}() = pcre_uint8(int32(0))
		noarch.Strcpy((*byte)(unsafe.Pointer(pbuffer)), (*byte)(unsafe.Pointer(p)))
		options = int32(uint32((default_options)))
		study_options = force_study_options
		log_store = showstore
		for int32(uint8((*pp))) != int32(0) {
			if noarch.Strchr((*byte)(unsafe.Pointer(&lockout[0])), int32(uint8((pcre_uint8(*pp))))) != nil {
				if (int32(uint8((*pp))) == int32('<')) && (noarch.Strchr((*byte)(unsafe.Pointer(&lockout[0])), int32('>')) != nil) {
					var x int32 = check_mc_option(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(int32(1))*unsafe.Sizeof(*pp)))), outfile, BOOL((int32(0))), (&[]byte("modifier\x00")[0]))
					if x == int32(0) {
						goto_SKIP_DATA = true
						goto SKIP_DATA_CONTAINER_OUTER
					}
					for ppp = &lockout[0]; int32(uint8((*ppp))) != int32(0); ppp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ppp)) + (uintptr)(1)*unsafe.Sizeof(*ppp)))) {
						if int32(uint8((*ppp))) == int32('<') {
							var y int32 = check_mc_option(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ppp)) + (uintptr)(int32(1))*unsafe.Sizeof(*ppp)))), outfile, BOOL((int32(0))), (&[]byte("modifier\x00")[0]))
							if y == int32(0) {
								noarch.Printf((&[]byte("** Error in modifier forbid data - giving up.\n\x00")[0]))
								yield = int32(1)
								goto EXIT
							}
							if x == y {
								ppp = pp
								for int32(uint8((*ppp))) != int32('>') {
									ppp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ppp)) + (uintptr)(1)*unsafe.Sizeof(*ppp))))
								}
								noarch.Printf((&[]byte("** The %.*s modifier is locked out - giving up.\n\x00")[0]), (int32(((int64(uintptr(unsafe.Pointer(ppp))) - int64(uintptr(unsafe.Pointer(pp)))) + int64(int32(1))))), pp)
								yield = int32(1)
								goto EXIT
							}
						}
					}
				} else {
					noarch.Printf((&[]byte("** The /%c modifier is locked out - giving up.\n\x00")[0]), int32(uint8((pcre_uint8(*pp)))))
					yield = int32(1)
					goto EXIT
				}
			}
			switch int32(uint8((pcre_uint8(*func() *pcre_uint8 {
				defer func() {
					pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
				}()
				return pp
			}())))) {
			case 'f':
				goto SW_GENERATED_LABEL_1
			case 'g':
				goto SW_GENERATED_LABEL_2
			case 'i':
				goto SW_GENERATED_LABEL_3
			case 'm':
				goto SW_GENERATED_LABEL_4
			case 's':
				goto SW_GENERATED_LABEL_5
			case 'x':
				goto SW_GENERATED_LABEL_6
			case '+':
				goto SW_GENERATED_LABEL_7
			case '=':
				goto SW_GENERATED_LABEL_8
			case 'A':
				goto SW_GENERATED_LABEL_9
			case 'B':
				goto SW_GENERATED_LABEL_10
			case 'C':
				goto SW_GENERATED_LABEL_11
			case 'D':
				goto SW_GENERATED_LABEL_12
			case 'E':
				goto SW_GENERATED_LABEL_13
			case 'F':
				goto SW_GENERATED_LABEL_14
			case 'G':
				goto SW_GENERATED_LABEL_15
			case 'I':
				goto SW_GENERATED_LABEL_16
			case 'J':
				goto SW_GENERATED_LABEL_17
			case 'K':
				goto SW_GENERATED_LABEL_18
			case 'M':
				goto SW_GENERATED_LABEL_19
			case 'N':
				goto SW_GENERATED_LABEL_20
			case 'O':
				goto SW_GENERATED_LABEL_21
			case 'P':
				goto SW_GENERATED_LABEL_22
			case 'Q':
				goto SW_GENERATED_LABEL_23
			case 'S':
				goto SW_GENERATED_LABEL_24
			case 'U':
				goto SW_GENERATED_LABEL_25
			case 'W':
				goto SW_GENERATED_LABEL_26
			case 'X':
				goto SW_GENERATED_LABEL_27
			case 'Y':
				goto SW_GENERATED_LABEL_28
			case 'Z':
				goto SW_GENERATED_LABEL_29
			case '8':
				goto SW_GENERATED_LABEL_30
			case '9':
				goto SW_GENERATED_LABEL_31
			case '?':
				goto SW_GENERATED_LABEL_32
			case 'T':
				goto SW_GENERATED_LABEL_33
			case 'L':
				goto SW_GENERATED_LABEL_34
			case '>':
				goto SW_GENERATED_LABEL_35
			case '<':
				goto SW_GENERATED_LABEL_36
			case '\r':
				goto SW_GENERATED_LABEL_37
			case '\n':
				goto SW_GENERATED_LABEL_38
			case ' ':
				goto SW_GENERATED_LABEL_39
			default:
				goto SW_GENERATED_LABEL_40
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_1:
			{
				options |= int32(262144)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_2:
			{
				do_g = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_3:
			{
				options |= int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_4:
			{
				options |= int32(2)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_5:
			{
				options |= int32(4)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_6:
			{
				options |= int32(8)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_7:
			{
				if do_showrest != 0 {
					do_showcaprest = int32(1)
				} else {
					do_showrest = int32(1)
				}
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_8:
			{
				do_allcaps = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_9:
			{
				options |= int32(16)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_10:
			{
				do_debug = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_11:
			{
				options |= int32(16384)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_12:
			{
				do_showinfo = int32(1)
				do_debug = do_showinfo
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_13:
			{
				options |= int32(32)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_14:
			{
				do_flip = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_15:
			{
				do_G = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_16:
			{
				do_showinfo = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_17:
			{
				options |= int32(524288)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_18:
			{
				do_mark = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_19:
			{
				log_store = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_20:
			{
				options |= int32(4096)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_21:
			{
				options |= int32(131072)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_22:
			{
				do_posix = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_23:
			{
				switch int32(uint8((pcre_uint8(*pp)))) {
				case '0':
					fallthrough
				case '1':
					{
						stack_guard_return = (int32(uint8((*func() *pcre_uint8 {
							defer func() {
								pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
							}()
							return pp
						}()))) - int32('0'))
					}
				default:
					{
						noarch.Fprintf(outfile, (&[]byte("** Missing 0 or 1 after /Q\n\x00")[0]))
						goto_SKIP_DATA = true
						goto SKIP_DATA_CONTAINER_OUTER
					}
				}
				pcre_stack_guard = stack_guard
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_24:
			{
				do_study = int32(1)
				for {
					switch int32(uint8((pcre_uint8(*func() *pcre_uint8 {
						defer func() {
							pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
						}()
						return pp
					}())))) {
					case 'S':
						{
							do_study = int32(0)
							no_force_study = int32(1)
						}
					case '!':
						{
							study_options |= int32(8)
						}
					case '+':
						{
							if int32(uint8((*pp))) == int32('+') {
								verify_jit = int32(1)
								pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
							}
							if (int32(uint8((*pp))) >= int32('1')) && (int32(uint8((*pp))) <= int32('7')) {
								study_options |= *((*int32)(func() unsafe.Pointer {
									tempVar := &jit_study_bits[0]
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)((int32(uint8((*func() *pcre_uint8 {
										defer func() {
											pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
										}()
										return pp
									}())))-int32('1')))*unsafe.Sizeof(*tempVar))
								}()))
							} else {
								study_options |= *((*int32)(func() unsafe.Pointer {
									tempVar := &jit_study_bits[0]
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
								}()))
							}
						}
					case '-':
						{
							study_options &= ^((int32(1) | int32(2)) | int32(4))
						}
					default:
						{
							pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) - (uintptr)(1)*unsafe.Sizeof(*pp))))
							goto ENDLOOP
						}
					}
				}
			}
		ENDLOOP:
			{
				goto SW_GENERATED_LABEL_0
			}
		SW_GENERATED_LABEL_25:
			{
				options |= int32(512)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_26:
			{
				options |= int32(536870912)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_27:
			{
				options |= int32(64)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_28:
			{
				options |= int32(67108864)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_29:
			{
				debug_lengths = int32(0)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_30:
			{
				options |= int32(2048)
				use_utf = int32(1)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_31:
			{
				options |= int32(65536)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_32:
			{
				options |= int32(8192)
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_33:
			{
				switch int32(uint8((pcre_uint8(*func() *pcre_uint8 {
					defer func() {
						pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
					}()
					return pp
				}())))) {
				case '0':
					{
						tables = &tables0[0]
					}
				case '1':
					{
						tables = &tables1[0]
					}
				case '\r':
					fallthrough
				case '\n':
					fallthrough
				case ' ':
					fallthrough
				case int32(0):
					{
						noarch.Fprintf(outfile, (&[]byte("** Missing table number after /T\n\x00")[0]))
						goto_SKIP_DATA = true
						goto SKIP_DATA_CONTAINER_OUTER
					}
					fallthrough
				default:
					{
						noarch.Fprintf(outfile, (&[]byte("** Bad table number \"%c\" after /T\n\x00")[0]), int32(uint8((pcre_uint8(*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) - (uintptr)(1)*unsafe.Sizeof(*pp)))))))))
						goto_SKIP_DATA = true
						goto SKIP_DATA_CONTAINER_OUTER
					}
				}
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_34:
			{
				ppp = pp
				for (((int32(uint8((*ppp))) != int32(0)) && (int32(uint8((*ppp))) != int32('\n'))) && (int32(uint8((*ppp))) != int32('\r'))) && (int32(uint8((*ppp))) != int32(' ')) {
					ppp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ppp)) + (uintptr)(1)*unsafe.Sizeof(*ppp))))
				}
				*ppp = pcre_uint8(int32(0))
				if setlocale(int32(0), (*byte)(unsafe.Pointer(pp))) == nil {
					noarch.Fprintf(outfile, (&[]byte("** Failed to set locale \"%s\"\n\x00")[0]), pp)
					goto_SKIP_DATA = true
					goto SKIP_DATA_CONTAINER_OUTER
				}
				locale_set = int32(1)
				tables = (*pcre_uint8)(unsafe.Pointer(pcre_maketables()))
				pp = ppp
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_35:
			{
				to_file = pp
				for int32(uint8((*pp))) != int32(0) {
					pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
				}
				for (int32(*((*uint16)(func() unsafe.Pointer {
					tempVar := (*linux.CtypeLoc())
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) - (uintptr)(1)*unsafe.Sizeof(*pp))))))))))*unsafe.Sizeof(*tempVar))
				}()))) & int32(uint16(_ISspace))) != 0 {
					pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) - (uintptr)(1)*unsafe.Sizeof(*pp))))
				}
				*pp = pcre_uint8(int32(0))
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_36:
			{
				var x int32 = check_mc_option(pp, outfile, BOOL((int32(0))), (&[]byte("modifier\x00")[0]))
				if x == int32(0) {
					goto_SKIP_DATA = true
					goto SKIP_DATA_CONTAINER_OUTER
				}
				options |= x
				for int32(uint8((*func() *pcre_uint8 {
					defer func() {
						pp = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + (uintptr)(1)*unsafe.Sizeof(*pp))))
					}()
					return pp
				}()))) != int32('>') {
				}
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_37:
			;
		SW_GENERATED_LABEL_38:
			;
		SW_GENERATED_LABEL_39:
			{
			}
			goto SW_GENERATED_LABEL_0
		SW_GENERATED_LABEL_40:
			{
				noarch.Fprintf(outfile, (&[]byte("** Unknown modifier '%c'\n\x00")[0]), int32(uint8((pcre_uint8(*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) - (uintptr)(1)*unsafe.Sizeof(*pp)))))))))
				goto_SKIP_DATA = true
				goto SKIP_DATA_CONTAINER_OUTER
			}
		SW_GENERATED_LABEL_0:
		}
	SHOW_INFO_CONTAINER:
	SKIP_DATA_CONTAINER_OUTER:
		if !goto_SHOW_INFO && !goto_SKIP_DATA && ((posix != 0) || (do_posix != 0)) {
			var rc int32
			var cflags int32 = int32(0)
			if (options & int32(1)) != int32(0) {
				cflags |= int32(1)
			}
			if (options & int32(2)) != int32(0) {
				cflags |= int32(2)
			}
			if (options & int32(4)) != int32(0) {
				cflags |= int32(16)
			}
			if (options & int32(4096)) != int32(0) {
				cflags |= int32(32)
			}
			if (options & int32(2048)) != int32(0) {
				cflags |= int32(64)
			}
			if (options & int32(536870912)) != int32(0) {
				cflags |= int32(1024)
			}
			if (options & int32(512)) != int32(0) {
				cflags |= int32(512)
			}
			rc = regcomp(&preg, (*byte)(unsafe.Pointer(p)), cflags)
			if rc != int32(0) {
				_ = regerror(rc, &preg, (*byte)(unsafe.Pointer(buffer)), size_t(buffer_size))
				noarch.Fprintf(outfile, (&[]byte("Failed: POSIX code %d: %s\n\x00")[0]), rc, buffer)
				goto_SKIP_DATA = true
				goto SKIP_DATA_CONTAINER_OUTER
			}
		} else {
			if goto_SKIP_DATA {
				goto SKIP_DATA_CONTAINER
			}
			if goto_SHOW_INFO {
				goto SHOW_INFO
			}
			if timeit > int32(0) {
				var i int32
				var time_taken clock_t
				var start_time clock_t = clock()
				for i = int32(0); i < timeit; i++ {
					re = pcre_compile((*byte)(unsafe.Pointer(p)), options, &error, &erroroffset, (*uint8)(unsafe.Pointer(tables)))
					if re != nil {
						noarch.Free(unsafe.Pointer(re))
					}
				}
				total_compile_time += (func() clock_t {
					tempVar := (clock() - start_time)
					time_taken = tempVar
					return tempVar
				}())
				noarch.Fprintf(outfile, (&[]byte("Compile time %.4f milliseconds\n\x00")[0]), (((float64(int32((__clock_t((clock_t(time_taken)))))) * 1000) / float64(timeit)) / float64(int32((__clock_t((clock_t(int32(1000000)))))))))
			}
			re = pcre_compile((*byte)(unsafe.Pointer(p)), options, &error, &erroroffset, (*uint8)(unsafe.Pointer(tables)))
		SKIP_DATA_CONTAINER:
			if goto_SKIP_DATA || re == nil {
				if goto_SKIP_DATA {
					goto SKIP_DATA
				}
				noarch.Fprintf(outfile, (&[]byte("Failed: %s at offset %d\n\x00")[0]), error, erroroffset)
			SKIP_DATA:
				goto_SKIP_DATA = false
				if int64(uintptr(unsafe.Pointer(infile))) != int64(uintptr(unsafe.Pointer(stdin))) {
					for {
						if extend_inputline(infile, buffer, nil) == nil {
							done = int32(1)
							goto CONTINUE
						}
						len = noarch.Strlen((*byte)(unsafe.Pointer(buffer)))
						for (len > int32(0)) && ((int32(*((*uint16)(func() unsafe.Pointer {
							tempVar := (*linux.CtypeLoc())
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buffer)) + (uintptr)((len-int32(1)))*unsafe.Sizeof(*buffer))))))))))*unsafe.Sizeof(*tempVar))
						}()))) & int32(uint16(_ISspace))) != 0) {
							len -= 1
						}
						if len == int32(0) {
							break
						}
					}
					noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
				}
				goto CONTINUE
			}
			if new_info(re, nil, int32(0), unsafe.Pointer(&get_options)) < int32(0) {
				goto_SKIP_DATA = true
				goto SKIP_DATA_CONTAINER_OUTER
			}
			if (get_options & uint32(int32(2048))) != uint32(int32(0)) {
				use_utf = int32(1)
			}
			true_size = uint32(((*((*real_pcre)(unsafe.Pointer(re)))).size))
			if log_store != 0 {
				var name_count int32
				var name_entry_size int32
				var real_pcre_size int32
				new_info(re, nil, int32(8), unsafe.Pointer(&name_count))
				new_info(re, nil, int32(7), unsafe.Pointer(&name_entry_size))
				real_pcre_size = int32(0)
				if uint32(((*((*real_pcre)(unsafe.Pointer(re)))).flags & pcre_uint32((uint32(int32(1)))))) != 0 {
					real_pcre_size = int32(64)
				}
				new_info(re, nil, int32(1), unsafe.Pointer(&size))
				noarch.Fprintf(outfile, (&[]byte("Memory allocation (code space): %d\n\x00")[0]), int32((uint32(((size - size_t((uint32(real_pcre_size)))) - size_t((uint32((name_count * name_entry_size)))))))))
			}
			if (do_study != 0) || ((force_study >= int32(0)) && (noarch.NotInt32(no_force_study) != 0)) {
				if timeit > int32(0) {
					var i int32
					var time_taken clock_t
					var start_time clock_t = clock()
					for i = int32(0); i < timeit; i++ {
						extra = pcre_study(re, study_options, &error)
					}
					total_study_time = (func() clock_t {
						tempVar := (clock() - start_time)
						time_taken = tempVar
						return tempVar
					}())
					if extra != nil {
						pcre_free_study(extra)
					}
					noarch.Fprintf(outfile, (&[]byte("  Study time %.4f milliseconds\n\x00")[0]), (((float64(int32((__clock_t((clock_t(time_taken)))))) * 1000) / float64(timeit)) / float64(int32((__clock_t((clock_t(int32(1000000)))))))))
				}
				extra = pcre_study(re, study_options, &error)
				if error != nil {
					noarch.Fprintf(outfile, (&[]byte("Failed to study: %s\n\x00")[0]), error)
				} else if extra != nil {
					true_study_size = uint32((pcre_uint32((*((*pcre_study_data)(((*extra).study_data)))).size)))
					if log_store != 0 {
						var jitsize size_t
						if (new_info(re, extra, int32(17), unsafe.Pointer(&jitsize)) == int32(0)) && (jitsize != size_t((uint32(int32(0))))) {
							noarch.Fprintf(outfile, (&[]byte("Memory allocation (JIT code): %d\n\x00")[0]), int32(uint32((size_t(jitsize)))))
						}
					}
				}
			}
			if do_mark != 0 {
				if extra == nil {
					extra = (*pcre_extra)(noarch.Malloc(int32(64)))
					(*extra).flags = uint32(int32(0))
				}
				(*extra).mark = (**uint8)(unsafe.Pointer(&markptr))
				(*extra).flags |= uint32(int32(32))
			}
		SHOW_INFO:
			goto_SHOW_INFO = false
			if do_debug != 0 {
				noarch.Fprintf(outfile, (&[]byte("------------------------------------------------------------------\n\x00")[0]))
				pcre_printint(re, outfile, BOOL((debug_lengths)))
			}
			if do_showinfo != 0 {
				var all_options uint32
				var first_char pcre_uint32
				var need_char pcre_uint32
				var match_limit pcre_uint32
				var recursion_limit pcre_uint32
				var count int32
				var backrefmax int32
				var first_char_set int32
				var need_char_set int32
				var okpartial int32
				var jchanged int32
				var hascrorlf int32
				var maxlookbehind int32
				var match_empty int32
				var nameentrysize int32
				var namecount int32
				var nametable *pcre_uint8
				if (((((((((((((new_info(re, nil, int32(2), unsafe.Pointer(&count)) + new_info(re, nil, int32(3), unsafe.Pointer(&backrefmax))) + new_info(re, nil, int32(19), unsafe.Pointer(&first_char))) + new_info(re, nil, int32(20), unsafe.Pointer(&first_char_set))) + new_info(re, nil, int32(21), unsafe.Pointer(&need_char))) + new_info(re, nil, int32(22), unsafe.Pointer(&need_char_set))) + new_info(re, nil, int32(7), unsafe.Pointer(&nameentrysize))) + new_info(re, nil, int32(8), unsafe.Pointer(&namecount))) + new_info(re, nil, int32(9), unsafe.Pointer(&nametable))) + new_info(re, nil, int32(12), unsafe.Pointer(&okpartial))) + new_info(re, nil, int32(13), unsafe.Pointer(&jchanged))) + new_info(re, nil, int32(14), unsafe.Pointer(&hascrorlf))) + new_info(re, nil, int32(25), unsafe.Pointer(&match_empty))) + new_info(re, nil, int32(18), unsafe.Pointer(&maxlookbehind))) != int32(0) {
					goto_SKIP_DATA = true
					goto SKIP_DATA_CONTAINER_OUTER
				}
				noarch.Fprintf(outfile, (&[]byte("Capturing subpattern count = %d\n\x00")[0]), count)
				if backrefmax > int32(0) {
					noarch.Fprintf(outfile, (&[]byte("Max back reference = %d\n\x00")[0]), backrefmax)
				}
				if maxlookbehind > int32(0) {
					noarch.Fprintf(outfile, (&[]byte("Max lookbehind = %d\n\x00")[0]), maxlookbehind)
				}
				if new_info(re, nil, int32(23), unsafe.Pointer(&match_limit)) == int32(0) {
					noarch.Fprintf(outfile, (&[]byte("Match limit = %u\n\x00")[0]), pcre_uint32(match_limit))
				}
				if new_info(re, nil, int32(24), unsafe.Pointer(&recursion_limit)) == int32(0) {
					noarch.Fprintf(outfile, (&[]byte("Recursion limit = %u\n\x00")[0]), pcre_uint32(recursion_limit))
				}
				if namecount > int32(0) {
					noarch.Fprintf(outfile, (&[]byte("Named capturing subpatterns:\n\x00")[0]))
					for func() int32 {
						defer func() {
							namecount -= 1
						}()
						return namecount
					}() > int32(0) {
						var imm2_size int32 = func() int32 {
							if pcre_mode == PCRE8_MODE {
								return int32(2)
							} else {
								return int32(1)
							}
						}()
						var length int32 = (noarch.Strlen(((*byte)(func() unsafe.Pointer {
							tempVar := (*byte)(unsafe.Pointer(nametable))
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(imm2_size)*unsafe.Sizeof(*tempVar))
						}()))))
						noarch.Fprintf(outfile, (&[]byte("  \x00")[0]))
						_ = pchars(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((nametable))) + (uintptr)(imm2_size)*unsafe.Sizeof(*(nametable))))), length, outfile)
						for func() int32 {
							defer func() {
								length += 1
							}()
							return length
						}() < (nameentrysize - imm2_size) {
							noarch.Fputc(int32(' '), outfile)
						}
						if pcre_mode == PCRE8_MODE {
							noarch.Fprintf(outfile, (&[]byte("%3d\n\x00")[0]), ((int32(uint8((pcre_uint8(*nametable)))) << uint64(int32(8))) | int32(uint8((pcre_uint8(*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(nametable)) + (uintptr)(int32(1))*unsafe.Sizeof(*nametable))))))))))
						}
						nametable = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(nametable)) + (uintptr)((nameentrysize*int32(1)))*unsafe.Sizeof(*nametable))))
					}
				}
				if noarch.NotInt32(okpartial) != 0 {
					noarch.Fprintf(outfile, (&[]byte("Partial matching not supported\n\x00")[0]))
				}
				if hascrorlf != 0 {
					noarch.Fprintf(outfile, (&[]byte("Contains explicit CR or LF match\n\x00")[0]))
				}
				if match_empty != 0 {
					noarch.Fprintf(outfile, (&[]byte("May match empty string\n\x00")[0]))
				}
				all_options = uint32(((*((*real_pcre)(unsafe.Pointer(re)))).options))
				if do_flip != 0 {
					all_options = uint32((swap_uint32(pcre_uint32(all_options))))
				}
				if get_options == uint32(int32(0)) {
					noarch.Fprintf(outfile, (&[]byte("No options\n\x00")[0]))
				} else {
					noarch.Fprintf(outfile, (&[]byte("Options:%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s\n\x00")[0]), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(16))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" anchored\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(1))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" caseless\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(8))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" extended\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(2))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" multiline\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(262144))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" firstline\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(4))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" dotall\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(8388608))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" bsr_anycrlf\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(16777216))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" bsr_unicode\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(32))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" dollar_endonly\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(64))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" extra\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(512))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" ungreedy\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(4096))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" no_auto_capture\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(131072))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" no_auto_possessify\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(2048))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" utf\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(536870912))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" ucp\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(8192))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" no_utf_check\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(67108864))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" no_start_optimize\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(524288))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" dupnames\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}(), func() *byte {
						if (map[bool]int32{false: 0, true: 1}[((get_options & uint32(int32(65536))) != uint32(int32(0)))]) != 0 {
							return (&[]byte(" never_utf\x00")[0])
						} else {
							return (&[]byte("\x00")[0])
						}
					}())
				}
				if jchanged != 0 {
					noarch.Fprintf(outfile, (&[]byte("Duplicate name status changes\n\x00")[0]))
				}
				switch get_options & uint32((((int32(1048576) | int32(2097152)) | int32(4194304)) | int32(5242880))) {
				case uint32(int32(1048576)):
					{
						noarch.Fprintf(outfile, (&[]byte("Forced newline sequence: CR\n\x00")[0]))
					}
				case uint32(int32(2097152)):
					{
						noarch.Fprintf(outfile, (&[]byte("Forced newline sequence: LF\n\x00")[0]))
					}
				case uint32(int32(3145728)):
					{
						noarch.Fprintf(outfile, (&[]byte("Forced newline sequence: CRLF\n\x00")[0]))
					}
				case uint32(int32(5242880)):
					{
						noarch.Fprintf(outfile, (&[]byte("Forced newline sequence: ANYCRLF\n\x00")[0]))
					}
				case uint32(int32(4194304)):
					{
						noarch.Fprintf(outfile, (&[]byte("Forced newline sequence: ANY\n\x00")[0]))
					}
				default:
					{
						break
					}
				}
				if first_char_set == int32(2) {
					noarch.Fprintf(outfile, (&[]byte("First char at start or follows newline\n\x00")[0]))
				} else if first_char_set == int32(1) {
					var caseless *byte = func() *byte {
						if (map[bool]int32{false: 0, true: 1}[(((*((*real_pcre)(unsafe.Pointer(re)))).flags & pcre_uint32((uint32(int32(32))))) == pcre_uint32((uint32(int32(0)))))]) != 0 {
							return (&[]byte("\x00")[0])
						} else {
							return (&[]byte(" (caseless)\x00")[0])
						}
					}()
					if (func() int32 {
						if locale_set != 0 {
							return (map[bool]int32{false: 0, true: 1}[((first_char < pcre_uint32((uint32(int32(256))))) && ((int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((pcre_uint32((first_char))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISprint))) != 0))])
						} else {
							return (map[bool]int32{false: 0, true: 1}[((first_char >= pcre_uint32((uint32(int32(32))))) && (first_char < pcre_uint32((uint32(int32(127))))))])
						}
					}()) != 0 {
						noarch.Fprintf(outfile, (&[]byte("First char = '%c'%s\n\x00")[0]), pcre_uint32(first_char), caseless)
					} else {
						noarch.Fprintf(outfile, (&[]byte("First char = \x00")[0]))
						pchar(pcre_uint32(first_char), outfile)
						noarch.Fprintf(outfile, (&[]byte("%s\n\x00")[0]), caseless)
					}
				} else {
					noarch.Fprintf(outfile, (&[]byte("No first char\n\x00")[0]))
				}
				if need_char_set == int32(0) {
					noarch.Fprintf(outfile, (&[]byte("No need char\n\x00")[0]))
				} else {
					var caseless *byte = func() *byte {
						if (map[bool]int32{false: 0, true: 1}[(((*((*real_pcre)(unsafe.Pointer(re)))).flags & pcre_uint32((uint32(int32(128))))) == pcre_uint32((uint32(int32(0)))))]) != 0 {
							return (&[]byte("\x00")[0])
						} else {
							return (&[]byte(" (caseless)\x00")[0])
						}
					}()
					if (func() int32 {
						if locale_set != 0 {
							return (map[bool]int32{false: 0, true: 1}[((need_char < pcre_uint32((uint32(int32(256))))) && ((int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((pcre_uint32((need_char))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISprint))) != 0))])
						} else {
							return (map[bool]int32{false: 0, true: 1}[((need_char >= pcre_uint32((uint32(int32(32))))) && (need_char < pcre_uint32((uint32(int32(127))))))])
						}
					}()) != 0 {
						noarch.Fprintf(outfile, (&[]byte("Need char = '%c'%s\n\x00")[0]), pcre_uint32(need_char), caseless)
					} else {
						noarch.Fprintf(outfile, (&[]byte("Need char = \x00")[0]))
						pchar(pcre_uint32(need_char), outfile)
						noarch.Fprintf(outfile, (&[]byte("%s\n\x00")[0]), caseless)
					}
				}
				if (do_study != 0) || (((force_study >= int32(0)) && (showinfo != 0)) && (noarch.NotInt32(no_force_study) != 0)) {
					if extra == nil {
						noarch.Fprintf(outfile, (&[]byte("Study returned NULL\n\x00")[0]))
					} else {
						var start_bits *pcre_uint8 = nil
						var minlength int32
						if new_info(re, extra, int32(15), unsafe.Pointer(&minlength)) == int32(0) {
							noarch.Fprintf(outfile, (&[]byte("Subject length lower bound = %d\n\x00")[0]), minlength)
						}
						if new_info(re, extra, int32(5), unsafe.Pointer(&start_bits)) == int32(0) {
							if start_bits == nil {
								noarch.Fprintf(outfile, (&[]byte("No starting char list\n\x00")[0]))
							} else {
								var i int32
								var c int32 = int32(24)
								noarch.Fprintf(outfile, (&[]byte("Starting chars: \x00")[0]))
								for i = int32(0); i < int32(256); i++ {
									if (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(start_bits)) + (uintptr)((i/int32(8)))*unsafe.Sizeof(*start_bits))))))) & (int32(1) << uint64((i & int32(7))))) != int32(0) {
										if c > int32(75) {
											noarch.Fprintf(outfile, (&[]byte("\n  \x00")[0]))
											c = int32(2)
										}
										if (func() int32 {
											if locale_set != 0 {
												return (map[bool]int32{false: 0, true: 1}[((i < int32(256)) && ((int32(*((*uint16)(func() unsafe.Pointer {
													tempVar := (*linux.CtypeLoc())
													return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)((i))*unsafe.Sizeof(*tempVar))
												}()))) & int32(uint16(_ISprint))) != 0))])
											} else {
												return (map[bool]int32{false: 0, true: 1}[((i >= int32(32)) && (i < int32(127)))])
											}
										}() != 0) && (i != int32(' ')) {
											noarch.Fprintf(outfile, (&[]byte("%c \x00")[0]), i)
											c += int32(2)
										} else {
											noarch.Fprintf(outfile, (&[]byte("\\x%02x \x00")[0]), i)
											c += int32(5)
										}
									}
								}
								noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
							}
						}
					}
					if ((study_options & ((int32(1) | int32(2)) | int32(4))) != int32(0)) && ((force_study_options & ((int32(1) | int32(2)) | int32(4))) == int32(0)) {
						var jit int32
						if new_info(re, extra, int32(16), unsafe.Pointer(&jit)) == int32(0) {
							if jit != 0 {
								noarch.Fprintf(outfile, (&[]byte("JIT study was successful\n\x00")[0]))
							} else {
								noarch.Fprintf(outfile, (&[]byte("JIT support is not available in this version of PCRE\n\x00")[0]))
							}
						}
					}
				}
			}
			if to_file != nil {
				var f *noarch.File = noarch.Fopen((*byte)(unsafe.Pointer(to_file)), (&[]byte("wb\x00")[0]))
				if f == nil {
					noarch.Fprintf(outfile, (&[]byte("Unable to open %s: %s\n\x00")[0]), to_file, noarch.Strerror((*noarch.Errno())))
				} else {
					var sbuf []pcre_uint8 = make([]pcre_uint8, 8, 8)
					if do_flip != 0 {
						regexflip(re, extra)
					}
					*&sbuf[0] = pcre_uint8(((true_size >> uint64(int32(24))) & uint32(int32(255))))
					*((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &sbuf[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
					}())) = pcre_uint8(((true_size >> uint64(int32(16))) & uint32(int32(255))))
					*((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &sbuf[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
					}())) = pcre_uint8(((true_size >> uint64(int32(8))) & uint32(int32(255))))
					*((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &sbuf[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
					}())) = pcre_uint8((true_size & uint32(int32(255))))
					*((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &sbuf[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
					}())) = pcre_uint8(((true_study_size >> uint64(int32(24))) & uint32(int32(255))))
					*((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &sbuf[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
					}())) = pcre_uint8(((true_study_size >> uint64(int32(16))) & uint32(int32(255))))
					*((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &sbuf[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
					}())) = pcre_uint8(((true_study_size >> uint64(int32(8))) & uint32(int32(255))))
					*((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &sbuf[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
					}())) = pcre_uint8((true_study_size & uint32(int32(255))))
					if (noarch.Fwrite((*byte)(unsafe.Pointer(&sbuf[0])), int32(int32(1)), int32(int32(8)), f) < int32(uint32(int32(8)))) || (noarch.Fwrite((*byte)(unsafe.Pointer(re)), int32(int32(1)), int32(true_size), f) < int32(true_size)) {
						noarch.Fprintf(outfile, (&[]byte("Write error on %s: %s\n\x00")[0]), to_file, noarch.Strerror((*noarch.Errno())))
					} else {
						noarch.Fprintf(outfile, (&[]byte("Compiled pattern written to %s\n\x00")[0]), to_file)
						if extra != nil {
							if noarch.Fwrite((*byte)((*extra).study_data), int32(int32(1)), int32(true_study_size), f) < int32(true_study_size) {
								noarch.Fprintf(outfile, (&[]byte("Write error on %s: %s\n\x00")[0]), to_file, noarch.Strerror((*noarch.Errno())))
							} else {
								noarch.Fprintf(outfile, (&[]byte("Study data written to %s\n\x00")[0]), to_file)
							}
						}
					}
					noarch.Fclose(f)
				}
				new_free(unsafe.Pointer(re))
				if extra != nil {
					pcre_free_study(extra)
				}
				if locale_set != 0 {
					new_free(unsafe.Pointer(tables))
					setlocale(int32(0), (&[]byte("C\x00")[0]))
					locale_set = int32(0)
				}
				continue
			}
		}
		for {
			var q8 *pcre_uint8
			var bptr *pcre_uint8
			var use_offsets *int32 = offsets
			var use_size_offsets int32 = size_offsets
			var callout_data int32 = int32(0)
			var callout_data_set int32 = int32(0)
			var count int32
			var c pcre_uint32
			var copystrings int32 = int32(0)
			var find_match_limit int32 = default_find_match_limit
			var getstrings int32 = int32(0)
			var getlist int32 = int32(0)
			var gmatched int32 = int32(0)
			var start_offset int32 = int32(0)
			var start_offset_sign int32 = int32(1)
			var g_notempty int32 = int32(0)
			var use_dfa int32 = int32(0)
			copynames[0] = pcre_uint32(int32(0))
			getnames[0] = pcre_uint32(int32(0))
			cn8ptr = copynames8
			gn8ptr = getnames8
			pcre_callout = callout
			first_callout = int32(1)
			last_callout_mark = nil
			callout_extra = int32(0)
			callout_count = int32(0)
			callout_fail_count = int32(999999)
			callout_fail_id = -int32(1)
			show_malloc = int32(0)
			options = int32(0)
			if extra != nil {
				(*extra).flags &= ^uint32((int32(2) | int32(16)))
			}
			len = int32(0)
			for {
				if extend_inputline(infile, ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buffer))+(uintptr)(len)*unsafe.Sizeof(*buffer)))), (&[]byte("data> \x00")[0])) == nil {
					if len > int32(0) {
						noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
						break
					}
					done = int32(1)
					goto CONTINUE
				}
				if int64(uintptr(unsafe.Pointer(infile))) != int64(uintptr(unsafe.Pointer(stdin))) {
					noarch.Fprintf(outfile, (&[]byte("%s\x00")[0]), (*byte)(unsafe.Pointer(buffer)))
				}
				len = noarch.Strlen((*byte)(unsafe.Pointer(buffer)))
				if int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buffer)) + (uintptr)((len-int32(1)))*unsafe.Sizeof(*buffer))))))) == int32('\n') {
					break
				}
			}
			for (len > int32(0)) && ((int32(*((*uint16)(func() unsafe.Pointer {
				tempVar := (*linux.CtypeLoc())
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buffer)) + (uintptr)((len-int32(1)))*unsafe.Sizeof(*buffer))))))))))*unsafe.Sizeof(*tempVar))
			}()))) & int32(uint16(_ISspace))) != 0) {
				len -= 1
			}
			*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buffer)) + (uintptr)(len)*unsafe.Sizeof(*buffer)))) = pcre_uint8(int32(0))
			if len == int32(0) {
				break
			}
			p = buffer
			for (int32(*((*uint16)(func() unsafe.Pointer {
				tempVar := (*linux.CtypeLoc())
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
			}()))) & int32(uint16(_ISspace))) != 0 {
				p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
			}
			if use_utf != 0 {
				var q *pcre_uint8
				var cc pcre_uint32
				var n int32 = int32(1)
				for q = p; (n > int32(0)) && (int32(uint8((*q))) != 0); q = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + (uintptr)(n)*unsafe.Sizeof(*q)))) {
					n = utf82ord(q, &cc)
				}
				if n <= int32(0) {
					noarch.Fprintf(outfile, (&[]byte("**Failed: invalid UTF-8 string cannot be used as input in UTF mode\n\x00")[0]))
					goto NEXT_DATA
				}
			}
			for (dbuffer == nil) || (size_t(len) >= dbuffer_size) {
				dbuffer_size *= size_t((uint32(int32(2))))
				dbuffer = (*pcre_uint8)(noarch.Malloc(int32(uint32((dbuffer_size * size_t((uint32(int32(1)))))))))
				if dbuffer == nil {
					noarch.Fprintf(stderr, (&[]byte("pcretest: realloc(%d) failed\n\x00")[0]), int32(uint32((size_t(dbuffer_size)))))
					noarch.Exit(int32(1))
				}
			}
			q8 = dbuffer
			for (func() pcre_uint32 {
				tempVar := pcre_uint32((uint32(uint8((*func() *pcre_uint8 {
					defer func() {
						p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
					}()
					return p
				}())))))
				c = tempVar
				return tempVar
			}()) != pcre_uint32((uint32(int32(0)))) {
				var i int32 = int32(0)
				var n int32 = int32(0)
				if c != pcre_uint32((uint32('\\'))) {
					if (use_utf != 0) && (c >= pcre_uint32((uint32(int32(192))))) {
						{
							if (c & pcre_uint32((uint32(int32(32))))) == pcre_uint32((uint32(int32(0)))) {
								c = (((c & pcre_uint32((uint32(int32(31))))) << uint64(int32(6))) | pcre_uint32((uint32((int32(uint8((*func() *pcre_uint8 {
									defer func() {
										p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
									}()
									return p
								}()))) & int32(63))))))
							} else if (c & pcre_uint32((uint32(int32(16))))) == pcre_uint32((uint32(int32(0)))) {
								c = ((((c & pcre_uint32((uint32(int32(15))))) << uint64(int32(12))) | pcre_uint32((uint32(((int32(uint8((*p))) & int32(63)) << uint64(int32(6))))))) | pcre_uint32((uint32((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(1))*unsafe.Sizeof(*p))))))) & int32(63))))))
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(2))*unsafe.Sizeof(*p))))
							} else if (c & pcre_uint32((uint32(int32(8))))) == pcre_uint32((uint32(int32(0)))) {
								c = (((((c & pcre_uint32((uint32(int32(7))))) << uint64(int32(18))) | pcre_uint32((uint32(((int32(uint8((*p))) & int32(63)) << uint64(int32(12))))))) | pcre_uint32((uint32(((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(1))*unsafe.Sizeof(*p))))))) & int32(63)) << uint64(int32(6))))))) | pcre_uint32((uint32((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(2))*unsafe.Sizeof(*p))))))) & int32(63))))))
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(3))*unsafe.Sizeof(*p))))
							} else if (c & pcre_uint32((uint32(int32(4))))) == pcre_uint32((uint32(int32(0)))) {
								c = ((((((c & pcre_uint32((uint32(int32(3))))) << uint64(int32(24))) | pcre_uint32((uint32(((int32(uint8((*p))) & int32(63)) << uint64(int32(18))))))) | pcre_uint32((uint32(((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(1))*unsafe.Sizeof(*p))))))) & int32(63)) << uint64(int32(12))))))) | pcre_uint32((uint32(((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(2))*unsafe.Sizeof(*p))))))) & int32(63)) << uint64(int32(6))))))) | pcre_uint32((uint32((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(3))*unsafe.Sizeof(*p))))))) & int32(63))))))
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(4))*unsafe.Sizeof(*p))))
							} else {
								c = (((((((c & pcre_uint32((uint32(int32(1))))) << uint64(int32(30))) | pcre_uint32((uint32(((int32(uint8((*p))) & int32(63)) << uint64(int32(24))))))) | pcre_uint32((uint32(((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(1))*unsafe.Sizeof(*p))))))) & int32(63)) << uint64(int32(18))))))) | pcre_uint32((uint32(((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(2))*unsafe.Sizeof(*p))))))) & int32(63)) << uint64(int32(12))))))) | pcre_uint32((uint32(((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(3))*unsafe.Sizeof(*p))))))) & int32(63)) << uint64(int32(6))))))) | pcre_uint32((uint32((int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(4))*unsafe.Sizeof(*p))))))) & int32(63))))))
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(int32(5))*unsafe.Sizeof(*p))))
							}
						}
					}
				} else {
					switch func() pcre_uint32 {
						c = pcre_uint32((uint32(uint8((*func() *pcre_uint8 {
							defer func() {
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
							}()
							return p
						}())))))
						return c
					}() {
					case pcre_uint32('a'):
						{
							c = pcre_uint32((uint32('\a')))
						}
					case pcre_uint32('b'):
						{
							c = pcre_uint32((uint32('\b')))
						}
					case pcre_uint32('e'):
						{
							c = pcre_uint32((uint32('\x1b')))
						}
					case pcre_uint32('f'):
						{
							c = pcre_uint32((uint32('\f')))
						}
					case pcre_uint32('n'):
						{
							c = pcre_uint32((uint32('\n')))
						}
					case pcre_uint32('r'):
						{
							c = pcre_uint32((uint32('\r')))
						}
					case pcre_uint32('t'):
						{
							c = pcre_uint32((uint32('\t')))
						}
					case pcre_uint32('v'):
						{
							c = pcre_uint32((uint32('\v')))
						}
					case pcre_uint32('0'):
						fallthrough
					case pcre_uint32('1'):
						fallthrough
					case pcre_uint32('2'):
						fallthrough
					case pcre_uint32('3'):
						fallthrough
					case pcre_uint32('4'):
						fallthrough
					case pcre_uint32('5'):
						fallthrough
					case pcre_uint32('6'):
						fallthrough
					case pcre_uint32('7'):
						{
							c -= pcre_uint32((uint32('0')))
							for (((func() int32 {
								defer func() {
									i += 1
								}()
								return i
							}() < int32(2)) && ((int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISdigit))) != 0)) && (int32(uint8((*p))) != int32('8'))) && (int32(uint8((*p))) != int32('9')) {
								c = (((c * pcre_uint32((uint32(int32(8))))) + pcre_uint32((uint32(uint8((*func() *pcre_uint8 {
									defer func() {
										p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
									}()
									return p
								}())))))) - pcre_uint32((uint32('0'))))
							}
						}
					case pcre_uint32('o'):
						{
							if int32(uint8((*p))) == int32('{') {
								var pt *pcre_uint8 = p
								c = pcre_uint32(int32(0))
								for pt = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pt)) + (uintptr)(1)*unsafe.Sizeof(*pt)))); (((int32(*((*uint16)(func() unsafe.Pointer {
									tempVar := (*linux.CtypeLoc())
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*pt))))))*unsafe.Sizeof(*tempVar))
								}()))) & int32(uint16(_ISdigit))) != 0) && (int32(uint8((*pt))) != int32('8'))) && (int32(uint8((*pt))) != int32('9')); pt = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pt)) + (uintptr)(1)*unsafe.Sizeof(*pt)))) {
									if func() int32 {
										i += 1
										return i
									}() == int32(12) {
										noarch.Fprintf(outfile, (&[]byte("** Too many octal digits in \\o{...} item; using only the first twelve.\n\x00")[0]))
									} else {
										c = (((c * pcre_uint32((uint32(int32(8))))) + pcre_uint32((uint32(uint8((*pt)))))) - pcre_uint32((uint32('0'))))
									}
								}
								if int32(uint8((*pt))) == int32('}') {
									p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pt)) + (uintptr)(int32(1))*unsafe.Sizeof(*pt))))
								} else {
									noarch.Fprintf(outfile, (&[]byte("** Missing } after \\o{ (assumed)\n\x00")[0]))
								}
							}
						}
					case pcre_uint32('x'):
						{
							if int32(uint8((*p))) == int32('{') {
								var pt *pcre_uint8 = p
								c = pcre_uint32(int32(0))
								for pt = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pt)) + (uintptr)(1)*unsafe.Sizeof(*pt)))); (int32(*((*uint16)(func() unsafe.Pointer {
									tempVar := (*linux.CtypeLoc())
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*pt))))))*unsafe.Sizeof(*tempVar))
								}()))) & int32(uint16(_ISxdigit))) != 0; pt = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pt)) + (uintptr)(1)*unsafe.Sizeof(*pt)))) {
									if func() int32 {
										i += 1
										return i
									}() == int32(9) {
										noarch.Fprintf(outfile, (&[]byte("** Too many hex digits in \\x{...} item; using only the first eight.\n\x00")[0]))
									} else {
										c = (((c * pcre_uint32((uint32(int32(16))))) + pcre_uint32((uint32(linux.ToLower(int32(uint8((pcre_uint8(*pt))))))))) - pcre_uint32((uint32(func() int32 {
											if (int32(*((*uint16)(func() unsafe.Pointer {
												tempVar := (*linux.CtypeLoc())
												return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*pt))))))*unsafe.Sizeof(*tempVar))
											}()))) & int32(uint16(_ISdigit))) != 0 {
												return int32('0')
											} else {
												return int32(('a' - byte(int32(10))))
											}
										}()))))
									}
								}
								if int32(uint8((*pt))) == int32('}') {
									p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pt)) + (uintptr)(int32(1))*unsafe.Sizeof(*pt))))
									break
								}
							}
							c = pcre_uint32(int32(0))
							for (func() int32 {
								defer func() {
									i += 1
								}()
								return i
							}() < int32(2)) && ((int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISxdigit))) != 0) {
								c = (((c * pcre_uint32((uint32(int32(16))))) + pcre_uint32((uint32(linux.ToLower(int32(uint8((pcre_uint8(*p))))))))) - pcre_uint32((uint32(func() int32 {
									if (int32(*((*uint16)(func() unsafe.Pointer {
										tempVar := (*linux.CtypeLoc())
										return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
									}()))) & int32(uint16(_ISdigit))) != 0 {
										return int32('0')
									} else {
										return int32(('a' - byte(int32(10))))
									}
								}()))))
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
							}
							if (use_utf != 0) && (pcre_mode == PCRE8_MODE) {
								*func() *pcre_uint8 {
									defer func() {
										q8 = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(q8)) + (uintptr)(1)*unsafe.Sizeof(*q8))))
									}()
									return q8
								}() = pcre_uint8((uint8(uint32((c)))))
								continue
							}
						}
					case pcre_uint32(int32(0)):
						{
							p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) - (uintptr)(1)*unsafe.Sizeof(*p))))
							continue
						}
						fallthrough
					case pcre_uint32('>'):
						{
							if int32(uint8((*p))) == int32('-') {
								start_offset_sign = -int32(1)
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
							}
							for (int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISdigit))) != 0 {
								start_offset = (((start_offset * int32(10)) + int32(uint8((*func() *pcre_uint8 {
									defer func() {
										p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
									}()
									return p
								}())))) - int32('0'))
							}
							start_offset *= start_offset_sign
							continue
						}
						fallthrough
					case pcre_uint32('A'):
						{
							options |= int32(16)
							continue
						}
						fallthrough
					case pcre_uint32('B'):
						{
							options |= int32(128)
							continue
						}
						fallthrough
					case pcre_uint32('C'):
						{
							if (int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISdigit))) != 0 {
								for (int32(*((*uint16)(func() unsafe.Pointer {
									tempVar := (*linux.CtypeLoc())
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
								}()))) & int32(uint16(_ISdigit))) != 0 {
									n = (((n * int32(10)) + int32(uint8((*func() *pcre_uint8 {
										defer func() {
											p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
										}()
										return p
									}())))) - int32('0'))
								}
								copystrings |= (int32(1) << uint64(n))
							} else if (int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISalnum))) != 0 {
								p = read_capture_name8(p, &cn8ptr, re)
							} else if int32(uint8((*p))) == int32('+') {
								callout_extra = int32(1)
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
							} else if int32(uint8((*p))) == int32('-') {
								pcre_callout = nil
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
							} else if int32(uint8((*p))) == int32('!') {
								callout_fail_id = int32(0)
								p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
								for (int32(*((*uint16)(func() unsafe.Pointer {
									tempVar := (*linux.CtypeLoc())
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
								}()))) & int32(uint16(_ISdigit))) != 0 {
									callout_fail_id = (((callout_fail_id * int32(10)) + int32(uint8((*func() *pcre_uint8 {
										defer func() {
											p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
										}()
										return p
									}())))) - int32('0'))
								}
								callout_fail_count = int32(0)
								if int32(uint8((*p))) == int32('!') {
									p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
									for (int32(*((*uint16)(func() unsafe.Pointer {
										tempVar := (*linux.CtypeLoc())
										return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
									}()))) & int32(uint16(_ISdigit))) != 0 {
										callout_fail_count = (((callout_fail_count * int32(10)) + int32(uint8((*func() *pcre_uint8 {
											defer func() {
												p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
											}()
											return p
										}())))) - int32('0'))
									}
								}
							} else if int32(uint8((*p))) == int32('*') {
								var sign int32 = int32(1)
								callout_data = int32(0)
								if int32(uint8((*func() *pcre_uint8 {
									p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
									return p
								}()))) == int32('-') {
									sign = -int32(1)
									p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
								}
								for (int32(*((*uint16)(func() unsafe.Pointer {
									tempVar := (*linux.CtypeLoc())
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
								}()))) & int32(uint16(_ISdigit))) != 0 {
									callout_data = (((callout_data * int32(10)) + int32(uint8((*func() *pcre_uint8 {
										defer func() {
											p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
										}()
										return p
									}())))) - int32('0'))
								}
								callout_data *= sign
								callout_data_set = int32(1)
							}
							continue
						}
						fallthrough
					case pcre_uint32('D'):
						{
							if (posix != 0) || (do_posix != 0) {
								noarch.Printf((&[]byte("** Can't use dfa matching in POSIX mode: \\D ignored\n\x00")[0]))
							} else {
								use_dfa = int32(1)
							}
							continue
						}
						fallthrough
					case pcre_uint32('F'):
						{
							options |= int32(65536)
							continue
						}
						fallthrough
					case pcre_uint32('G'):
						{
							if (int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISdigit))) != 0 {
								for (int32(*((*uint16)(func() unsafe.Pointer {
									tempVar := (*linux.CtypeLoc())
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
								}()))) & int32(uint16(_ISdigit))) != 0 {
									n = (((n * int32(10)) + int32(uint8((*func() *pcre_uint8 {
										defer func() {
											p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
										}()
										return p
									}())))) - int32('0'))
								}
								getstrings |= (int32(1) << uint64(n))
							} else if (int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISalnum))) != 0 {
								p = read_capture_name8(p, &gn8ptr, re)
							}
							continue
						}
						fallthrough
					case pcre_uint32('J'):
						{
							for (int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISdigit))) != 0 {
								n = (((n * int32(10)) + int32(uint8((*func() *pcre_uint8 {
									defer func() {
										p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
									}()
									return p
								}())))) - int32('0'))
							}
							if ((extra != nil) && (((*extra).flags & uint32(int32(64))) != uint32(int32(0)))) && ((*extra).executable_jit != nil) {
								if jit_stack != nil {
									pcre_jit_stack_free(jit_stack)
								}
								jit_stack = pcre_jit_stack_alloc(int32(1), (n * int32(1024)))
								pcre_assign_jit_stack(extra, jit_callback, unsafe.Pointer(jit_stack))
							}
							continue
						}
						fallthrough
					case pcre_uint32('L'):
						{
							getlist = int32(1)
							continue
						}
						fallthrough
					case pcre_uint32('M'):
						{
							find_match_limit = int32(1)
							continue
						}
						fallthrough
					case pcre_uint32('N'):
						{
							if (options & int32(1024)) != int32(0) {
								options = ((options & ^int32(1024)) | int32(268435456))
							} else {
								options |= int32(1024)
							}
							continue
						}
						fallthrough
					case pcre_uint32('O'):
						{
							for (int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISdigit))) != 0 {
								if n > ((int32(2147483647) - int32(10)) / int32(10)) {
									noarch.Printf((&[]byte("** \\O argument is too big\n\x00")[0]))
									yield = int32(1)
									goto EXIT
								}
								n = (((n * int32(10)) + int32(uint8((*func() *pcre_uint8 {
									defer func() {
										p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
									}()
									return p
								}())))) - int32('0'))
							}
							if n > size_offsets_max {
								size_offsets_max = n
								noarch.Free(unsafe.Pointer(offsets))
								offsets = (*int32)(noarch.Malloc(int32((uint32(size_offsets_max) * 4))))
								use_offsets = offsets
								if offsets == nil {
									noarch.Printf((&[]byte("** Failed to get %d bytes of memory for offsets vector\n\x00")[0]), int32((uint32(size_offsets_max) * 4)))
									yield = int32(1)
									goto EXIT
								}
							}
							use_size_offsets = n
							if n == int32(0) {
								use_offsets = nil
							} else {
								use_offsets = ((*int32)(func() unsafe.Pointer {
									tempVar := ((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(offsets)) + (uintptr)(size_offsets_max)*unsafe.Sizeof(*offsets))))
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) - (uintptr)(n)*unsafe.Sizeof(*tempVar))
								}()))
							}
							continue
						}
						fallthrough
					case pcre_uint32('P'):
						{
							options |= func() int32 {
								if (map[bool]int32{false: 0, true: 1}[((options & int32(32768)) == int32(0))]) != 0 {
									return int32(32768)
								} else {
									return int32(134217728)
								}
							}()
							continue
						}
						fallthrough
					case pcre_uint32('Q'):
						{
							for (int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISdigit))) != 0 {
								n = (((n * int32(10)) + int32(uint8((*func() *pcre_uint8 {
									defer func() {
										p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
									}()
									return p
								}())))) - int32('0'))
							}
							if extra == nil {
								extra = (*pcre_extra)(noarch.Malloc(int32(64)))
								(*extra).flags = uint32(int32(0))
							}
							(*extra).flags |= uint32(int32(16))
							(*extra).match_limit_recursion = uint32(n)
							continue
						}
						fallthrough
					case pcre_uint32('q'):
						{
							for (int32(*((*uint16)(func() unsafe.Pointer {
								tempVar := (*linux.CtypeLoc())
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((pcre_uint8((*p))))))*unsafe.Sizeof(*tempVar))
							}()))) & int32(uint16(_ISdigit))) != 0 {
								n = (((n * int32(10)) + int32(uint8((*func() *pcre_uint8 {
									defer func() {
										p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
									}()
									return p
								}())))) - int32('0'))
							}
							if extra == nil {
								extra = (*pcre_extra)(noarch.Malloc(int32(64)))
								(*extra).flags = uint32(int32(0))
							}
							(*extra).flags |= uint32(int32(2))
							(*extra).match_limit = uint32(n)
							continue
						}
						fallthrough
					case pcre_uint32('R'):
						{
							options |= int32(131072)
							continue
						}
						fallthrough
					case pcre_uint32('S'):
						{
							show_malloc = int32(1)
							continue
						}
						fallthrough
					case pcre_uint32('Y'):
						{
							options |= int32(67108864)
							continue
						}
						fallthrough
					case pcre_uint32('Z'):
						{
							options |= int32(256)
							continue
						}
						fallthrough
					case pcre_uint32('?'):
						{
							options |= int32(8192)
							continue
						}
						fallthrough
					case pcre_uint32('<'):
						{
							var x int32 = check_mc_option(p, outfile, BOOL((int32(1))), (&[]byte("escape sequence\x00")[0]))
							if x == int32(0) {
								goto NEXT_DATA
							}
							options |= x
							for int32(uint8((*func() *pcre_uint8 {
								defer func() {
									p = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
								}()
								return p
							}()))) != int32('>') {
							}
							continue
						}
					}
				}
				if pcre_mode == PCRE8_MODE {
					if use_utf != 0 {
						if c > pcre_uint32((uint32(int32(2147483647)))) {
							noarch.Fprintf(outfile, (&[]byte("** Character \\x{%x} is greater than 0x7fffffff and so cannot be converted to UTF-8\n\x00")[0]), pcre_uint32(c))
							goto NEXT_DATA
						}
						q8 = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(q8)) + (uintptr)(ord2utf8(pcre_uint32(c), q8))*unsafe.Sizeof(*q8))))
					} else {
						if c > pcre_uint32((uint32(255))) {
							noarch.Fprintf(outfile, (&[]byte("** Character \\x{%x} is greater than 255 and UTF-8 mode is not enabled.\n\x00")[0]), pcre_uint32(c))
							noarch.Fprintf(outfile, (&[]byte("** Truncation will probably give the wrong result.\n\x00")[0]))
						}
						*func() *pcre_uint8 {
							defer func() {
								q8 = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(q8)) + (uintptr)(1)*unsafe.Sizeof(*q8))))
							}()
							return q8
						}() = pcre_uint8((uint8(uint32((c)))))
					}
				}
			}
			if pcre_mode == PCRE8_MODE {
				*q8 = pcre_uint8(int32(0))
				len = (int32((int64(uintptr(unsafe.Pointer(q8))) - int64(uintptr(unsafe.Pointer(dbuffer))))))
			}
			bptr = dbuffer
			if (posix != 0) || (do_posix != 0) {
				noarch.Memcpy(unsafe.Pointer(((*pcre_uint8)(func() unsafe.Pointer {
					tempVar := ((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(bptr)) + (uintptr)(int32(uint32((dbuffer_size))))*unsafe.Sizeof(*bptr))))
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) - (uintptr)(len)*unsafe.Sizeof(*tempVar))
					}()))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) - (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}()))), unsafe.Pointer(bptr), int32(uint32((len + int32(1)))))
				bptr = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(bptr)) + (uintptr)(int32(uint32(((dbuffer_size-size_t((uint32(len))))-size_t((uint32(int32(1))))))))*unsafe.Sizeof(*bptr))))
			} else {
				bptr = (*pcre_uint8)(noarch.Memcpy(unsafe.Pointer(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(bptr)) + (uintptr)(int32(uint32(((dbuffer_size-size_t((uint32(len))))*size_t((uint32(int32(1))))))))*unsafe.Sizeof(*bptr))))), unsafe.Pointer(bptr), int32(uint32((len * int32(1))))))
			}
			if ((all_use_dfa != 0) || (use_dfa != 0)) && (find_match_limit != 0) {
				noarch.Printf((&[]byte("** Match limit not relevant for DFA matching: ignored\n\x00")[0]))
				find_match_limit = int32(0)
			}
			if (posix != 0) || (do_posix != 0) {
				var rc int32
				var eflags int32 = int32(0)
				var pmatch *regmatch_t = nil
				if use_size_offsets > int32(0) {
					pmatch = (*regmatch_t)(noarch.Malloc(int32((8 * uint32(use_size_offsets)))))
				}
				if (options & int32(128)) != int32(0) {
					eflags |= int32(4)
				}
				if (options & int32(256)) != int32(0) {
					eflags |= int32(8)
				}
				if (options & int32(1024)) != int32(0) {
					eflags |= int32(256)
				}
				rc = regexec(&preg, (*byte)(unsafe.Pointer(bptr)), size_t(use_size_offsets), pmatch, eflags)
				if rc != int32(0) {
					_ = regerror(rc, &preg, (*byte)(unsafe.Pointer(buffer)), size_t(buffer_size))
					noarch.Fprintf(outfile, (&[]byte("No match: POSIX code %d: %s\n\x00")[0]), rc, buffer)
				} else if ((*((*real_pcre)(preg.re_pcre))).options & pcre_uint32((uint32(int32(4096))))) != pcre_uint32((uint32(int32(0)))) {
					noarch.Fprintf(outfile, (&[]byte("Matched with REG_NOSUB\n\x00")[0]))
				} else {
					var i size_t
					for i = size_t(int32(0)); i < size_t(use_size_offsets); i++ {
						if regoff_t((*((*regmatch_t)(unsafe.Pointer(uintptr(unsafe.Pointer(pmatch)) + (uintptr)(int32(uint32((i))))*unsafe.Sizeof(*pmatch))))).rm_so) >= regoff_t((int32(0))) {
							noarch.Fprintf(outfile, (&[]byte("%2d: \x00")[0]), int32(uint32((size_t(i)))))
							_ = pchars(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((dbuffer))) + (uintptr)(int32((regoff_t((*((*regmatch_t)(unsafe.Pointer(uintptr(unsafe.Pointer(pmatch)) + (uintptr)(int32(uint32((i))))*unsafe.Sizeof(*pmatch))))).rm_so))))*unsafe.Sizeof(*(dbuffer))))), int32((regoff_t((*((*regmatch_t)(unsafe.Pointer(uintptr(unsafe.Pointer(pmatch)) + (uintptr)(int32(uint32((i))))*unsafe.Sizeof(*pmatch))))).rm_eo) - regoff_t((*((*regmatch_t)(unsafe.Pointer(uintptr(unsafe.Pointer(pmatch)) + (uintptr)(int32(uint32((i))))*unsafe.Sizeof(*pmatch))))).rm_so))), outfile)
							noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
							if (do_showcaprest != 0) || ((i == size_t((uint32(int32(0))))) && (do_showrest != 0)) {
								noarch.Fprintf(outfile, (&[]byte("%2d+ \x00")[0]), int32(uint32((size_t(i)))))
								_ = pchars(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((dbuffer))) + (uintptr)(int32((regoff_t((*((*regmatch_t)(unsafe.Pointer(uintptr(unsafe.Pointer(pmatch)) + (uintptr)(int32(uint32((i))))*unsafe.Sizeof(*pmatch))))).rm_eo))))*unsafe.Sizeof(*(dbuffer))))), (len - int32((regoff_t((*((*regmatch_t)(unsafe.Pointer(uintptr(unsafe.Pointer(pmatch)) + (uintptr)(int32(uint32((i))))*unsafe.Sizeof(*pmatch))))).rm_eo)))), outfile)
								noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
							}
						}
					}
				}
				noarch.Free(unsafe.Pointer(pmatch))
				goto NEXT_DATA
			}
			if ((verify_jit != 0) && (jit_stack == nil)) && (extra != nil) {
				pcre_assign_jit_stack(extra, jit_callback, unsafe.Pointer(jit_stack))
			}
			for ; ; gmatched++ {
				markptr = nil
				jit_was_used = int32(0)
				if timeitm > int32(0) {
					var i int32
					var time_taken clock_t
					var start_time clock_t = clock()
					if (all_use_dfa != 0) || (use_dfa != 0) {
						if (options & int32(131072)) != int32(0) {
							noarch.Fprintf(outfile, (&[]byte("Timing DFA restarts is not supported\n\x00")[0]))
							break
						}
						if dfa_workspace == nil {
							dfa_workspace = (*int32)(noarch.Malloc(int32((uint32(int32(1000)) * 4))))
						}
						for i = int32(0); i < timeitm; i++ {
							count = pcre_dfa_exec(re, extra, (*byte)(unsafe.Pointer(bptr)), len, start_offset, (options | g_notempty), use_offsets, use_size_offsets, dfa_workspace, int32(1000))
						}
					} else {
						for i = int32(0); i < timeitm; i++ {
							count = pcre_exec(re, extra, (*byte)(unsafe.Pointer(bptr)), len, start_offset, (options | g_notempty), use_offsets, use_size_offsets)
						}
					}
					total_match_time += (func() clock_t {
						tempVar := (clock() - start_time)
						time_taken = tempVar
						return tempVar
					}())
					noarch.Fprintf(outfile, (&[]byte("Execute time %.4f milliseconds\n\x00")[0]), (((float64(int32((__clock_t((clock_t(time_taken)))))) * 1000) / float64(timeitm)) / float64(int32((__clock_t((clock_t(int32(1000000)))))))))
				}
				if find_match_limit != 0 {
					if extra != nil {
						pcre_free_study(extra)
					}
					extra = (*pcre_extra)(noarch.Malloc(int32(64)))
					(*extra).flags = uint32(int32(0))
					_ = check_match_limit(re, extra, bptr, len, start_offset, (options | g_notempty), use_offsets, use_size_offsets, int32(2), &((*extra).match_limit), (-int32(8)), (&[]byte("match()\x00")[0]))
					count = check_match_limit(re, extra, bptr, len, start_offset, (options | g_notempty), use_offsets, use_size_offsets, int32(16), &((*extra).match_limit_recursion), (-int32(21)), (&[]byte("match() recursion\x00")[0]))
				} else if callout_data_set != 0 {
					if extra == nil {
						extra = (*pcre_extra)(noarch.Malloc(int32(64)))
						(*extra).flags = uint32(int32(0))
					}
					(*extra).flags |= uint32(int32(4))
					(*extra).callout_data = unsafe.Pointer(&callout_data)
					count = pcre_exec(re, extra, (*byte)(unsafe.Pointer(bptr)), len, start_offset, (options | g_notempty), use_offsets, use_size_offsets)
					(*extra).flags &= ^uint32(int32(4))
				} else if (all_use_dfa != 0) || (use_dfa != 0) {
					if dfa_workspace == nil {
						dfa_workspace = (*int32)(noarch.Malloc(int32((uint32(int32(1000)) * 4))))
					}
					if func() int32 {
						defer func() {
							dfa_matched += 1
						}()
						return dfa_matched
					}() == int32(0) {
						*dfa_workspace = -int32(1)
					}
					count = pcre_dfa_exec(re, extra, (*byte)(unsafe.Pointer(bptr)), len, start_offset, (options | g_notempty), use_offsets, use_size_offsets, dfa_workspace, int32(1000))
					if count == int32(0) {
						noarch.Fprintf(outfile, (&[]byte("Matched, but offsets vector is too small to show all matches\n\x00")[0]))
						count = (use_size_offsets / int32(2))
					}
				} else {
					count = pcre_exec(re, extra, (*byte)(unsafe.Pointer(bptr)), len, start_offset, (options | g_notempty), use_offsets, use_size_offsets)
					if count == int32(0) {
						noarch.Fprintf(outfile, (&[]byte("Matched, but too many substrings\n\x00")[0]))
						count = func() int32 {
							if (map[bool]int32{false: 0, true: 1}[(use_size_offsets == int32(2))]) != 0 {
								return int32(1)
							} else {
								return (use_size_offsets / int32(3))
							}
						}()
					}
				}
				if count >= int32(0) {
					var i int32
					var maxcount int32
					var cnptr unsafe.Pointer
					var gnptr unsafe.Pointer
					if (all_use_dfa != 0) || (use_dfa != 0) {
						maxcount = (use_size_offsets / int32(2))
					} else {
						maxcount = func() int32 {
							if (map[bool]int32{false: 0, true: 1}[(use_size_offsets == int32(2))]) != 0 {
								return int32(1)
							} else {
								return (use_size_offsets / int32(3))
							}
						}()
					}
					if count > maxcount {
						noarch.Fprintf(outfile, (&[]byte("** PCRE error: returned count %d is too big for offset size %d\n\x00")[0]), count, use_size_offsets)
						count = (use_size_offsets / int32(3))
						if (do_g != 0) || (do_G != 0) {
							noarch.Fprintf(outfile, (&[]byte("** /%c loop abandoned\n\x00")[0]), func() int32 {
								if do_g != 0 {
									return int32('g')
								} else {
									return int32('G')
								}
							}())
							do_G = int32(0)
							do_g = do_G
						}
					}
					if do_allcaps != 0 {
						if (all_use_dfa != 0) || (use_dfa != 0) {
							noarch.Fprintf(outfile, (&[]byte("** Show all captures ignored after DFA matching\n\x00")[0]))
						} else {
							if new_info(re, nil, int32(2), unsafe.Pointer(&count)) < int32(0) {
								goto_SKIP_DATA = true
								goto SKIP_DATA_CONTAINER_OUTER
							}
							count += 1
							if (count * int32(2)) > use_size_offsets {
								count = (use_size_offsets / int32(2))
							}
						}
					}
					for i = int32(0); i < (count * int32(2)); i += int32(2) {
						if *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(i)*unsafe.Sizeof(*use_offsets)))) < int32(0) {
							if *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(i)*unsafe.Sizeof(*use_offsets)))) != -int32(1) {
								noarch.Fprintf(outfile, (&[]byte("ERROR: bad negative value %d for offset %d\n\x00")[0]), *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(i)*unsafe.Sizeof(*use_offsets)))), i)
							}
							if *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)((i+int32(1)))*unsafe.Sizeof(*use_offsets)))) != -int32(1) {
								noarch.Fprintf(outfile, (&[]byte("ERROR: bad negative value %d for offset %d\n\x00")[0]), *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)((i+int32(1)))*unsafe.Sizeof(*use_offsets)))), (i + int32(1)))
							}
							noarch.Fprintf(outfile, (&[]byte("%2d: <unset>\n\x00")[0]), (i / int32(2)))
						} else {
							var start int32 = *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(i)*unsafe.Sizeof(*use_offsets))))
							var end int32 = *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)((i+int32(1)))*unsafe.Sizeof(*use_offsets))))
							if start > end {
								start = *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)((i+int32(1)))*unsafe.Sizeof(*use_offsets))))
								end = *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(i)*unsafe.Sizeof(*use_offsets))))
								noarch.Fprintf(outfile, (&[]byte("Start of matched string is beyond its end - displaying from end to start.\n\x00")[0]))
							}
							noarch.Fprintf(outfile, (&[]byte("%2d: \x00")[0]), (i / int32(2)))
							_ = pchars(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((bptr))) + (uintptr)(start)*unsafe.Sizeof(*(bptr))))), (end - start), outfile)
							if (verify_jit != 0) && (jit_was_used != 0) {
								noarch.Fprintf(outfile, (&[]byte(" (JIT)\x00")[0]))
							}
							noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
							if (do_showcaprest != 0) || ((i == int32(0)) && (do_showrest != 0)) {
								noarch.Fprintf(outfile, (&[]byte("%2d+ \x00")[0]), (i / int32(2)))
								_ = pchars(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((bptr))) + (uintptr)(*((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)((i+int32(1)))*unsafe.Sizeof(*use_offsets)))))*unsafe.Sizeof(*(bptr))))), (len - *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)((i+int32(1)))*unsafe.Sizeof(*use_offsets))))), outfile)
								noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
							}
						}
					}
					if markptr != nil {
						noarch.Fprintf(outfile, (&[]byte("MK: \x00")[0]))
						_ = pchars(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((markptr))) + (uintptr)(int32(0))*unsafe.Sizeof(*(markptr))))), -int32(1), outfile)
						noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
					}
					for i = int32(0); i < int32(32); i++ {
						if (copystrings & (int32(1) << uint64(i))) != int32(0) {
							var rc int32
							var copybuffer []byte = make([]byte, 256, 256)
							rc = pcre_copy_substring((*byte)(unsafe.Pointer(bptr)), use_offsets, count, i, &copybuffer[0], int32(256))
							if rc < int32(0) {
								noarch.Fprintf(outfile, (&[]byte("copy substring %d failed %d\n\x00")[0]), i, rc)
							} else {
								noarch.Fprintf(outfile, (&[]byte("%2dC \x00")[0]), i)
								_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
									tempVar := (*pcre_uint8)(unsafe.Pointer(&(copybuffer)[0]))
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
								}())), rc, outfile)
								noarch.Fprintf(outfile, (&[]byte(" (%d)\n\x00")[0]), rc)
							}
						}
					}
					cnptr = unsafe.Pointer(&copynames[0])
					for {
						var rc int32
						var copybuffer []byte = make([]byte, 256, 256)
						if pcre_mode == PCRE8_MODE {
							if int32(uint8((*(*pcre_uint8)(cnptr)))) == int32(0) {
								break
							}
						}
						rc = pcre_copy_named_substring(re, (*byte)(unsafe.Pointer(bptr)), use_offsets, count, (*byte)(cnptr), &copybuffer[0], int32(256))
						if rc < int32(0) {
							noarch.Fprintf(outfile, (&[]byte("copy substring \x00")[0]))
							_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
								tempVar := (*pcre_uint8)((cnptr))
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
							}())), -int32(1), outfile)
							noarch.Fprintf(outfile, (&[]byte(" failed %d\n\x00")[0]), rc)
						} else {
							noarch.Fprintf(outfile, (&[]byte("  C \x00")[0]))
							_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
								tempVar := (*pcre_uint8)(unsafe.Pointer(&(copybuffer)[0]))
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
							}())), rc, outfile)
							noarch.Fprintf(outfile, (&[]byte(" (%d) \x00")[0]), rc)
							_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
								tempVar := (*pcre_uint8)((cnptr))
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
							}())), -int32(1), outfile)
							noarch.Fputc(int32('\n'), outfile)
						}
						cnptr = unsafe.Pointer(((*byte)(func() unsafe.Pointer {
							tempVar := (*byte)(cnptr)
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(((noarch.Strlen((*byte)(cnptr))+int32(1))*int32(1)))*unsafe.Sizeof(*tempVar))
						}())))
					}
					for i = int32(0); i < int32(32); i++ {
						if (getstrings & (int32(1) << uint64(i))) != int32(0) {
							var rc int32
							var substring *byte
							rc = pcre_get_substring((*byte)(unsafe.Pointer(bptr)), use_offsets, count, i, &substring)
							if rc < int32(0) {
								noarch.Fprintf(outfile, (&[]byte("get substring %d failed %d\n\x00")[0]), i, rc)
							} else {
								noarch.Fprintf(outfile, (&[]byte("%2dG \x00")[0]), i)
								_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
									tempVar := (*pcre_uint8)(unsafe.Pointer((substring)))
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
								}())), rc, outfile)
								noarch.Fprintf(outfile, (&[]byte(" (%d)\n\x00")[0]), rc)
								pcre_free_substring(substring)
							}
						}
					}
					gnptr = unsafe.Pointer(&getnames[0])
					for {
						var rc int32
						var substring *byte
						if pcre_mode == PCRE8_MODE {
							if int32(uint8((*(*pcre_uint8)(gnptr)))) == int32(0) {
								break
							}
						}
						rc = pcre_get_named_substring(re, (*byte)(unsafe.Pointer(bptr)), use_offsets, count, (*byte)(gnptr), &substring)
						if rc < int32(0) {
							noarch.Fprintf(outfile, (&[]byte("get substring \x00")[0]))
							_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
								tempVar := (*pcre_uint8)((gnptr))
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
							}())), -int32(1), outfile)
							noarch.Fprintf(outfile, (&[]byte(" failed %d\n\x00")[0]), rc)
						} else {
							noarch.Fprintf(outfile, (&[]byte("  G \x00")[0]))
							_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
								tempVar := (*pcre_uint8)(unsafe.Pointer((substring)))
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
							}())), rc, outfile)
							noarch.Fprintf(outfile, (&[]byte(" (%d) \x00")[0]), rc)
							_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
								tempVar := (*pcre_uint8)((gnptr))
								return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
							}())), -int32(1), outfile)
							pcre_free_substring(substring)
							noarch.Fputc(int32('\n'), outfile)
						}
						gnptr = unsafe.Pointer(((*byte)(func() unsafe.Pointer {
							tempVar := (*byte)(gnptr)
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(((noarch.Strlen((*byte)(gnptr))+int32(1))*int32(1)))*unsafe.Sizeof(*tempVar))
						}())))
					}
					if getlist != 0 {
						var rc int32
						var stringlist **byte
						rc = pcre_get_substring_list((*byte)(unsafe.Pointer(bptr)), use_offsets, count, &stringlist)
						if rc < int32(0) {
							noarch.Fprintf(outfile, (&[]byte("get substring list failed %d\n\x00")[0]), rc)
						} else {
							for i = int32(0); i < count; i++ {
								noarch.Fprintf(outfile, (&[]byte("%2dL \x00")[0]), i)
								_ = pchars(((*pcre_uint8)(func() unsafe.Pointer {
									tempVar := (*pcre_uint8)(unsafe.Pointer((*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(stringlist)) + (uintptr)(i)*unsafe.Sizeof(*stringlist)))))))
									return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(0))*unsafe.Sizeof(*tempVar))
								}())), -int32(1), outfile)
								noarch.Fputc(int32('\n'), outfile)
							}
							if *((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(stringlist)) + (uintptr)(i)*unsafe.Sizeof(*stringlist)))) != nil {
								noarch.Fprintf(outfile, (&[]byte("string list not terminated by NULL\n\x00")[0]))
							}
							pcre_free_substring_list(stringlist)
						}
					}
				} else if count == -int32(12) {
					noarch.Fprintf(outfile, (&[]byte("Partial match\x00")[0]))
					if (use_size_offsets > int32(2)) && (*use_offsets != *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(2))*unsafe.Sizeof(*use_offsets))))) {
						noarch.Fprintf(outfile, (&[]byte(" at offset %d\x00")[0]), *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(2))*unsafe.Sizeof(*use_offsets)))))
					}
					if markptr != nil {
						noarch.Fprintf(outfile, (&[]byte(", mark=\x00")[0]))
						_ = pchars(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((markptr))) + (uintptr)(int32(0))*unsafe.Sizeof(*(markptr))))), -int32(1), outfile)
					}
					if use_size_offsets > int32(1) {
						noarch.Fprintf(outfile, (&[]byte(": \x00")[0]))
						_ = pchars(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((bptr))) + (uintptr)(*use_offsets)*unsafe.Sizeof(*(bptr))))), (*((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(1))*unsafe.Sizeof(*use_offsets)))) - *use_offsets), outfile)
					}
					if (verify_jit != 0) && (jit_was_used != 0) {
						noarch.Fprintf(outfile, (&[]byte(" (JIT)\x00")[0]))
					}
					noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
					break
				} else {
					if g_notempty != int32(0) {
						var onechar int32 = int32(1)
						var obits uint32 = uint32(((*((*real_pcre)(unsafe.Pointer(re)))).options))
						*use_offsets = start_offset
						if (obits & uint32((((int32(1048576) | int32(2097152)) | int32(4194304)) | int32(5242880)))) == uint32(int32(0)) {
							var d int32
							_ = pcre_config(int32(1), unsafe.Pointer(&d))
							obits = uint32(func() int32 {
								if (map[bool]int32{false: 0, true: 1}[(d == int32(13))]) != 0 {
									return int32(1048576)
								} else {
									return func() int32 {
										if (map[bool]int32{false: 0, true: 1}[(d == int32(10))]) != 0 {
											return int32(2097152)
										} else {
											return func() int32 {
												if (map[bool]int32{false: 0, true: 1}[(d == ((int32(13) << uint64(int32(8))) | int32(10)))]) != 0 {
													return int32(3145728)
												} else {
													return func() int32 {
														if (map[bool]int32{false: 0, true: 1}[(d == -int32(2))]) != 0 {
															return int32(5242880)
														} else {
															return func() int32 {
																if (map[bool]int32{false: 0, true: 1}[(d == -int32(1))]) != 0 {
																	return int32(4194304)
																} else {
																	return int32(0)
																}
															}()
														}
													}()
												}
											}()
										}
									}()
								}
							}())
						}
						if (((((obits & uint32((((int32(1048576) | int32(2097152)) | int32(4194304)) | int32(5242880)))) == uint32(int32(4194304))) || ((obits & uint32((((int32(1048576) | int32(2097152)) | int32(4194304)) | int32(5242880)))) == uint32(int32(3145728)))) || ((obits & uint32((((int32(1048576) | int32(2097152)) | int32(4194304)) | int32(5242880)))) == uint32(int32(5242880)))) && (start_offset < (len - int32(1)))) && ((((pcre_mode == PCRE8_MODE) && (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(bptr)) + (uintptr)(start_offset)*unsafe.Sizeof(*bptr))))))) == int32('\r'))) && (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(bptr)) + (uintptr)((start_offset+int32(1)))*unsafe.Sizeof(*bptr))))))) == int32('\n'))) || (int32(0) != 0)) {
							onechar += 1
						} else if use_utf != 0 {
							for (start_offset + onechar) < len {
								if (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(bptr)) + (uintptr)((start_offset+onechar))*unsafe.Sizeof(*bptr))))))) & int32(192)) != int32(128) {
									break
								}
								onechar += 1
							}
						}
						*((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(1))*unsafe.Sizeof(*use_offsets)))) = (start_offset + onechar)
					} else {
						switch count {
						case (-int32(1)):
							{
								if gmatched == int32(0) {
									if markptr == nil {
										noarch.Fprintf(outfile, (&[]byte("No match\x00")[0]))
									} else {
										noarch.Fprintf(outfile, (&[]byte("No match, mark = \x00")[0]))
										_ = pchars(((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((markptr))) + (uintptr)(int32(0))*unsafe.Sizeof(*(markptr))))), -int32(1), outfile)
									}
									if (verify_jit != 0) && (jit_was_used != 0) {
										noarch.Fprintf(outfile, (&[]byte(" (JIT)\x00")[0]))
									}
									noarch.Fputc(int32('\n'), outfile)
								}
							}
						case (-int32(10)):
							fallthrough
						case (-int32(25)):
							{
								noarch.Fprintf(outfile, (&[]byte("Error %d (%s UTF-%d string)\x00")[0]), count, func() *byte {
									if (map[bool]int32{false: 0, true: 1}[(count == -int32(10))]) != 0 {
										return (&[]byte("bad\x00")[0])
									} else {
										return (&[]byte("short\x00")[0])
									}
								}(), (int32(8) * int32(1)))
								if use_size_offsets >= int32(2) {
									noarch.Fprintf(outfile, (&[]byte(" offset=%d reason=%d\x00")[0]), *use_offsets, *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(1))*unsafe.Sizeof(*use_offsets)))))
								}
								noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
							}
						case (-int32(11)):
							{
								noarch.Fprintf(outfile, (&[]byte("Error %d (bad UTF-%d offset)\n\x00")[0]), count, (int32(8) * int32(1)))
							}
						default:
							{
								if (count < int32(0)) && (-count < int32((264 / 8))) {
									noarch.Fprintf(outfile, (&[]byte("Error %d (%s)\n\x00")[0]), count, *((**byte)(func() unsafe.Pointer {
										tempVar := &errtexts[0]
										return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(-count)*unsafe.Sizeof(*tempVar))
									}())))
								} else {
									noarch.Fprintf(outfile, (&[]byte("Error %d (Unexpected value)\n\x00")[0]), count)
								}
								break
							}
						}
						break
					}
				}
				if (noarch.NotInt32(do_g) != 0) && (noarch.NotInt32(do_G) != 0) {
					break
				}
				if use_offsets == nil {
					noarch.Fprintf(outfile, (&[]byte("Cannot do global matching without an ovector\n\x00")[0]))
					break
				}
				if use_size_offsets < int32(2) {
					noarch.Fprintf(outfile, (&[]byte("Cannot do global matching with an ovector size < 2\n\x00")[0]))
					break
				}
				g_notempty = int32(0)
				if *use_offsets == *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(1))*unsafe.Sizeof(*use_offsets)))) {
					if *use_offsets == len {
						break
					}
					g_notempty = (int32(268435456) | int32(16))
				}
				if do_g != 0 {
					if (g_notempty == int32(0)) && (*((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(1))*unsafe.Sizeof(*use_offsets)))) <= start_offset) {
						if start_offset >= len {
							break
						}
						start_offset += 1
						if use_utf != 0 {
							for start_offset < len {
								if (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(bptr)) + (uintptr)(start_offset)*unsafe.Sizeof(*bptr))))))) & int32(192)) != int32(128) {
									break
								}
								start_offset += 1
							}
						}
					} else {
						start_offset = *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(1))*unsafe.Sizeof(*use_offsets))))
					}
				} else {
					bptr = ((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(bptr)) + (uintptr)((*((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(1))*unsafe.Sizeof(*use_offsets))))*int32(1)))*unsafe.Sizeof(*bptr))))
					len -= *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(use_offsets)) + (uintptr)(int32(1))*unsafe.Sizeof(*use_offsets))))
				}
			}
		NEXT_DATA:
			;
			continue
		}
	CONTINUE:
		;
		if ((posix != 0) || (do_posix != 0)) && (preg.re_pcre != nil) {
			regfree(&preg)
		}
		if re != nil {
			new_free(unsafe.Pointer(re))
		}
		if extra != nil {
			pcre_free_study(extra)
		}
		if locale_set != 0 {
			new_free(unsafe.Pointer(tables))
			setlocale(int32(0), (&[]byte("C\x00")[0]))
			locale_set = int32(0)
		}
		if jit_stack != nil {
			pcre_jit_stack_free(jit_stack)
			jit_stack = nil
		}
	}
	if int64(uintptr(unsafe.Pointer(infile))) == int64(uintptr(unsafe.Pointer(stdin))) {
		noarch.Fprintf(outfile, (&[]byte("\n\x00")[0]))
	}
	if showtotaltimes != 0 {
		noarch.Fprintf(outfile, (&[]byte("--------------------------------------\n\x00")[0]))
		if timeit > int32(0) {
			noarch.Fprintf(outfile, (&[]byte("Total compile time %.4f milliseconds\n\x00")[0]), (((float64(int32((__clock_t((clock_t(total_compile_time)))))) * 1000) / float64(timeit)) / float64(int32((__clock_t((clock_t(int32(1000000)))))))))
			noarch.Fprintf(outfile, (&[]byte("Total study time   %.4f milliseconds\n\x00")[0]), (((float64(int32((__clock_t((clock_t(total_study_time)))))) * 1000) / float64(timeit)) / float64(int32((__clock_t((clock_t(int32(1000000)))))))))
		}
		noarch.Fprintf(outfile, (&[]byte("Total execute time %.4f milliseconds\n\x00")[0]), (((float64(int32((__clock_t((clock_t(total_match_time)))))) * 1000) / float64(timeitm)) / float64(int32((__clock_t((clock_t(int32(1000000)))))))))
	}
EXIT:
	;
	if (infile != nil) && (int64(uintptr(unsafe.Pointer(infile))) != int64(uintptr(unsafe.Pointer(stdin)))) {
		noarch.Fclose(infile)
	}
	if (outfile != nil) && (int64(uintptr(unsafe.Pointer(outfile))) != int64(uintptr(unsafe.Pointer(stdout)))) {
		noarch.Fclose(outfile)
	}
	noarch.Free(unsafe.Pointer(buffer))
	noarch.Free(unsafe.Pointer(dbuffer))
	noarch.Free(unsafe.Pointer(pbuffer))
	noarch.Free(unsafe.Pointer(offsets))
	if dfa_workspace != nil {
		noarch.Free(unsafe.Pointer(dfa_workspace))
	}
	os.Exit(int(yield))
}

var priv_OP_names []*byte = []*byte{(&[]byte("End\x00")[0]), (&[]byte("\\A\x00")[0]), (&[]byte("\\G\x00")[0]), (&[]byte("\\K\x00")[0]), (&[]byte("\\B\x00")[0]), (&[]byte("\\b\x00")[0]), (&[]byte("\\D\x00")[0]), (&[]byte("\\d\x00")[0]), (&[]byte("\\S\x00")[0]), (&[]byte("\\s\x00")[0]), (&[]byte("\\W\x00")[0]), (&[]byte("\\w\x00")[0]), (&[]byte("Any\x00")[0]), (&[]byte("AllAny\x00")[0]), (&[]byte("Anybyte\x00")[0]), (&[]byte("notprop\x00")[0]), (&[]byte("prop\x00")[0]), (&[]byte("\\R\x00")[0]), (&[]byte("\\H\x00")[0]), (&[]byte("\\h\x00")[0]), (&[]byte("\\V\x00")[0]), (&[]byte("\\v\x00")[0]), (&[]byte("extuni\x00")[0]), (&[]byte("\\Z\x00")[0]), (&[]byte("\\z\x00")[0]), (&[]byte("$\x00")[0]), (&[]byte("$\x00")[0]), (&[]byte("^\x00")[0]), (&[]byte("^\x00")[0]), (&[]byte("char\x00")[0]), (&[]byte("chari\x00")[0]), (&[]byte("not\x00")[0]), (&[]byte("noti\x00")[0]), (&[]byte("*\x00")[0]), (&[]byte("*?\x00")[0]), (&[]byte("+\x00")[0]), (&[]byte("+?\x00")[0]), (&[]byte("?\x00")[0]), (&[]byte("??\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*+\x00")[0]), (&[]byte("++\x00")[0]), (&[]byte("?+\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*\x00")[0]), (&[]byte("*?\x00")[0]), (&[]byte("+\x00")[0]), (&[]byte("+?\x00")[0]), (&[]byte("?\x00")[0]), (&[]byte("??\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*+\x00")[0]), (&[]byte("++\x00")[0]), (&[]byte("?+\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*\x00")[0]), (&[]byte("*?\x00")[0]), (&[]byte("+\x00")[0]), (&[]byte("+?\x00")[0]), (&[]byte("?\x00")[0]), (&[]byte("??\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*+\x00")[0]), (&[]byte("++\x00")[0]), (&[]byte("?+\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*\x00")[0]), (&[]byte("*?\x00")[0]), (&[]byte("+\x00")[0]), (&[]byte("+?\x00")[0]), (&[]byte("?\x00")[0]), (&[]byte("??\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*+\x00")[0]), (&[]byte("++\x00")[0]), (&[]byte("?+\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*\x00")[0]), (&[]byte("*?\x00")[0]), (&[]byte("+\x00")[0]), (&[]byte("+?\x00")[0]), (&[]byte("?\x00")[0]), (&[]byte("??\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*+\x00")[0]), (&[]byte("++\x00")[0]), (&[]byte("?+\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*\x00")[0]), (&[]byte("*?\x00")[0]), (&[]byte("+\x00")[0]), (&[]byte("+?\x00")[0]), (&[]byte("?\x00")[0]), (&[]byte("??\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("*+\x00")[0]), (&[]byte("++\x00")[0]), (&[]byte("?+\x00")[0]), (&[]byte("{\x00")[0]), (&[]byte("class\x00")[0]), (&[]byte("nclass\x00")[0]), (&[]byte("xclass\x00")[0]), (&[]byte("Ref\x00")[0]), (&[]byte("Refi\x00")[0]), (&[]byte("DnRef\x00")[0]), (&[]byte("DnRefi\x00")[0]), (&[]byte("Recurse\x00")[0]), (&[]byte("Callout\x00")[0]), (&[]byte("Alt\x00")[0]), (&[]byte("Ket\x00")[0]), (&[]byte("KetRmax\x00")[0]), (&[]byte("KetRmin\x00")[0]), (&[]byte("KetRpos\x00")[0]), (&[]byte("Reverse\x00")[0]), (&[]byte("Assert\x00")[0]), (&[]byte("Assert not\x00")[0]), (&[]byte("AssertB\x00")[0]), (&[]byte("AssertB not\x00")[0]), (&[]byte("Once\x00")[0]), (&[]byte("Once_NC\x00")[0]), (&[]byte("Bra\x00")[0]), (&[]byte("BraPos\x00")[0]), (&[]byte("CBra\x00")[0]), (&[]byte("CBraPos\x00")[0]), (&[]byte("Cond\x00")[0]), (&[]byte("SBra\x00")[0]), (&[]byte("SBraPos\x00")[0]), (&[]byte("SCBra\x00")[0]), (&[]byte("SCBraPos\x00")[0]), (&[]byte("SCond\x00")[0]), (&[]byte("Cond ref\x00")[0]), (&[]byte("Cond dnref\x00")[0]), (&[]byte("Cond rec\x00")[0]), (&[]byte("Cond dnrec\x00")[0]), (&[]byte("Cond def\x00")[0]), (&[]byte("Brazero\x00")[0]), (&[]byte("Braminzero\x00")[0]), (&[]byte("Braposzero\x00")[0]), (&[]byte("*MARK\x00")[0]), (&[]byte("*PRUNE\x00")[0]), (&[]byte("*PRUNE\x00")[0]), (&[]byte("*SKIP\x00")[0]), (&[]byte("*SKIP\x00")[0]), (&[]byte("*THEN\x00")[0]), (&[]byte("*THEN\x00")[0]), (&[]byte("*COMMIT\x00")[0]), (&[]byte("*FAIL\x00")[0]), (&[]byte("*ACCEPT\x00")[0]), (&[]byte("*ASSERT_ACCEPT\x00")[0]), (&[]byte("Close\x00")[0]), (&[]byte("Skip zero\x00")[0])}
var priv_OP_lengths []pcre_uint8 = []pcre_uint8{pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8(int32(2)), pcre_uint8((int32(2) + int32(2))), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((uint32(int32(1)) + (uint32(int32(32)) / 1))), pcre_uint8((uint32(int32(1)) + (uint32(int32(32)) / 1))), pcre_uint8(int32(0)), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(2) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8(((int32(1) + int32(2)) + int32(2))), pcre_uint8(((int32(1) + int32(2)) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8(((int32(1) + int32(2)) + int32(2))), pcre_uint8(((int32(1) + int32(2)) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8((int32(1) + int32(2))), pcre_uint8((int32(1) + (int32(2) * int32(2)))), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(3)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8(int32(1)), pcre_uint8((int32(1) + int32(2))), pcre_uint8(int32(1))}

func print_char(f *noarch.File, ptr *pcre_uchar, utf BOOL) (c2goDefaultReturn uint32) {
	var c pcre_uint32 = pcre_uint32((uint32(uint8((*ptr)))))
	if (int32((NotBOOL(BOOL(utf)))) != 0) || ((c & pcre_uint32((uint32(int32(192))))) != pcre_uint32((uint32(int32(192))))) {
		if (map[bool]int32{false: 0, true: 1}[((c >= pcre_uint32((uint32(int32(32))))) && (c < pcre_uint32((uint32(int32(127))))))]) != 0 {
			noarch.Fprintf(f, (&[]byte("%c\x00")[0]), int32(byte(uint32((pcre_uint32(c))))))
		} else if c < pcre_uint32((uint32(int32(128)))) {
			noarch.Fprintf(f, (&[]byte("\\x%02x\x00")[0]), pcre_uint32(c))
		} else {
			noarch.Fprintf(f, (&[]byte("\\x{%02x}\x00")[0]), pcre_uint32(c))
		}
		return uint32(int32(0))
	} else {
		var i int32
		var a int32 = int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
			tempVar := &utf8_table4[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((c&pcre_uint32((uint32(int32(63))))))))*unsafe.Sizeof(*tempVar))
		}())))))
		var s int32 = (int32(6) * a)
		c = ((c & pcre_uint32((uint32(*((*int32)(func() unsafe.Pointer {
			tempVar := &utf8_table3[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(a)*unsafe.Sizeof(*tempVar))
		}())))))) << uint64(s))
		for i = int32(1); i <= a; i++ {
			if (int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + (uintptr)(i)*unsafe.Sizeof(*ptr))))))) & int32(192)) != int32(128) {
				noarch.Fprintf(f, (&[]byte("\\X{%x}\x00")[0]), pcre_uint32(c))
				return uint32((i - int32(1)))
			}
			s -= int32(6)
			c |= pcre_uint32((uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + (uintptr)(i)*unsafe.Sizeof(*ptr))))))) & int32(63)) << uint64(s)))))
		}
		noarch.Fprintf(f, (&[]byte("\\x{%x}\x00")[0]), pcre_uint32(c))
		return uint32(a)
	}
	return
}

func print_puchar(f *noarch.File, ptr *pcre_uchar) {
	for int32(uint8((*ptr))) != int32('\x00') {
		var c pcre_uint32 = pcre_uint32((uint32(uint8((*func() *pcre_uchar {
			defer func() {
				ptr = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + (uintptr)(1)*unsafe.Sizeof(*ptr))))
			}()
			return ptr
		}())))))
		if (map[bool]int32{false: 0, true: 1}[((c >= pcre_uint32((uint32(int32(32))))) && (c < pcre_uint32((uint32(int32(127))))))]) != 0 {
			noarch.Fprintf(f, (&[]byte("%c\x00")[0]), pcre_uint32(c))
		} else {
			noarch.Fprintf(f, (&[]byte("\\x{%x}\x00")[0]), pcre_uint32(c))
		}
	}
}

func get_ucpname(ptype uint32, pvalue uint32) *byte {
	var i int32
	for i = (utt_size - int32(1)); i >= int32(0); i-- {
		if (ptype == uint32(uint16((pcre_uint16((*((*ucp_type_table)(func() unsafe.Pointer {
			tempVar := &utt[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
		}()))).type_))))) && (pvalue == uint32(uint16((pcre_uint16((*((*ucp_type_table)(func() unsafe.Pointer {
			tempVar := &utt[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
		}()))).value))))) {
			break
		}
	}
	return func() *byte {
		if (map[bool]int32{false: 0, true: 1}[(i >= int32(0))]) != 0 {
			return ((*byte)(func() unsafe.Pointer {
				tempVar := &utt_names[0]
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint16((pcre_uint16((*((*ucp_type_table)(func() unsafe.Pointer {
					tempVar := &utt[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
				}()))).name_offset)))))*unsafe.Sizeof(*tempVar))
			}()))
		} else {
			return (&[]byte("??\x00")[0])
		}
	}()
}

func print_prop(f *noarch.File, code *pcre_uchar, before *byte, after *byte) {
	if int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code))))))) != int32(9) {
		noarch.Fprintf(f, (&[]byte("%s%s %s%s\x00")[0]), before, *((**byte)(func() unsafe.Pointer {
			tempVar := &priv_OP_names[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
		}())), get_ucpname(uint32(uint8((pcre_uchar(*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code)))))))), uint32(uint8((pcre_uchar(*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(2))*unsafe.Sizeof(*code))))))))), after)
	} else {
		var not *byte = func() *byte {
			if (map[bool]int32{false: 0, true: 1}[(int32(uint8((*code))) == OP_PROP)]) != 0 {
				return (&[]byte("\x00")[0])
			} else {
				return (&[]byte("not \x00")[0])
			}
		}()
		var p *pcre_uint32 = ((*pcre_uint32)(func() unsafe.Pointer {
			tempVar := &ucd_caseless_sets[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(2))*unsafe.Sizeof(*code))))))))*unsafe.Sizeof(*tempVar))
		}()))
		noarch.Fprintf(f, (&[]byte("%s%sclist\x00")[0]), before, not)
		for *p < pcre_uint32((uint32(4294967295))) {
			noarch.Fprintf(f, (&[]byte(" %04x\x00")[0]), pcre_uint32(*func() *pcre_uint32 {
				defer func() {
					p = ((*pcre_uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + (uintptr)(1)*unsafe.Sizeof(*p))))
				}()
				return p
			}()))
		}
		noarch.Fprintf(f, (&[]byte("%s\x00")[0]), after)
	}
}

func pcre_printint(external_re *pcre, f *noarch.File, print_lengths BOOL) {
	var re *real_pcre = (*real_pcre)(unsafe.Pointer(external_re))
	var codestart *pcre_uchar
	var code *pcre_uchar
	var utf BOOL
	var options uint32 = uint32((pcre_uint32((*re).options)))
	var offset int32 = int32(uint16((pcre_uint16((*re).name_table_offset))))
	var count int32 = int32(uint16((pcre_uint16((*re).name_count))))
	var size int32 = int32(uint16((pcre_uint16((*re).name_entry_size))))
	if uint32((pcre_uint32((*re).magic_number))) != uint32(1346589253) {
		offset = (((offset << uint64(int32(8))) & int32(65280)) | ((offset >> uint64(int32(8))) & int32(255)))
		count = (((count << uint64(int32(8))) & int32(65280)) | ((count >> uint64(int32(8))) & int32(255)))
		size = (((size << uint64(int32(8))) & int32(65280)) | ((size >> uint64(int32(8))) & int32(255)))
		options = (((((options << uint64(int32(24))) & uint32(4278190080)) | ((options << uint64(int32(8))) & uint32(int32(16711680)))) | ((options >> uint64(int32(8))) & uint32(int32(65280)))) | ((options >> uint64(int32(24))) & uint32(int32(255))))
	}
	codestart = ((*pcre_uchar)(func() unsafe.Pointer {
		tempVar := ((*pcre_uchar)(func() unsafe.Pointer {
			tempVar := (*pcre_uchar)(unsafe.Pointer(re))
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(offset)*unsafe.Sizeof(*tempVar))
		}()))
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)((count*size))*unsafe.Sizeof(*tempVar))
	}()))
	code = codestart
	utf = BOOL((map[bool]int32{false: 0, true: 1}[((options & uint32(int32(2048))) != uint32(int32(0)))]))
	for {
		var ccode *pcre_uchar
		var flag *byte = (&[]byte("  \x00")[0])
		var c pcre_uint32
		var extra uint32 = uint32(int32(0))
		var goto_CLASS_REF_REPEAT bool
		if int32((BOOL(print_lengths))) != 0 {
			noarch.Fprintf(f, (&[]byte("%3d \x00")[0]), (int32((int64(uintptr(unsafe.Pointer(code))) - int64(uintptr(unsafe.Pointer(codestart)))))))
		} else {
			noarch.Fprintf(f, (&[]byte("    \x00")[0]))
		}
		switch int32(uint8((pcre_uchar(*code)))) {
		case OP_TABLE_LENGTH:
			fallthrough
		case (OP_TABLE_LENGTH + map[bool]int32{false: 0, true: 1}[(((1296/8) == uint32(OP_TABLE_LENGTH)) && (162 == uint32(OP_TABLE_LENGTH)))]):
			{
			}
		case OP_END:
			{
				noarch.Fprintf(f, (&[]byte("    %s\n\x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())))
				noarch.Fprintf(f, (&[]byte("------------------------------------------------------------------\n\x00")[0]))
				return
			}
		case OP_CHAR:
			{
				noarch.Fprintf(f, (&[]byte("    \x00")[0]))
				for {
					code = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(1)*unsafe.Sizeof(*code))))
					code = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32((uint32(int32(1))+print_char(f, code, BOOL(utf)))))*unsafe.Sizeof(*code))))
					if noarch.NotInt32((map[bool]int32{false: 0, true: 1}[(int32(uint8((*code))) == OP_CHAR)])) != 0 {
						break
					}
				}
				noarch.Fprintf(f, (&[]byte("\n\x00")[0]))
				continue
			}
			fallthrough
		case OP_CHARI:
			{
				noarch.Fprintf(f, (&[]byte(" /i \x00")[0]))
				for {
					code = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(1)*unsafe.Sizeof(*code))))
					code = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32((uint32(int32(1))+print_char(f, code, BOOL(utf)))))*unsafe.Sizeof(*code))))
					if noarch.NotInt32((map[bool]int32{false: 0, true: 1}[(int32(uint8((*code))) == OP_CHARI)])) != 0 {
						break
					}
				}
				noarch.Fprintf(f, (&[]byte("\n\x00")[0]))
				continue
			}
			fallthrough
		case OP_CBRA:
			fallthrough
		case OP_CBRAPOS:
			fallthrough
		case OP_SCBRA:
			fallthrough
		case OP_SCBRAPOS:
			{
				if int32((BOOL(print_lengths))) != 0 {
					noarch.Fprintf(f, (&[]byte("%3d \x00")[0]), ((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code))))))))))
				} else {
					noarch.Fprintf(f, (&[]byte("    \x00")[0]))
				}
				noarch.Fprintf(f, (&[]byte("%s %d\x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())), uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(2)))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(((int32(1)+int32(2))+int32(1)))*unsafe.Sizeof(*(code)))))))))))
			}
		case OP_BRA:
			fallthrough
		case OP_BRAPOS:
			fallthrough
		case OP_SBRA:
			fallthrough
		case OP_SBRAPOS:
			fallthrough
		case OP_KETRMAX:
			fallthrough
		case OP_KETRMIN:
			fallthrough
		case OP_KETRPOS:
			fallthrough
		case OP_ALT:
			fallthrough
		case OP_KET:
			fallthrough
		case OP_ASSERT:
			fallthrough
		case OP_ASSERT_NOT:
			fallthrough
		case OP_ASSERTBACK:
			fallthrough
		case OP_ASSERTBACK_NOT:
			fallthrough
		case OP_ONCE:
			fallthrough
		case OP_ONCE_NC:
			fallthrough
		case OP_COND:
			fallthrough
		case OP_SCOND:
			fallthrough
		case OP_REVERSE:
			{
				if int32((BOOL(print_lengths))) != 0 {
					noarch.Fprintf(f, (&[]byte("%3d \x00")[0]), ((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code))))))))))
				} else {
					noarch.Fprintf(f, (&[]byte("    \x00")[0]))
				}
				noarch.Fprintf(f, (&[]byte("%s\x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())))
			}
		case OP_CLOSE:
			{
				noarch.Fprintf(f, (&[]byte("    %s %d\x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())), uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code)))))))))))
			}
		case OP_CREF:
			{
				noarch.Fprintf(f, (&[]byte("%3d %s\x00")[0]), uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code)))))))))), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())))
			}
		case OP_DNCREF:
			{
				var entry *pcre_uchar = ((*pcre_uchar)(func() unsafe.Pointer {
					tempVar := ((*pcre_uchar)(func() unsafe.Pointer {
						tempVar := ((*pcre_uchar)(func() unsafe.Pointer {
							tempVar := (*pcre_uchar)(unsafe.Pointer(re))
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(offset)*unsafe.Sizeof(*tempVar))
						}()))
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32((uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code))))))))<<uint64(int32(8)))|int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code))))))))))*uint32(size))))*unsafe.Sizeof(*tempVar))
					}()))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}()))
				noarch.Fprintf(f, (&[]byte(" %s Cond ref <\x00")[0]), flag)
				print_puchar(f, entry)
				noarch.Fprintf(f, (&[]byte(">%d\x00")[0]), uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(2)))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(((int32(1)+int32(2))+int32(1)))*unsafe.Sizeof(*(code)))))))))))
			}
		case OP_RREF:
			{
				c = pcre_uint32((uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code))))))))))))
				if c == pcre_uint32((uint32(int32(65535)))) {
					noarch.Fprintf(f, (&[]byte("    Cond recurse any\x00")[0]))
				} else {
					noarch.Fprintf(f, (&[]byte("    Cond recurse %d\x00")[0]), pcre_uint32(c))
				}
			}
		case OP_DNRREF:
			{
				var entry *pcre_uchar = ((*pcre_uchar)(func() unsafe.Pointer {
					tempVar := ((*pcre_uchar)(func() unsafe.Pointer {
						tempVar := ((*pcre_uchar)(func() unsafe.Pointer {
							tempVar := (*pcre_uchar)(unsafe.Pointer(re))
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(offset)*unsafe.Sizeof(*tempVar))
						}()))
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32((uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code))))))))<<uint64(int32(8)))|int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code))))))))))*uint32(size))))*unsafe.Sizeof(*tempVar))
					}()))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}()))
				noarch.Fprintf(f, (&[]byte(" %s Cond recurse <\x00")[0]), flag)
				print_puchar(f, entry)
				noarch.Fprintf(f, (&[]byte(">%d\x00")[0]), uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(2)))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(((int32(1)+int32(2))+int32(1)))*unsafe.Sizeof(*(code)))))))))))
			}
		case OP_DEF:
			{
				noarch.Fprintf(f, (&[]byte("    Cond def\x00")[0]))
			}
		case OP_STARI:
			fallthrough
		case OP_MINSTARI:
			fallthrough
		case OP_POSSTARI:
			fallthrough
		case OP_PLUSI:
			fallthrough
		case OP_MINPLUSI:
			fallthrough
		case OP_POSPLUSI:
			fallthrough
		case OP_QUERYI:
			fallthrough
		case OP_MINQUERYI:
			fallthrough
		case OP_POSQUERYI:
			{
				flag = (&[]byte("/i\x00")[0])
			}
			fallthrough
		case OP_STAR:
			fallthrough
		case OP_MINSTAR:
			fallthrough
		case OP_POSSTAR:
			fallthrough
		case OP_PLUS:
			fallthrough
		case OP_MINPLUS:
			fallthrough
		case OP_POSPLUS:
			fallthrough
		case OP_QUERY:
			fallthrough
		case OP_MINQUERY:
			fallthrough
		case OP_POSQUERY:
			fallthrough
		case OP_TYPESTAR:
			fallthrough
		case OP_TYPEMINSTAR:
			fallthrough
		case OP_TYPEPOSSTAR:
			fallthrough
		case OP_TYPEPLUS:
			fallthrough
		case OP_TYPEMINPLUS:
			fallthrough
		case OP_TYPEPOSPLUS:
			fallthrough
		case OP_TYPEQUERY:
			fallthrough
		case OP_TYPEMINQUERY:
			fallthrough
		case OP_TYPEPOSQUERY:
			{
				noarch.Fprintf(f, (&[]byte(" %s \x00")[0]), flag)
				if int32(uint8((*code))) >= OP_TYPESTAR {
					if (int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code))))))) == OP_PROP) || (int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code))))))) == OP_NOTPROP) {
						print_prop(f, ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code)))), (&[]byte("\x00")[0]), (&[]byte(" \x00")[0]))
						extra = uint32(int32(2))
					} else {
						noarch.Fprintf(f, (&[]byte("%s\x00")[0]), *((**byte)(func() unsafe.Pointer {
							tempVar := &priv_OP_names[0]
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code))))))))*unsafe.Sizeof(*tempVar))
						}())))
					}
				} else {
					extra = print_char(f, ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code)))), BOOL(utf))
				}
				noarch.Fprintf(f, (&[]byte("%s\x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())))
			}
		case OP_EXACTI:
			fallthrough
		case OP_UPTOI:
			fallthrough
		case OP_MINUPTOI:
			fallthrough
		case OP_POSUPTOI:
			{
				flag = (&[]byte("/i\x00")[0])
			}
			fallthrough
		case OP_EXACT:
			fallthrough
		case OP_UPTO:
			fallthrough
		case OP_MINUPTO:
			fallthrough
		case OP_POSUPTO:
			{
				noarch.Fprintf(f, (&[]byte(" %s \x00")[0]), flag)
				extra = print_char(f, ((*pcre_uchar)(func() unsafe.Pointer {
					tempVar := ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code))))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}())), BOOL(utf))
				noarch.Fprintf(f, (&[]byte("{\x00")[0]))
				if (int32(uint8((*code))) != OP_EXACT) && (int32(uint8((*code))) != OP_EXACTI) {
					noarch.Fprintf(f, (&[]byte("0,\x00")[0]))
				}
				noarch.Fprintf(f, (&[]byte("%d}\x00")[0]), uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code)))))))))))
				if (int32(uint8((*code))) == OP_MINUPTO) || (int32(uint8((*code))) == OP_MINUPTOI) {
					noarch.Fprintf(f, (&[]byte("?\x00")[0]))
				} else if (int32(uint8((*code))) == OP_POSUPTO) || (int32(uint8((*code))) == OP_POSUPTOI) {
					noarch.Fprintf(f, (&[]byte("+\x00")[0]))
				}
			}
		case OP_TYPEEXACT:
			fallthrough
		case OP_TYPEUPTO:
			fallthrough
		case OP_TYPEMINUPTO:
			fallthrough
		case OP_TYPEPOSUPTO:
			{
				if (int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)((int32(1)+int32(2)))*unsafe.Sizeof(*code))))))) == OP_PROP) || (int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)((int32(1)+int32(2)))*unsafe.Sizeof(*code))))))) == OP_NOTPROP) {
					print_prop(f, ((*pcre_uchar)(func() unsafe.Pointer {
						tempVar := ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(2))*unsafe.Sizeof(*code))))
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
					}())), (&[]byte("    \x00")[0]), (&[]byte(" \x00")[0]))
					extra = uint32(int32(2))
				} else {
					noarch.Fprintf(f, (&[]byte("    %s\x00")[0]), *((**byte)(func() unsafe.Pointer {
						tempVar := &priv_OP_names[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)((int32(1)+int32(2)))*unsafe.Sizeof(*code))))))))*unsafe.Sizeof(*tempVar))
					}())))
				}
				noarch.Fprintf(f, (&[]byte("{\x00")[0]))
				if int32(uint8((*code))) != OP_TYPEEXACT {
					noarch.Fprintf(f, (&[]byte("0,\x00")[0]))
				}
				noarch.Fprintf(f, (&[]byte("%d}\x00")[0]), uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code)))))))))))
				if int32(uint8((*code))) == OP_TYPEMINUPTO {
					noarch.Fprintf(f, (&[]byte("?\x00")[0]))
				} else if int32(uint8((*code))) == OP_TYPEPOSUPTO {
					noarch.Fprintf(f, (&[]byte("+\x00")[0]))
				}
			}
		case OP_NOTI:
			{
				flag = (&[]byte("/i\x00")[0])
			}
			fallthrough
		case OP_NOT:
			{
				noarch.Fprintf(f, (&[]byte(" %s [^\x00")[0]), flag)
				extra = print_char(f, ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code)))), BOOL(utf))
				noarch.Fprintf(f, (&[]byte("]\x00")[0]))
			}
		case OP_NOTSTARI:
			fallthrough
		case OP_NOTMINSTARI:
			fallthrough
		case OP_NOTPOSSTARI:
			fallthrough
		case OP_NOTPLUSI:
			fallthrough
		case OP_NOTMINPLUSI:
			fallthrough
		case OP_NOTPOSPLUSI:
			fallthrough
		case OP_NOTQUERYI:
			fallthrough
		case OP_NOTMINQUERYI:
			fallthrough
		case OP_NOTPOSQUERYI:
			{
				flag = (&[]byte("/i\x00")[0])
			}
			fallthrough
		case OP_NOTSTAR:
			fallthrough
		case OP_NOTMINSTAR:
			fallthrough
		case OP_NOTPOSSTAR:
			fallthrough
		case OP_NOTPLUS:
			fallthrough
		case OP_NOTMINPLUS:
			fallthrough
		case OP_NOTPOSPLUS:
			fallthrough
		case OP_NOTQUERY:
			fallthrough
		case OP_NOTMINQUERY:
			fallthrough
		case OP_NOTPOSQUERY:
			{
				noarch.Fprintf(f, (&[]byte(" %s [^\x00")[0]), flag)
				extra = print_char(f, ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code)))), BOOL(utf))
				noarch.Fprintf(f, (&[]byte("]%s\x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())))
			}
		case OP_NOTEXACTI:
			fallthrough
		case OP_NOTUPTOI:
			fallthrough
		case OP_NOTMINUPTOI:
			fallthrough
		case OP_NOTPOSUPTOI:
			{
				flag = (&[]byte("/i\x00")[0])
			}
			fallthrough
		case OP_NOTEXACT:
			fallthrough
		case OP_NOTUPTO:
			fallthrough
		case OP_NOTMINUPTO:
			fallthrough
		case OP_NOTPOSUPTO:
			{
				noarch.Fprintf(f, (&[]byte(" %s [^\x00")[0]), flag)
				extra = print_char(f, ((*pcre_uchar)(func() unsafe.Pointer {
					tempVar := ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code))))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}())), BOOL(utf))
				noarch.Fprintf(f, (&[]byte("]{\x00")[0]))
				if (int32(uint8((*code))) != OP_NOTEXACT) && (int32(uint8((*code))) != OP_NOTEXACTI) {
					noarch.Fprintf(f, (&[]byte("0,\x00")[0]))
				}
				noarch.Fprintf(f, (&[]byte("%d}\x00")[0]), uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code)))))))))))
				if (int32(uint8((*code))) == OP_NOTMINUPTO) || (int32(uint8((*code))) == OP_NOTMINUPTOI) {
					noarch.Fprintf(f, (&[]byte("?\x00")[0]))
				} else if (int32(uint8((*code))) == OP_NOTPOSUPTO) || (int32(uint8((*code))) == OP_NOTPOSUPTOI) {
					noarch.Fprintf(f, (&[]byte("+\x00")[0]))
				}
			}
		case OP_RECURSE:
			{
				if int32((BOOL(print_lengths))) != 0 {
					noarch.Fprintf(f, (&[]byte("%3d \x00")[0]), ((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code))))))))))
				} else {
					noarch.Fprintf(f, (&[]byte("    \x00")[0]))
				}
				noarch.Fprintf(f, (&[]byte("%s\x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())))
			}
		case OP_REFI:
			{
				flag = (&[]byte("/i\x00")[0])
			}
			fallthrough
		case OP_REF:
			{
				noarch.Fprintf(f, (&[]byte(" %s \\%d\x00")[0]), flag, uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code)))))))))))
				ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
					tempVar := &priv_OP_lengths[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}()))))))*unsafe.Sizeof(*code))))
				goto_CLASS_REF_REPEAT = true
				goto CLASS_REF_REPEAT_CONTAINER
			}
			fallthrough
		case OP_DNREFI:
			{
				flag = (&[]byte("/i\x00")[0])
			}
			fallthrough
		case OP_DNREF:
			{
				var entry *pcre_uchar = ((*pcre_uchar)(func() unsafe.Pointer {
					tempVar := ((*pcre_uchar)(func() unsafe.Pointer {
						tempVar := ((*pcre_uchar)(func() unsafe.Pointer {
							tempVar := (*pcre_uchar)(unsafe.Pointer(re))
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(offset)*unsafe.Sizeof(*tempVar))
						}()))
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32((uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code))))))))<<uint64(int32(8)))|int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code))))))))))*uint32(size))))*unsafe.Sizeof(*tempVar))
					}()))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}()))
				noarch.Fprintf(f, (&[]byte(" %s \\k<\x00")[0]), flag)
				print_puchar(f, entry)
				noarch.Fprintf(f, (&[]byte(">%d\x00")[0]), uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(2)))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(((int32(1)+int32(2))+int32(1)))*unsafe.Sizeof(*(code)))))))))))
				ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
					tempVar := &priv_OP_lengths[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}()))))))*unsafe.Sizeof(*code))))
				goto_CLASS_REF_REPEAT = true
				goto CLASS_REF_REPEAT_CONTAINER
			}
			fallthrough
		case OP_CALLOUT:
			{
				noarch.Fprintf(f, (&[]byte("    %s %d %d %d\x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())), int32(uint8((pcre_uchar(*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code)))))))), ((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(2))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(2)+int32(1)))*unsafe.Sizeof(*(code))))))))), ((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(2)+int32(2)))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(((int32(2)+int32(2))+int32(1)))*unsafe.Sizeof(*(code))))))))))
			}
		case OP_PROP:
			fallthrough
		case OP_NOTPROP:
			{
				print_prop(f, code, (&[]byte("    \x00")[0]), (&[]byte("\x00")[0]))
			}
		case OP_CLASS:
			fallthrough
		case OP_NCLASS:
			fallthrough
		case OP_XCLASS:
			goto SW_GENERATED_LABEL_41
		case OP_MARK:
			fallthrough
		case OP_PRUNE_ARG:
			fallthrough
		case OP_SKIP_ARG:
			fallthrough
		case OP_THEN_ARG:
			{
				noarch.Fprintf(f, (&[]byte("    %s \x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())))
				print_puchar(f, ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(2))*unsafe.Sizeof(*code)))))
				extra += uint32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code)))))))
			}
		case OP_THEN:
			{
				noarch.Fprintf(f, (&[]byte("    %s\x00")[0]), *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())))
			}
		case OP_CIRCM:
			fallthrough
		case OP_DOLLM:
			{
				flag = (&[]byte("/m\x00")[0])
			}
			fallthrough
		default:
			{
				noarch.Fprintf(f, (&[]byte(" %s %s\x00")[0]), flag, *((**byte)(func() unsafe.Pointer {
					tempVar := &priv_OP_names[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
				}())))
				break
			}
		}
		goto SW_GENERATED_LABEL_42
	CLASS_REF_REPEAT_CONTAINER:
		;
	SW_GENERATED_LABEL_41:
		{
			var i int32
			var min uint32
			var max uint32
			var printmap BOOL
			var invertmap BOOL = BOOL((int32(0)))
			var map_ *pcre_uint8
			var inverted_map []pcre_uint8 = make([]pcre_uint8, 32, 32)
			if goto_CLASS_REF_REPEAT {
				goto CLASS_REF_REPEAT
			}
			noarch.Fprintf(f, (&[]byte("    [\x00")[0]))
			if int32(uint8((*code))) == OP_XCLASS {
				extra = uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)(int32(1))*unsafe.Sizeof(*(code)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((code))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(code))))))))))
				ccode = ((*pcre_uchar)(func() unsafe.Pointer {
					tempVar := ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(2))*unsafe.Sizeof(*code))))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}()))
				printmap = BOOL((map[bool]int32{false: 0, true: 1}[((int32(uint8((*ccode))) & int32(2)) != int32(0))]))
				if (int32(uint8((*ccode))) & int32(1)) != int32(0) {
					invertmap = BOOL((map[bool]int32{false: 0, true: 1}[((int32(uint8((*ccode))) & int32(4)) == int32(0))]))
					noarch.Fprintf(f, (&[]byte("^\x00")[0]))
				}
				ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ccode)) + (uintptr)(1)*unsafe.Sizeof(*ccode))))
			} else {
				printmap = BOOL((int32(1)))
				ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32(1))*unsafe.Sizeof(*code))))
			}
			if int32((BOOL(printmap))) != 0 {
				map_ = (*pcre_uint8)(unsafe.Pointer(ccode))
				if int32((BOOL(invertmap))) != 0 {
					for i = int32(0); i < int32(32); i++ {
						*((*pcre_uint8)(func() unsafe.Pointer {
							tempVar := &inverted_map[0]
							return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
						}())) = pcre_uint8((^uint8(int32(uint8((pcre_uint8(*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(map_)) + (uintptr)(i)*unsafe.Sizeof(*map_)))))))))))
					}
					map_ = &inverted_map[0]
				}
				for i = int32(0); i < int32(256); i++ {
					if (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(map_)) + (uintptr)((i/int32(8)))*unsafe.Sizeof(*map_))))))) & (int32(1) << uint64((i & int32(7))))) != int32(0) {
						var j int32
						for j = (i + int32(1)); j < int32(256); j++ {
							if (int32(uint8((*((*pcre_uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(map_)) + (uintptr)((j/int32(8)))*unsafe.Sizeof(*map_))))))) & (int32(1) << uint64((j & int32(7))))) == int32(0) {
								break
							}
						}
						if (i == int32('-')) || (i == int32(']')) {
							noarch.Fprintf(f, (&[]byte("\\\x00")[0]))
						}
						if (map[bool]int32{false: 0, true: 1}[((i >= int32(32)) && (i < int32(127)))]) != 0 {
							noarch.Fprintf(f, (&[]byte("%c\x00")[0]), i)
						} else {
							noarch.Fprintf(f, (&[]byte("\\x%02x\x00")[0]), i)
						}
						if func() int32 {
							j -= 1
							return j
						}() > i {
							if j != (i + int32(1)) {
								noarch.Fprintf(f, (&[]byte("-\x00")[0]))
							}
							if (j == int32('-')) || (j == int32(']')) {
								noarch.Fprintf(f, (&[]byte("\\\x00")[0]))
							}
							if (map[bool]int32{false: 0, true: 1}[((j >= int32(32)) && (j < int32(127)))]) != 0 {
								noarch.Fprintf(f, (&[]byte("%c\x00")[0]), j)
							} else {
								noarch.Fprintf(f, (&[]byte("\\x%02x\x00")[0]), j)
							}
						}
						i = j
					}
				}
				ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ccode)) + (uintptr)(int32((uint32(int32(32))/1)))*unsafe.Sizeof(*ccode))))
			}
			if int32(uint8((*code))) == OP_XCLASS {
				var ch pcre_uchar
				for int32(uint8((func() pcre_uchar {
					tempVar := *func() *pcre_uchar {
						defer func() {
							ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ccode)) + (uintptr)(1)*unsafe.Sizeof(*ccode))))
						}()
						return ccode
					}()
					ch = tempVar
					return tempVar
				}()))) != int32(0) {
					var not BOOL = BOOL((int32(0)))
					var notch *byte = (&[]byte("\x00")[0])
					switch int32(uint8((pcre_uchar(ch)))) {
					case int32(4):
						{
							not = BOOL((int32(1)))
							notch = (&[]byte("^\x00")[0])
						}
						fallthrough
					case int32(3):
						{
							var ptype uint32 = uint32(uint8((*func() *pcre_uchar {
								defer func() {
									ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ccode)) + (uintptr)(1)*unsafe.Sizeof(*ccode))))
								}()
								return ccode
							}())))
							var pvalue uint32 = uint32(uint8((*func() *pcre_uchar {
								defer func() {
									ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ccode)) + (uintptr)(1)*unsafe.Sizeof(*ccode))))
								}()
								return ccode
							}())))
							switch ptype {
							case uint32(int32(11)):
								{
									noarch.Fprintf(f, (&[]byte("[:%sgraph:]\x00")[0]), notch)
								}
							case uint32(int32(12)):
								{
									noarch.Fprintf(f, (&[]byte("[:%sprint:]\x00")[0]), notch)
								}
							case uint32(int32(13)):
								{
									noarch.Fprintf(f, (&[]byte("[:%spunct:]\x00")[0]), notch)
								}
							default:
								{
									noarch.Fprintf(f, (&[]byte("\\%c{%s}\x00")[0]), (func() int32 {
										if int32((BOOL(not))) != 0 {
											return int32('P')
										} else {
											return int32('p')
										}
									}()), get_ucpname(ptype, pvalue))
									break
								}
							}
						}
					default:
						{
							ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ccode)) + (uintptr)(int32((uint32(int32(1))+print_char(f, ccode, BOOL(utf)))))*unsafe.Sizeof(*ccode))))
							if int32(uint8((ch))) == int32(2) {
								noarch.Fprintf(f, (&[]byte("-\x00")[0]))
								ccode = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ccode)) + (uintptr)(int32((uint32(int32(1))+print_char(f, ccode, BOOL(utf)))))*unsafe.Sizeof(*ccode))))
							}
							break
						}
					}
				}
			}
			noarch.Fprintf(f, (&[]byte("]%s\x00")[0]), func() *byte {
				if (map[bool]int32{false: 0, true: 1}[(int32(uint8((*code))) == OP_NCLASS)]) != 0 {
					return (&[]byte(" (neg)\x00")[0])
				} else {
					return (&[]byte("\x00")[0])
				}
			}())
		CLASS_REF_REPEAT:
			goto_CLASS_REF_REPEAT = false
			switch int32(uint8((pcre_uchar(*ccode)))) {
			case OP_CRSTAR:
				fallthrough
			case OP_CRMINSTAR:
				fallthrough
			case OP_CRPLUS:
				fallthrough
			case OP_CRMINPLUS:
				fallthrough
			case OP_CRQUERY:
				fallthrough
			case OP_CRMINQUERY:
				fallthrough
			case OP_CRPOSSTAR:
				fallthrough
			case OP_CRPOSPLUS:
				fallthrough
			case OP_CRPOSQUERY:
				{
					noarch.Fprintf(f, (&[]byte("%s\x00")[0]), *((**byte)(func() unsafe.Pointer {
						tempVar := &priv_OP_names[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*ccode))))*unsafe.Sizeof(*tempVar))
					}())))
					extra += uint32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &priv_OP_lengths[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*ccode))))*unsafe.Sizeof(*tempVar))
					}())))))
				}
			case OP_CRRANGE:
				fallthrough
			case OP_CRMINRANGE:
				fallthrough
			case OP_CRPOSRANGE:
				{
					min = uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((ccode))) + (uintptr)(int32(1))*unsafe.Sizeof(*(ccode)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((ccode))) + (uintptr)((int32(1)+int32(1)))*unsafe.Sizeof(*(ccode))))))))))
					max = uint32(((int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((ccode))) + (uintptr)((int32(1)+int32(2)))*unsafe.Sizeof(*(ccode)))))))) << uint64(int32(8))) | int32(uint8((*((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer((ccode))) + (uintptr)(((int32(1)+int32(2))+int32(1)))*unsafe.Sizeof(*(ccode))))))))))
					if max == uint32(int32(0)) {
						noarch.Fprintf(f, (&[]byte("{%u,}\x00")[0]), min)
					} else {
						noarch.Fprintf(f, (&[]byte("{%u,%u}\x00")[0]), min, max)
					}
					if int32(uint8((*ccode))) == OP_CRMINRANGE {
						noarch.Fprintf(f, (&[]byte("?\x00")[0]))
					} else if int32(uint8((*ccode))) == OP_CRPOSRANGE {
						noarch.Fprintf(f, (&[]byte("+\x00")[0]))
					}
					extra += uint32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
						tempVar := &priv_OP_lengths[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*ccode))))*unsafe.Sizeof(*tempVar))
					}())))))
				}
			default:
				{
					break
				}
			}
		}
	SW_GENERATED_LABEL_42:
		;
		code = ((*pcre_uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(code)) + (uintptr)(int32((uint32(uint8((*((*pcre_uint8)(func() unsafe.Pointer {
			tempVar := &priv_OP_lengths[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint8((*code))))*unsafe.Sizeof(*tempVar))
		}())))))+extra)))*unsafe.Sizeof(*code))))
		noarch.Fprintf(f, (&[]byte("\n\x00")[0]))
	}
}
