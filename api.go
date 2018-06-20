// Package pcre provides access to the Perl Compatible Regular
// Expresion library, PCRE.
//
// It implements two main types, Regexp and Matcher.  Regexp objects
// store a compiled regular expression. They consist of two immutable
// parts: pcre and pcre_extra. Compile()/MustCompile() initialize pcre.
// Calling Study() on a compiled Regexp initializes pcre_extra.
// Compilation of regular expressions using Compile or MustCompile is
// slightly expensive, so these objects should be kept and reused,
// instead of compiling them from scratch for each matching attempt.
// CompileJIT and MustCompileJIT are way more expensive, because they
// run Study() after compiling a Regexp, but they tend to give
// much better perfomance:
// http://sljit.sourceforge.net/regex_perf.html
//
// Matcher objects keeps the results of a match against a []byte or
// string subject.  The Group and GroupString functions provide access
// to capture groups; both versions work no matter if the subject was a
// []byte or string, but the version with the matching type is slightly
// more efficient.
//
// Matcher objects contain some temporary space and refer the original
// subject.  They are mutable and can be reused (using Match,
// MatchString, Reset or ResetString).
//
// For details on the regular expression language implemented by this
// package and the flags defined below, see the PCRE documentation.
// http://www.pcre.org/pcre.txt
package libpcre

import (
	"fmt"
	"github.com/elliotchance/c2go/noarch"
	"strconv"
	"unsafe"
)

// Flags for Compile and Match functions.
const (
	ANCHORED          = int(pcre_ANCHORED)
	BSR_ANYCRLF       = int(pcre_BSR_ANYCRLF)
	BSR_UNICODE       = int(pcre_BSR_UNICODE)
	NEWLINE_ANY       = int(pcre_NEWLINE_ANY)
	NEWLINE_ANYCRLF   = int(pcre_NEWLINE_ANYCRLF)
	NEWLINE_CR        = int(pcre_NEWLINE_CR)
	NEWLINE_CRLF      = int(pcre_NEWLINE_CRLF)
	NEWLINE_LF        = int(pcre_NEWLINE_LF)
	NO_START_OPTIMIZE = int(pcre_NO_START_OPTIMIZE)
	NO_UTF8_CHECK     = int(pcre_NO_UTF8_CHECK)
)

// Flags for Compile functions
const (
	CASELESS          = int(pcre_CASELESS)
	DOLLAR_ENDONLY    = int(pcre_DOLLAR_ENDONLY)
	DOTALL            = int(pcre_DOTALL)
	DUPNAMES          = int(pcre_DUPNAMES)
	EXTENDED          = int(pcre_EXTENDED)
	EXTRA             = int(pcre_EXTRA)
	FIRSTLINE         = int(pcre_FIRSTLINE)
	JAVASCRIPT_COMPAT = int(pcre_JAVASCRIPT_COMPAT)
	MULTILINE         = int(pcre_MULTILINE)
	NEVER_UTF         = int(pcre_NEVER_UTF)
	NO_AUTO_CAPTURE   = int(pcre_NO_AUTO_CAPTURE)
	UNGREEDY          = int(pcre_UNGREEDY)
	UTF8              = int(pcre_UTF8)
	UCP               = int(pcre_UCP)
)

// Flags for Match functions
const (
	NOTBOL           = int(pcre_NOTBOL)
	NOTEOL           = int(pcre_NOTEOL)
	NOTEMPTY         = int(pcre_NOTEMPTY)
	NOTEMPTY_ATSTART = int(pcre_NOTEMPTY_ATSTART)
	PARTIAL_HARD     = int(pcre_PARTIAL_HARD)
	PARTIAL_SOFT     = int(pcre_PARTIAL_SOFT)
)

// Flags for Study function
const (
	STUDY_JIT_COMPILE              = int(pcre_STUDY_JIT_COMPILE)
	STUDY_JIT_PARTIAL_SOFT_COMPILE = int(pcre_STUDY_JIT_PARTIAL_SOFT_COMPILE)
	STUDY_JIT_PARTIAL_HARD_COMPILE = int(pcre_STUDY_JIT_PARTIAL_HARD_COMPILE)
)

