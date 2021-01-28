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
	Tenant_access_token string
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

func HttpPost (uri string, body io.Reader, token string, contentType string) Res{
	var res Res
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Authorization", token)


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
	fmt.Printf("post res===> %v",*str)

	// json 解析
	_= json.Unmarshal(respBytes, &res)
	return  res
}

func uploadImg(paramName string, img io.Reader, token string)  string {
	var res Res
	path :="/tmp"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = io.Copy(part, img)

	// params map[string]string, 入参
	//for key, val := range params {
	//	_ = writer.WriteField(key, val)
	//}
	writer.WriteField("image_type", "message")
	err = writer.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	res = HttpPost("https://open.feishu.cn/open-apis/image/v4/put",body,token,writer.FormDataContentType())
	return res.Data.Image_key
}
/**
	获取 token
 */
func getToken() string{
	var res Res
	body_str :=`{
	"app_id":"cli_9e5cfce5663c100d",
	"app_secret":"6FYifQ7IBCPNyDe5zvFiSmXevq5Ic8Iq"
	}`
	// 转换为 byte 类型 再通过 bytes.NewReader(body) 转换为 io.reader 类型
	body :=[]byte(body_str)
	res = HttpPost("https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal/", bytes.NewReader(body), "", "application/json")
	return "Bearer "+res.Tenant_access_token
}
func postImgMsg(image_key string, token string)  {
	body_str :=`{
	"email":"liulin.wonder@bytedance.com",
	"msg_type": "image",
	"content":{
		"image_key": "`+image_key+`"
	}
	}`
	// 转换为 byte 类型 再通过 bytes.NewReader(body) 转换为 io.reader 类型
	body :=[]byte(body_str)
	HttpPost("https://open.feishu.cn/open-apis/message/v4/send/", bytes.NewReader(body),token,"application/json")
}
func postTextMsg(text string, token string)  {
	body_str :=`{
	"email":"liulin.wonder@bytedance.com",
	"msg_type": "text",
	"content":{
		"text": "我没太听懂你的意思"
	}
	}`
	// 转换为 byte 类型 再通过 bytes.NewReader(body) 转换为 io.reader 类型
	body :=[]byte(body_str)
	HttpPost("https://open.feishu.cn/open-apis/message/v4/send/", bytes.NewReader(body),token,"application/json")
}
func ImageCompress(
	base int,
	format string,
	outputType string,
	token string) io.Reader{
	var origin image.Image
	file_origin := HttpGet("https://open.feishu.cn/open-apis/image/v4/get?image_key=img_a1280840-3b1f-48aa-937c-4603248b206g", token)
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
	if outputType=="fixed" {
		canvas = imaging.Resize(origin, 240, 240, imaging.Lanczos)
	}else if outputType=="thumbnail" {
		canvas = imaging.Thumbnail(origin, 240, 120, imaging.Lanczos)
	}
	//return canvas
	buf := new(bytes.Buffer)
	// 将 image.Image 转化为 []byte
	_ = png.Encode(buf, canvas)
	return buf
}
func main()  {
	token :=getToken()
	img :=ImageCompress(
		240,
		"png",
		"thumbnail", token)
	//data := make(map[string]string)
	//img_key :=uploadImg(data,"image", img)
	img_key :=uploadImg("image", img, token)
	postImgMsg(img_key, token)
	postTextMsg("sdd", token)
}