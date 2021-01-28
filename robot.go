package main

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"unsafe"
)

func HttpGet(url string,token string) io.Reader  { //生成client 参数为默认
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", token)

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	img,err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return bytes.NewReader(img)
}

func HttpPost(params map[string]string, paramName, path string, img io.Reader)  error {
	uri :="https://open.feishu.cn/open-apis/image/v4/put"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return  err
	}
	_, err = io.Copy(part, img)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.WriteField("image_type", "message")
	err = writer.Close()
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer t-86e9e0e349751b66fd8ce6a75f61eeac2bbcfebd")


	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

	return  err
}
func ImageCompress(
	base int,
	format string,
	outputType string) io.Reader{
	var origin image.Image
	file_origin := HttpGet("https://open.feishu.cn/open-apis/image/v4/get?image_key=img_379f66f7-8169-478b-b381-00bf6a45502g", "Bearer t-86e9e0e349751b66fd8ce6a75f61eeac2bbcfebd")
	format = strings.ToLower(format)
	/** jpg 格式 */
	if format=="jpg" || format =="jpeg" {
		origin, _ = jpeg.Decode(file_origin)
	}else if format=="png" {
		origin, _ = png.Decode(file_origin)
	}
	/** 做等比缩放 */
	width  := uint(base)
	// 实际宽高
	height := uint(base*240/240)

	var canvas image.Image
	fmt.Printf("width: %v, height: %v", width, height)
	if outputType=="thumbnail" {
		canvas = imaging.Resize(origin, 128, 128, imaging.Lanczos)
		//canvas = resize.Thumbnail(width, height, origin, resize.Lanczos3)
	}else if outputType=="fixed" {
		canvas = imaging.Thumbnail(origin, 240, 240, imaging.Lanczos)
		//canvas = resize.Resize(240, 240, origin, resize.Lanczos3)
	}
	fmt.Printf("====>canvas %v\n",canvas)
	//return canvas
	buf := new(bytes.Buffer)
	// 将 image.Image 转化为 []byte
	_ = png.Encode(buf, canvas)
	return buf
}
func main()  {
	img :=ImageCompress(
		240,
		"png",
		"thumbnail")
	fmt.Printf("====img %v", img)
	data := make(map[string]string)
	HttpPost(data,"image","/Users/bytedance/Desktop/111.png", img)
}