// Exec-time and get/set-time error codes
const (
	ERROR_NOMATCH        = int(pcre_ERROR_NOMATCH)
	ERROR_NULL           = int(pcre_ERROR_NULL)
	ERROR_BADOPTION      = int(pcre_ERROR_BADOPTION)
	ERROR_BADMAGIC       = int(pcre_ERROR_BADMAGIC)
	ERROR_UNKNOWN_OPCODE = int(pcre_ERROR_UNKNOWN_OPCODE)
	ERROR_UNKNOWN_NODE   = int(pcre_ERROR_UNKNOWN_NODE)
	ERROR_NOMEMORY       = int(pcre_ERROR_NOMEMORY)
	ERROR_NOSUBSTRING    = int(pcre_ERROR_NOSUBSTRING)
	ERROR_MATCHLIMIT     = int(pcre_ERROR_MATCHLIMIT)
	ERROR_CALLOUT        = int(pcre_ERROR_CALLOUT)
	ERROR_BADUTF8        = int(pcre_ERROR_BADUTF8)
	ERROR_BADUTF8_OFFSET = int(pcre_ERROR_BADUTF8_OFFSET)
	ERROR_PARTIAL        = int(pcre_ERROR_PARTIAL)
	ERROR_BADPARTIAL     = int(pcre_ERROR_BADPARTIAL)
	ERROR_RECURSIONLIMIT = int(pcre_ERROR_RECURSIONLIMIT)
	ERROR_INTERNAL       = int(pcre_ERROR_INTERNAL)
	ERROR_BADCOUNT       = int(pcre_ERROR_BADCOUNT)
	ERROR_JIT_STACKLIMIT = int(pcre_ERROR_JIT_STACKLIMIT)
)

// Regexp holds a reference to a compiled regular expression.
// Use Compile or MustCompile to create such objects.
type Regexp struct {
	ptr   []byte
	extra []byte
}

// Number of bytes in the compiled pattern
func pcreSize(ptr *pcre) (size size_t) {
	pcre_fullinfo(ptr, nil, pcre_INFO_SIZE, unsafe.Pointer(&size))
	return
}

// Number of capture groups
func pcreGroups(ptr *pcre) (count int32) {
	pcre_fullinfo((*pcre)(unsafe.Pointer(ptr)), nil,
		pcre_INFO_CAPTURECOUNT, unsafe.Pointer(&count))
	return
}

// Move pattern to the Go heap so that we do not have to use a
// finalizer.  PCRE patterns are fully relocatable. (We do not use
// custom character tables.)
func toHeap(ptr *pcre) (re Regexp) {
	size := pcreSize(ptr)
	re.ptr = make([]byte, int(size))
	if size > 0 {
		noarch.Memcpy(unsafe.Pointer(&re.ptr[0]), unsafe.Pointer(ptr), int32(size))
	}
	return
}

// Compile the pattern and return a compiled regexp.
// If compilation fails, the second return value holds a *CompileError.
func Compile(pattern string, flags int) (Regexp, error) {
	pattern1 := noarch.StringToCString(pattern)
	if clen := int(noarch.Strlen(pattern1)); clen != len(pattern) {
		return Regexp{}, &CompileError{
			Pattern: pattern,
			Message: "NUL byte in pattern",
			Offset:  clen,
		}
	}
	var errptr *byte
	var erroffset int32
	ptr := pcre_compile(pattern1, int32(flags), &errptr, &erroffset, nil)
	if ptr == nil {
		return Regexp{}, &CompileError{
			Pattern: pattern,
			Message: noarch.CStringToString(errptr),
			Offset:  int(erroffset),
		}
	}
	heap := toHeap(ptr)
	return heap, nil
}

// CompileJIT is a combination of Compile and Study. It first compiles
// the pattern and if this succeeds calls Study on the compiled pattern.
// comFlags are Compile flags, jitFlags are study flags.
// If compilation fails, the second return value holds a *CompileError.
func CompileJIT(pattern string, comFlags, jitFlags int) (Regexp, error) {
	re, err := Compile(pattern, comFlags)
	if err == nil {
		err = (&re).Study(jitFlags)
	}
	return re, err
}

// MustCompile compiles the pattern.  If compilation fails, panic.
func MustCompile(pattern string, flags int) (re Regexp) {
	re, err := Compile(pattern, flags)
	if err != nil {
		panic(err)
	}
	return
}

// MustCompileJIT compiles and studies the pattern.  On failure it panics.
func MustCompileJIT(pattern string, comFlags, jitFlags int) (re Regexp) {
	re, err := CompileJIT(pattern, comFlags, jitFlags)
	if err != nil {
		panic(err)
	}
	return
}

