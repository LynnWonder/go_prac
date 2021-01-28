package main

import (
	"bytes"
	"encoding/json"
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

type Res struct {
	Code string
	Msg string
	Data Data
}
type Data struct {
	Image_key string
}

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

func HttpPost (uri string, body io.Reader, token string, contentType string) Data{
	var res Res
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Authorization", "Bearer t-1a871f6970b73ac44411eb9c24b62b489ec9ce02")


	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

	// json 解析
	_= json.Unmarshal(respBytes, &res)
	fmt.Println(res.Data.Image_key)
	return  res.Data
}

func uploadImg(params map[string]string, paramName string, img io.Reader)  string {
	var data Data
	path :="/tmp"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = io.Copy(part, img)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.WriteField("image_type", "message")
	err = writer.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	data = HttpPost("https://open.feishu.cn/open-apis/image/v4/put",body,"Bearer t-1a871f6970b73ac44411eb9c24b62b489ec9ce02",writer.FormDataContentType())
	return data.Image_key
}
//func getToken()  {
//	var data Data
//	data = HttpPost("https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/",)
//}
func ImageCompress(
	base int,
	format string,
	outputType string) io.Reader{
	var origin image.Image
	file_origin := HttpGet("https://open.feishu.cn/open-apis/image/v4/get?image_key=img_379f66f7-8169-478b-b381-00bf6a45502g", "Bearer t-1a871f6970b73ac44411eb9c24b62b489ec9ce02")
	format = strings.ToLower(format)
	/** jpg 格式 */
	if format=="jpg" || format =="jpeg" {
		origin, _ = jpeg.Decode(file_origin)
	}else if format=="png" {
		origin, _ = png.Decode(file_origin)
	}
	///** 做等比缩放 */
	//width  := uint(base)
	//// 实际宽高
	//height := uint(base*240/240)

	var canvas image.Image
	if outputType=="thumbnail" {
		canvas = imaging.Resize(origin, 128, 128, imaging.Lanczos)
	}else if outputType=="fixed" {
		canvas = imaging.Thumbnail(origin, 240, 240, imaging.Lanczos)
	}
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
	data := make(map[string]string)
	uploadImg(data,"image", img)
}