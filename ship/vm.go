package ship

/*
#cgo LDFLAGS: -L../lib -lhello
#include <stdio.h>
#include <stdlib.h>
#include "../lib/hello.h"
*/
import "C"
import "unsafe"

// Hello is a simple function that return the input value.
func Hello(s string) string {
	cstr := C.CString(s)
	rs := C.hello(cstr)
	str := C.GoString(rs)
	return str
}

// RunScript Run a test Move script
func RunScript(s string) string {
	cstr := C.CString(s)
	//defer C.free(unsafe.Pointer(cstr))
	rs := C.run_script(cstr)
	defer C.free(unsafe.Pointer(rs))
	str := C.GoString(rs)

	return str
}