// Study adds Just-In-Time compilation to a Regexp. This may give a huge
// speed boost when matching. If an error occurs, return value is non-nil.
// Flags optionally specifies JIT compilation options for partial matches.
func (re *Regexp) Study(flags int) error {
	if re.extra != nil {
		return fmt.Errorf("Study: Regexp has already been optimized")
	}
	if flags == 0 {
		flags = STUDY_JIT_COMPILE
	}

	ptr := (*pcre)(unsafe.Pointer(&re.ptr[0]))
	var err *byte
	extra := pcre_study(ptr, int32(flags), &err)
	if err != nil {
		return fmt.Errorf("%s", noarch.CStringToString(err))
	}
	if extra == nil {
		// Studying the pattern may not produce useful information.
		return nil
	}
	//defer free(unsafe.Pointer(extra))

	var size size_t
	rc := pcre_fullinfo(ptr, extra, pcre_INFO_JITSIZE, unsafe.Pointer(&size))
	if rc != 0 || size == 0 {
		return fmt.Errorf("Study failed to obtain JIT size (%d)", int(rc))
	}
	re.extra = make([]byte, int(size))
	noarch.Memcpy(unsafe.Pointer(&re.extra[0]), unsafe.Pointer(extra), int32(size))
	return nil
}

// Groups returns the number of capture groups in the compiled pattern.
func (re Regexp) Groups() int {
	if re.ptr == nil {
		panic("Regexp.Groups: uninitialized")
	}
	return int(pcreGroups((*pcre)(unsafe.Pointer(&re.ptr[0]))))
}

// Matcher objects provide a place for storing match results.
// They can be created by the Matcher and MatcherString functions,
// or they can be initialized with Reset or ResetString.
type Matcher struct {
	re       Regexp
	groups   int
	ovector  []int32 // scratch space for capture offsets
	matches  bool    // last match was successful
	partial  bool    // was the last match a partial match?
	subjects string  // one of these fields is set to record the subject,
	subjectb []byte  // so that Group/GroupString can return slices
}

// NewMatcher creates a new matcher object for the given Regexp.
func (re Regexp) NewMatcher() (m *Matcher) {
	m = new(Matcher)
	m.Init(&re)
	return
}

// Matcher creates a new matcher object, with the byte slice as subject.
// It also starts a first match on subject. Test for success with Matches().
func (re Regexp) Matcher(subject []byte, flags int) (m *Matcher) {
	m = re.NewMatcher()
	m.Match(subject, flags)
	return
}

// MatcherString creates a new matcher, with the specified subject string.
// It also starts a first match on subject. Test for success with Matches().
func (re Regexp) MatcherString(subject string, flags int) (m *Matcher) {
	m = re.NewMatcher()
	m.MatchString(subject, flags)
	return
}

// Reset switches the matcher object to the specified regexp and subject.
// It also starts a first match on subject.
func (m *Matcher) Reset(re Regexp, subject []byte, flags int) bool {
	m.Init(&re)
	return m.Match(subject, flags)
}

// ResetString switches the matcher object to the given regexp and subject.
// It also starts a first match on subject.
func (m *Matcher) ResetString(re Regexp, subject string, flags int) bool {
	m.Init(&re)
	return m.MatchString(subject, flags)
}

// Init binds an existing Matcher object to the given Regexp.
func (m *Matcher) Init(re *Regexp) {
	if re.ptr == nil {
		panic("Matcher.Init: uninitialized")
	}
	m.matches = false
	if m.re.ptr != nil && &m.re.ptr[0] == &re.ptr[0] {
		// Skip group count extraction if the matcher has
		// already been initialized with the same regular
		// expression.
		return
	}
	m.re = *re
	m.groups = re.Groups()
	if ovectorlen := 3 * (1 + m.groups); len(m.ovector) < ovectorlen {
		m.ovector = make([]int32, ovectorlen)
	}
}

var nullbyte = []byte{0}

// Match tries to match the specified byte slice to
// the current pattern by calling Exec and collects the result.
// Returns true if the match succeeds.
func (m *Matcher) Match(subject []byte, flags int) bool {
	if m.re.ptr == nil {
		panic("Matcher.Match: uninitialized")
	}
	rc := m.Exec(subject, flags)
	m.matches = matched(int32(rc))
	m.partial = (rc == ERROR_PARTIAL)
	return m.matches
}

// MatchString tries to match the specified subject string to
// the current pattern by calling ExecString and collects the result.
// Returns true if the match succeeds.
func (m *Matcher) MatchString(subject string, flags int) bool {
	if m.re.ptr == nil {
		panic("Matcher.MatchString: uninitialized")
	}
	rc := m.ExecString(subject, flags)
	m.matches = matched(int32(rc))
	m.partial = (rc == ERROR_PARTIAL)
	return m.matches
}

