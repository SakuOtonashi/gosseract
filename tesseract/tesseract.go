package tesseract

/*
#cgo pkg-config: tesseract

#include <stdlib.h>
#include <tesseract/capi.h>
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/SakuOtonashi/gosseract/leptonica"
)

type TessBaseAPI C.TessBaseAPI

type TessPageIteratorLevel C.TessPageIteratorLevel

const (
	RIL_BLOCK    = C.RIL_BLOCK
	RIL_PARA     = C.RIL_PARA
	RIL_TEXTLINE = C.RIL_TEXTLINE
	RIL_WORD     = C.RIL_WORD
	RIL_SYMBOL   = C.RIL_SYMBOL
)

func TessVersion() string {
	return C.GoString(C.TessVersion())
}

func NewTessBaseAPI() *TessBaseAPI {
	return (*TessBaseAPI)(C.TessBaseAPICreate())
}

func (api *TessBaseAPI) Delete() {
	C.TessBaseAPIDelete((*C.TessBaseAPI)(api))
}

func (api *TessBaseAPI) Init3(datapath, language string) error {
	var dp *C.char
	if datapath != "" {
		dp = C.CString(datapath)
		defer C.free(unsafe.Pointer(dp))
	}

	var lang *C.char
	if language != "" {
		lang = C.CString(language)
		defer C.free(unsafe.Pointer(lang))
	}

	code := C.TessBaseAPIInit3((*C.TessBaseAPI)(api), dp, lang)
	if code != 0 {
		return errors.New("failed to initialize tesseract")
	}
	return nil
}

func (api *TessBaseAPI) SetImage2(image *leptonica.Pix) {
	pix := (*C.struct_Pix)(unsafe.Pointer(image))
	C.TessBaseAPISetImage2((*C.TessBaseAPI)(api), pix)
}

func (api *TessBaseAPI) SetRectangle(left, top, width, height int32) {
	C.TessBaseAPISetRectangle((*C.TessBaseAPI)(api), C.int(left), C.int(top), C.int(width), C.int(height))
}

func (api *TessBaseAPI) GetComponentImages(
	level TessPageIteratorLevel, text_only bool,
	_pixa interface{}, _blockids interface{}) *leptonica.Boxa {
	textOnly := C.int(0)
	if text_only {
		textOnly = 1
	}
	boxa := C.TessBaseAPIGetComponentImages((*C.TessBaseAPI)(api), C.TessPageIteratorLevel(level), textOnly, nil, nil)
	return (*leptonica.Boxa)(unsafe.Pointer(boxa))
}

func (api *TessBaseAPI) GetUTF8Text() string {
	text := C.TessBaseAPIGetUTF8Text((*C.TessBaseAPI)(api))
	return C.GoString(text)
}

func (api *TessBaseAPI) MeanTextConf() int32 {
	conf := C.TessBaseAPIMeanTextConf((*C.TessBaseAPI)(api))
	return int32(conf)
}

func (api *TessBaseAPI) End() {
	C.TessBaseAPIEnd((*C.TessBaseAPI)(api))
}
