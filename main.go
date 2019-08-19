package main

/*
#include <emacs-module.h>
#include <stdlib.h>
#include <string.h>

static inline bool
copy_string_contents (emacs_env * env, emacs_value value, char *buf, ptrdiff_t *size)
{
        return env->copy_string_contents (env, value, buf, size);
}

static inline emacs_value
make_string (emacs_env * env, const char *s)
{
        return env->make_string (env, s, strlen(s));
}

static inline emacs_value
intern(emacs_env *env, const char* name) {
        return env->intern (env, name);
}
*/
import "C"
import (
	"encoding/base64"
	"reflect"
	"unsafe"
)

//export Fb64_encode
func Fb64_encode(env *C.emacs_env, nargs C.ptrdiff_t, args *C.emacs_value) C.emacs_value {
	var lisp_args []C.emacs_value
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&lisp_args))
	slice.Cap = int(nargs)
	slice.Len = int(nargs)
	slice.Data = uintptr(unsafe.Pointer(args))
	var lisp_str C.emacs_value = lisp_args[0]
	var size C.ptrdiff_t = 0
	var buf *C.char

	C.copy_string_contents(env, lisp_str, buf, &size)
	buf = (*C.char)(C.malloc(C.size_t(size)))
	defer C.free(unsafe.Pointer(buf))
	C.copy_string_contents(env, lisp_str, buf, &size)

	data := C.GoString(buf)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	cencoded := C.CString(encoded)
	defer C.free(unsafe.Pointer(cencoded))

	return C.make_string(env, cencoded)
}

//export Fb64_decode
func Fb64_decode(env *C.emacs_env, nargs C.ptrdiff_t, args *C.emacs_value) C.emacs_value {
	var lisp_args []C.emacs_value
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&lisp_args))
	slice.Cap = int(nargs)
	slice.Len = int(nargs)
	slice.Data = uintptr(unsafe.Pointer(args))
	var lisp_str C.emacs_value = lisp_args[0]
	var size C.ptrdiff_t = 0
	var buf *C.char

	C.copy_string_contents(env, lisp_str, buf, &size)
	buf = (*C.char)(C.malloc(C.size_t(size)))
	defer C.free(unsafe.Pointer(buf))
	C.copy_string_contents(env, lisp_str, buf, &size)

	data := C.GoString(buf)
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		cnil := C.CString("nil")
		defer C.free(unsafe.Pointer(cnil))
		return C.intern(env, cnil)
	}

	cdecoded := C.CString(string(decoded))
	defer C.free(unsafe.Pointer(cdecoded))

	return C.make_string(env, cdecoded)
}

func main() {}
