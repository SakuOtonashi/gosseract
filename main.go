package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/SakuOtonashi/gosseract/leptonica"
	"github.com/SakuOtonashi/gosseract/tesseract"
	"github.com/google/uuid"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Listen: " + ln.Addr().String())
	sm := http.NewServeMux()
	sm.Handle("/ocr/api", ocrApi{})
	srv := &http.Server{Handler: sm}
	err = srv.Serve(ln)
	if err != nil {
		panic(err)
	}
}

type ocrApi struct{}

func (ocr ocrApi) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	var v struct {
		Language  string
		ImagePath string
	}
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte("can't parse body: " + err.Error()))
		return
	}

	// 直接使用路径的话需要处理编码问题 utf8 -> gbk
	// image, err := leptonica.PixRead(v.ImagePath)
	data, err := os.ReadFile(v.ImagePath)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte("can't read file: " + err.Error()))
		return
	}
	image, err := leptonica.PixReadMem(data)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte("can't read image: " + err.Error()))
		return
	}
	defer image.Destroy()

	api := tesseract.NewTessBaseAPI()
	defer api.Delete()
	err = api.Init3("", v.Language)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("can't init tesseract: " + err.Error()))
		return
	}
	defer api.End()

	api.SetImage2(image)
	// outText := api.GetUTF8Text()
	boxa := api.GetComponentImages(tesseract.RIL_TEXTLINE, true, nil, nil)
	if boxa == nil {
		res := ocrApiResponse{
			Code:      -1,
			Message:   "can't get any text",
			RequestId: uuid.NewString(),
		}
		text, err := json.Marshal(res)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte("can't marshal response: " + err.Error()))
			return
		}
		rw.Write(text)
		return
	}
	outData := make([]ocrApiDataItem, boxa.N)
	for i := int32(0); i < boxa.N; i++ {
		box := boxa.GetBox(i, leptonica.L_CLONE)
		api.SetRectangle(box.X, box.Y, box.W, box.H)
		ocrResult := strings.TrimSuffix(api.GetUTF8Text(), "\n")
		conf := api.MeanTextConf()
		outData[i] = ocrApiDataItem{Words: ocrResult, Conf: conf, Box: ocrApiDataItemBox{box.X, box.Y, box.W, box.H}}
		box.Destroy()
	}
	boxa.Destroy()
	fmt.Println(v.Language + " OCR output:\n")
	for _, data := range outData {
		fmt.Println("得分: ", data.Conf, "\n文字: ", data.Words)
	}

	res := ocrApiResponse{
		Code:      0,
		Data:      outData,
		Message:   "Success",
		RequestId: uuid.NewString(),
	}
	text, err := json.Marshal(res)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("can't marshal response: " + err.Error()))
		return
	}
	rw.Write(text)
}

type ocrApiResponse struct {
	Code      int
	Data      []ocrApiDataItem
	Message   string
	RequestId string
}

type ocrApiDataItem struct {
	Words string
	Conf  int32
	Box   ocrApiDataItemBox
}

type ocrApiDataItemBox struct {
	X int32
	Y int32
	W int32
	H int32
}