// Exec tries to match the specified byte slice to
// the current pattern. Returns the raw pcre_exec error code.
func (m *Matcher) Exec(subject []byte, flags int) int {
	if m.re.ptr == nil {
		panic("Matcher.Exec: uninitialized")
	}
	length := len(subject)
	m.subjects = ""
	m.subjectb = subject
	if length == 0 {
		subject = nullbyte // make first character adressable
	}
	//subjectptr := (*char)(unsafe.Pointer(&subject[0]))
	return m.exec(&subject[0], length, flags)
}

// ExecString tries to match the specified subject string to
// the current pattern. It returns the raw pcre_exec error code.
func (m *Matcher) ExecString(subject string, flags int) int {
	if m.re.ptr == nil {
		panic("Matcher.ExecString: uninitialized")
	}
	length := len(subject)
	m.subjects = subject
	m.subjectb = nil
	if length == 0 {
		subject = "\000" // make first character addressable
	}
	// The following is a non-portable kludge to avoid a copy
	//subjectptr := *(**char)(unsafe.Pointer(&subject))
	return m.exec(&append([]byte(subject), 0)[0], length, flags)
}

func (m *Matcher) exec(subjectptr *byte, length, flags int) int {
	var extra *pcre_extra
	if m.re.extra != nil {
		extra = (*pcre_extra)(unsafe.Pointer(&m.re.extra[0]))
	}
	rc := pcre_exec((*pcre)(unsafe.Pointer(&m.re.ptr[0])), (*pcre_extra)(unsafe.Pointer(extra)),
		subjectptr, int32(length),
		0, int32(flags), (*int32)(unsafe.Pointer(&m.ovector[0])), int32(len(m.ovector)))
	return int(rc)
}

// matched checks the return code of a pattern match for success.
func matched(rc int32) bool {
	switch {
	case rc >= 0 || rc == pcre_ERROR_PARTIAL:
		return true
	case rc == pcre_ERROR_NOMATCH:
		return false
	case rc == pcre_ERROR_BADOPTION:
		panic("PCRE.Match: invalid option flag")
	}
	panic("unexpected return code from pcre_exec: " + strconv.Itoa(int(rc)))
}

// Matches returns true if a previous call to Matcher, MatcherString, Reset,
// ResetString, Match or MatchString succeeded.
func (m *Matcher) Matches() bool {
	return m.matches
}

// Partial returns true if a previous call to Matcher, MatcherString, Reset,
// ResetString, Match or MatchString found a partial match.
func (m *Matcher) Partial() bool {
	return m.partial
}

// Groups returns the number of groups in the current pattern.
func (m *Matcher) Groups() int {
	return m.groups
}

// Present returns true if the numbered capture group is present in the last
// match (performed by Matcher, MatcherString, Reset, ResetString,
// Match, or MatchString).  Group numbers start at 1.  A capture group
// can be present and match the empty string.
func (m *Matcher) Present(group int) bool {
	return m.ovector[2*group] >= 0
}

// Group returns the numbered capture group of the last match (performed by
// Matcher, MatcherString, Reset, ResetString, Match, or MatchString).
// Group 0 is the part of the subject which matches the whole pattern;
// the first actual capture group is numbered 1.  Capture groups which
// are not present return a nil slice.
func (m *Matcher) Group(group int) []byte {
	start := m.ovector[2*group]
	end := m.ovector[2*group+1]
	if start >= 0 {
		if m.subjectb != nil {
			return m.subjectb[start:end]
		}
		return []byte(m.subjects[start:end])
	}
	return nil
}

// Extract returns a slice of byte slices for a single match.
// The first byte slice contains the complete match.
// Subsequent byte slices contain the captured groups.
// If there was no match then nil is returned.
func (m *Matcher) Extract() [][]byte {
	if !m.matches {
		return nil
	}
	extract := make([][]byte, m.groups+1)
	extract[0] = m.subjectb
	for i := 1; i <= m.groups; i++ {
		x0 := m.ovector[2*i]
		x1 := m.ovector[2*i+1]
		extract[i] = m.subjectb[x0:x1]
	}
	return extract
}

// ExtractString returns a slice of strings for a single match.
// The first string contains the complete match.
// Subsequent strings in the slice contain the captured groups.
// If there was no match then nil is returned.
func (m *Matcher) ExtractString() []string {
	if !m.matches {
		return nil
	}
	extract := make([]string, m.groups+1)
	extract[0] = m.subjects
	for i := 1; i <= m.groups; i++ {
		x0 := m.ovector[2*i]
		x1 := m.ovector[2*i+1]
		extract[i] = m.subjects[x0:x1]
	}
	return extract
}

