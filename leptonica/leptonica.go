package leptonica

/*
#cgo pkg-config: lept

#include <stdlib.h>
#include <stdio.h>
#include <leptonica/allheaders.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Pix C.struct_Pix

type Box struct {
	X        int32
	Y        int32
	W        int32
	H        int32
	Refcount uint32
}

type Boxa struct {
	N        int32
	Nalloc   int32
	Refcount uint32
	Box      **Box
}

const (
	L_NOCOPY     = C.L_NOCOPY
	L_INSERT     = C.L_INSERT
	L_COPY       = C.L_COPY
	L_CLONE      = C.L_CLONE
	L_COPY_CLONE = C.L_COPY_CLONE
)

func (box *Box) Destroy() {
	p := (*C.struct_Box)(unsafe.Pointer(box))
	C.boxDestroy(&p)
}

func (boxa *Boxa) Destroy() {
	p := (*C.struct_Boxa)(unsafe.Pointer(boxa))
	C.boxaDestroy(&p)
}

func (boxa *Boxa) GetBox(index, accessflag int32) *Box {
	box := C.boxaGetBox((*C.struct_Boxa)(unsafe.Pointer(boxa)), C.int(index), C.int(accessflag))
	return (*Box)(unsafe.Pointer(box))
}

func (pix *Pix) Destroy() {
	p := (*C.struct_Pix)(pix)
	C.pixDestroy(&p)
}

func PixRead(filename string) (*Pix, error) {
	if filename == "" {
		return nil, errors.New("filename is empty")
	}
	fn := C.CString(filename)
	defer C.free(unsafe.Pointer(fn))

	pix := C.pixRead(fn)
	if pix == nil {
		return nil, errors.New("pixRead failed")
	}
	return (*Pix)(pix), nil
}

func PixReadMem(data []byte) (*Pix, error) {
	d := (*C.uchar)(unsafe.Pointer(&data[0]))
	l := C.ulonglong(len(data))
	pix := C.pixReadMem(d, l)
	if pix == nil {
		return nil, errors.New("pixReadMem failed")
	}
	return (*Pix)(pix), nil
}

func LeptVersion() string {
	return C.GoString(C.getLeptonicaVersion())
}
