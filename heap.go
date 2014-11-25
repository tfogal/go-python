package python

/*
#include <Python.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

const (
	PYTHON_API_VERSION = 1013
)

func Py_InitModule(name string) (*PyObject, error) {
	c_mname := C.CString(name)
	defer C.free(unsafe.Pointer(c_mname))
	// Py_InitModule doesn't actually exist; it's just a macro.
	obj := togo(C.Py_InitModule4(c_mname, nil, nil, nil, PYTHON_API_VERSION))
	if obj == nil {
		return nil, errors.New("internal error; module creation failed.")
	}
	return obj, nil
}

// EOF
