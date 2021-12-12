package leptonica

/*
#include <stdlib.h>
#include <leptonica/allheaders.h>
*/
import "C"
import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"unsafe"
)

func PixRead(filename string) (*Pix, error) {
	if filename == "" {
		return nil, errors.New("filename is empty")
	}

	// leptonica 在 windows 下接受的文件不是标准库里的文件对象
	// 也不接受 non-ANSI 格式的文件名，只能复制一份到临时文件夹里
	sf, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer sf.Close()
	ext := filepath.Ext(filename)
	tmp, err := os.CreateTemp("", "*"+ext)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(tmp, sf)
	tmp.Close()
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	if err != nil {
		return nil, err
	}

	cfn := C.CString(tmpPath)
	defer C.free(unsafe.Pointer(cfn))

	pix := C.pixRead(cfn)
	if pix == nil {
		return nil, errors.New("pixRead failed")
	}
	return (*Pix)(pix), nil
}
