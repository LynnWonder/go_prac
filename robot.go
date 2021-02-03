package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
	"unsafe"

	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/opentype"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
)

type Res struct {
	Code                string
	Msg                 string
	Data                Data
	Tenant_access_token string
}
type Data struct {
	Image_key string
}

func HttpGet(url string, token string) io.Reader { //生成client 参数为默认
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", token)

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	img, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return bytes.NewReader(img)
}

func HttpPost(uri string, body io.Reader, token string, contentType string) Res {
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
	fmt.Printf("post res===> %v", *str)

	// json 解析
	_ = json.Unmarshal(respBytes, &res)
	return res
}

func uploadImg(paramName string, img io.Reader, token string) string {
	var res Res
	path := "/tmp"
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
	res = HttpPost("https://open.feishu.cn/open-apis/image/v4/put", body, token, writer.FormDataContentType())
	return res.Data.Image_key
}

/**
获取 token
*/
func getToken() string {
	var res Res
	body_str := `{
	"app_id":"cli_a087de86a539100c",
	"app_secret":"h5G6QuWE3KmKjU7hKXQg9eiz7oiq0A7A"
	}`
	// 转换为 byte 类型 再通过 bytes.NewReader(body) 转换为 io.reader 类型
	body := []byte(body_str)
	res = HttpPost("https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal/", bytes.NewReader(body), "", "application/json")
	return "Bearer " + res.Tenant_access_token
}
func postImgMsg(image_key string, token string) {
	body_str := `{
	"email":"liulin.wonder@bytedance.com",
	"msg_type": "image",
	"content":{
		"image_key": "` + image_key + `"
	}
	}`
	// 转换为 byte 类型 再通过 bytes.NewReader(body) 转换为 io.reader 类型
	body := []byte(body_str)
	HttpPost("https://open.feishu.cn/open-apis/message/v4/send/", bytes.NewReader(body), token, "application/json")
}
func postTextMsg(text string, token string) {
	body_str := `{
	"email":"liulin.wonder@bytedance.com",
	"msg_type": "text",
	"content":{
		"text": "我没太听懂你的意思"
	}
	}`
	// 转换为 byte 类型 再通过 bytes.NewReader(body) 转换为 io.reader 类型
	body := []byte(body_str)
	HttpPost("https://open.feishu.cn/open-apis/message/v4/send/", bytes.NewReader(body), token, "application/json")
}
func postInteractMsg(chat_id string, token string) {
	body_str := `{
	"chat_id":"` + chat_id + `",
	"msg_type": "interactive",
 	"card": {
		"header": {
			"title": {
				"tag": "plain_text",
				"content": "图片处理"
			}
		},
        "elements": [
			{
				"tag": "div",
				"text": {
					"tag": "plain_text",
					"content": "选择一项你需要进行的操作，然后按返回的提示进行"
				}
			},
            {
                "tag": "action",
                "actions": [
                    {
                        "tag": "button",
                        "text": {
                            "tag": "plain_text",
                            "content": "指定宽高图片"
                        },
                        "type": "primary",
						"value":{
							"key":"fixed"
						}
                    },
                    {
                        "tag": "button",
                        "text": {
                            "tag": "plain_text",
                            "content": "缩略图"
                        },
                        "type": "primary",
						"value":{
							"key":"thumbnail"
						}
                    }
        ]
    }
    ]
}
	}`
	// 转换为 byte 类型 再通过 bytes.NewReader(body) 转换为 io.reader 类型
	body := []byte(body_str)
	HttpPost("https://open.feishu.cn/open-apis/message/v4/send/", bytes.NewReader(body), token, "application/json")
}
func ImageCompress(
	base int,
	format string,
	outputType string,
	token string) io.Reader {
	var origin image.Image
	file_origin := HttpGet("https://open.feishu.cn/open-apis/image/v4/get?image_key=img_d326e1da-e12a-4203-acfc-e023f9b35f2g", token)
	format = strings.ToLower(format)
	/** jpg 格式 */
	//if format=="jpg" || format =="jpeg" {
	//	origin, _ = jpeg.Decode(file_origin)
	//}else if format=="png" {
	//	origin, _ = png.Decode(file_origin)
	//}
	// opted 图片 decode
	origin, _ = imaging.Decode(file_origin)
	///** 做等比缩放 */
	//width  := uint(base)
	//// 实际宽高
	//height := uint(base*240/240)

	var canvas image.Image
	//var col color.Color
	if outputType == "fixed" {
		canvas = imaging.Resize(origin, 240, 240, imaging.Lanczos)
	} else if outputType == "thumbnail" {
		canvas = imaging.Thumbnail(origin, 240, 120, imaging.Lanczos)
	}else if outputType == "blur" {
		// opted 图片高斯模糊
		canvas = imaging.Blur(origin, 15)
	}else if outputType == "sharpen" {
		// opted 图片锐化
		canvas = imaging.Sharpen(origin, 0.8)
	}else if outputType == "brightness" {
		// opted 图片亮度调节
		canvas = imaging.AdjustBrightness(origin, 20)
	}else if outputType == "gray"{
		// opted 图片灰度（一键遗照）
		canvas = imaging.Grayscale(origin)
	}else if outputType == "rotate" {
		// opted 图片逆时针旋转，背景部分默认设置为白色
		canvas = imaging.Rotate(origin, 90, color.RGBA{255, 255, 255, 255})
	}else if outputType == "invert" {
		// opted 图片反色
		canvas = imaging.Invert(origin)
	}
	//return canvas
	buf := new(bytes.Buffer)
	// 将 image.Image 转化为 []byte
	//_ = png.Encode(buf, canvas)
	// opted 修改成支持多种类型的图片转换
	_ = imaging.Encode(buf, canvas, 2)
	return buf
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

// WaterMark for adding a watermark on the image
func waterMark(image_key string, markText string, markSize float64, token string) (io.Reader, error) {
	file_origin := HttpGet("https://open.feishu.cn/open-apis/image/v4/get?image_key="+image_key, token)
	img, _ :=imaging.Decode(file_origin)
	// image's length to canvas's length
	bounds := img.Bounds()
	w := vg.Length(bounds.Max.X) * vg.Inch / vgimg.DefaultDPI
	h := vg.Length(bounds.Max.Y) * vg.Inch / vgimg.DefaultDPI
	diagonal := vg.Length(math.Sqrt(float64(w*w + h*h)))

	// create a canvas, which width and height are diagonal
	c := vgimg.New(diagonal, diagonal)

	// draw image on the center of canvas
	rect := vg.Rectangle{}
	rect.Min.X = diagonal/2 - w/2
	rect.Min.Y = diagonal/2 - h/2
	rect.Max.X = diagonal/2 + w/2
	rect.Max.Y = diagonal/2 + h/2
	c.DrawImage(rect, img)

	// make a fontStyle, which width is vg.Inch * 0.7
	fontStyle, _ := vg.MakeFont("Courier", vg.Inch*0.7)

	// repeat the markText
	markTextWidth := fontStyle.Width(markText)
	unitText := markText
	// 尽可能拉长距离让水印填充整张图片
	for markTextWidth <= w*5 {
		markText += "      " + unitText
		markTextWidth = fontStyle.Width(markText)
	}

	// set the color of markText
	c.SetColor(color.RGBA{0, 0, 0, 40})

	// set a random angle between 0 and π/2
	//θ := math.Pi * rand.Float64() / 2
	// 默认设置为 70 度
	c.Rotate(70)

	// set the lineHeight and add the markText
	lineHeight := fontStyle.Extents().Height * 1
	for offset := -2 * diagonal; offset < 2*diagonal; offset += lineHeight {
		var after font.Face
		f, _ := opentype.Parse(goitalic.TTF)
		after.Face = f
		var len font.Length
		len = font.Length(markSize)
		after.Font = font.Font{
			Typeface: "",
			Variant:  "Math",
			Style:    0,
			Weight:   20,
			Size:    len,
		}
		// 填充文字
		c.FillString(after, vg.Point{X: w/3, Y: offset}, markText)
	}

	// canvas writeto jpeg
	// canvas.img is private
	// so use a buffer to transfer
	jc := vgimg.PngCanvas{Canvas: c}
	buff := new(bytes.Buffer)
	jc.WriteTo(buff)
	img, _, err := image.Decode(buff)
	if err != nil {
		return nil, err
	}

	// get the center point of the image
	ctp := int(diagonal * vgimg.DefaultDPI / vg.Inch / 2)

	// cutout the marked image
	size := bounds.Size()
	bounds = image.Rect(ctp-size.X/2, ctp-size.Y/2, ctp+size.X/2, ctp+size.Y/2)
	rv := image.NewRGBA(bounds)
	draw.Draw(rv, bounds, img, bounds.Min, draw.Src)
	//return rv, nil
	buf := new(bytes.Buffer)
	// 将 image.Image 转化为 []byte
	_ = imaging.Encode(buf, rv, 2)
	return buf, nil
}

func main() {
	token := getToken()
	//file_origin := HttpGet("https://open.feishu.cn/open-apis/image/v4/get?image_key=img_94d78666-0ba3-4db0-bda9-7e408b7b67fg", token)
	//file_img, _ :=imaging.Decode(file_origin)
	img, _ := waterMark("img_0db4b3fd-5cb5-4467-826e-197eef09183g", "THIS IS A TEST", 20,token)
	//img := ImageCompress(
	//	240,
	//	"png",
	//	"invert", token)
	//data := make(map[string]string)
	img_key := uploadImg("image", img, token)
	postImgMsg(img_key, token)
	//postTextMsg("sdd", token)
	//postInteractMsg("6921306919621722115", token)
}