// GroupIndices returns the numbered capture group positions of the last
// match (performed by Matcher, MatcherString, Reset, ResetString, Match,
// or MatchString). Group 0 is the part of the subject which matches
// the whole pattern; the first actual capture group is numbered 1.
// Capture groups which are not present return a nil slice.
func (m *Matcher) GroupIndices(group int) []int {
	start := m.ovector[2*group]
	end := m.ovector[2*group+1]
	if start >= 0 {
		return []int{int(start), int(end)}
	}
	return nil
}

// GroupString returns the numbered capture group as a string.  Group 0
// is the part of the subject which matches the whole pattern; the first
// actual capture group is numbered 1.  Capture groups which are not
// present return an empty string.
func (m *Matcher) GroupString(group int) string {
	start := m.ovector[2*group]
	end := m.ovector[2*group+1]
	if start >= 0 {
		if m.subjectb != nil {
			return string(m.subjectb[start:end])
		}
		return m.subjects[start:end]
	}
	return ""
}

// Index returns the start and end of the first match, if a previous
// call to Matcher, MatcherString, Reset, ResetString, Match or
// MatchString succeeded. loc[0] is the start and loc[1] is the end.
func (m *Matcher) Index() (loc []int) {
	if !m.matches {
		return nil
	}
	loc = []int{int(m.ovector[0]), int(m.ovector[1])}
	return
}

// name2index converts a group name to its group index number.
func (m *Matcher) name2index(name string) (int, error) {
	if m.re.ptr == nil {
		return 0, fmt.Errorf("Matcher.Named: uninitialized")
	}
	name1 := noarch.StringToCString(name)
	group := int(pcre_get_stringnumber(
		(*pcre)(unsafe.Pointer(&m.re.ptr[0])), name1))
	if group < 0 {
		return group, fmt.Errorf("Matcher.Named: unknown name: " + name)
	}
	return group, nil
}

// Named returns the value of the named capture group.
// This is a nil slice if the capture group is not present.
// If the name does not refer to a group then error is non-nil.
func (m *Matcher) Named(group string) ([]byte, error) {
	groupNum, err := m.name2index(group)
	if err != nil {
		return []byte{}, err
	}
	return m.Group(groupNum), nil
}

// NamedString returns the value of the named capture group,
// or an empty string if the capture group is not present.
// If the name does not refer to a group then error is non-nil.
func (m *Matcher) NamedString(group string) (string, error) {
	groupNum, err := m.name2index(group)
	if err != nil {
		return "", err
	}
	return m.GroupString(groupNum), nil
}

// NamedPresent returns true if the named capture group is present.
// If the name does not refer to a group then error is non-nil.
func (m *Matcher) NamedPresent(group string) (bool, error) {
	groupNum, err := m.name2index(group)
	if err != nil {
		return false, err
	}
	return m.Present(groupNum), nil
}

// FindIndex returns the start and end of the first match,
// or nil if no match.  loc[0] is the start and loc[1] is the end.
func (re *Regexp) FindIndex(bytes []byte, flags int) (loc []int) {
	m := re.Matcher(bytes, flags)
	if m.Matches() {
		loc = []int{int(m.ovector[0]), int(m.ovector[1])}
		return
	}
	return nil
}

// ReplaceAll returns a copy of a byte slice
// where all pattern matches are replaced by repl.
func (re Regexp) ReplaceAll(bytes, repl []byte, flags int) []byte {
	m := re.Matcher(bytes, flags)
	r := []byte{}
	for m.matches {
		r = append(append(r, bytes[:m.ovector[0]]...), repl...)
		bytes = bytes[m.ovector[1]:]
		m.Match(bytes, flags)
	}
	return append(r, bytes...)
}

// ReplaceAllString is equivalent to ReplaceAll with string return type.
func (re Regexp) ReplaceAllString(in, repl string, flags int) string {
	return string(re.ReplaceAll([]byte(in), []byte(repl), flags))
}

// CompileError holds details about a compilation error,
// as returned by the Compile function.  The offset is
// the byte position in the pattern string at which the
// error was detected.
type CompileError struct {
	Pattern string // The failed pattern
	Message string // The error message
	Offset  int    // Byte position of error
}

// Error converts a compile error to a string
func (e *CompileError) Error() string {
	return e.Pattern + " (" + strconv.Itoa(e.Offset) + "): " + e.Message
}
