//go:build !windows

package leptonica

/*
#include <stdlib.h>
#include <leptonica/allheaders.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

func PixRead(filename string) (*Pix, error) {
	if filename == "" {
		return nil, errors.New("filename is empty")
	}
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))

	pix := C.pixRead(cfn)
	if pix == nil {
		return nil, errors.New("pixRead failed")
	}
	return (*Pix)(pix), nil
}
