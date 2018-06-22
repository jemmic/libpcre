# Go port of the PCRE library

The regular expression library PCRE (https://www.pcre.org/) has been transpiled from *C* to *Go* code.

This was done, by using *c2go* (https://github.com/elliotchance/c2go) to perform the majority of the work.
Some manual refinement was performed, to fix some remaining errors and to be able to launch the tests with `go test`.

An API similar to the Go regex functions was added based on the Go C binding for pcre (https://github.com/gijsbers/go-pcre).
