package main

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Get(url string,token string) io.Reader  { //生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url
	//url := "https://open.feishu.cn/open-apis/image/v4/get?image_key=img_379f66f7-8169-478b-b381-00bf6a45502g"
	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	//request.Header.Add("Authorization", "Bearer t-a215f7bfea3b8af18df9d9e33fdb1711b7bbee74")
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
/**
格式化输出图片路径
*/
func isPictureFormat(path string) (string) {
	// go 字符串也可以分割
	temp := strings.Split(path,".")
	if len(temp) <=1 {
		return ""
	}
	mapRule := make(map[string]int64)
	mapRule["jpg"]  = 1
	mapRule["png"]  = 1
	mapRule["jpeg"] = 1
	// 如果满足格式
	if mapRule[temp[1]] == 1  {
		println(temp[1])
		// 返回后缀
		return temp[1]
	}else{
		// 如果不满足格式或者说是批量进行修改
		return ""
	}
}
func imageCompress(
	base int,
	format string,
	outputType string) image.Image{
	var origin image.Image
	file_origin := Get("https://open.feishu.cn/open-apis/image/v4/get?image_key=img_379f66f7-8169-478b-b381-00bf6a45502g", "Bearer t-ba5f1063b84ba08842018382bf1901a841942ee0")
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
		canvas = resize.Thumbnail(width, height, origin, resize.Lanczos3)
	}else if outputType=="fixed" {
		canvas = resize.Resize(240, 240, origin, resize.Lanczos3)
	}
	fmt.Printf("====>canvas %v",canvas)
	return canvas
}
func main()  {
	img :=imageCompress(
		240,
		"png",
		"thumbnail")

}