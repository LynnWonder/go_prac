package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/golang/freetype"
)

func httpGet(url string, token string) io.Reader { //生成client 参数为默认
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

func main() {
	//需要加水印的图片
	imgfile, _ := os.Open("/Users/bytedance/Desktop/test.png")
	defer imgfile.Close()

	jpgimg, _ := png.Decode(imgfile)

	img := image.NewNRGBA(jpgimg.Bounds())

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, jpgimg.At(x, y))
		}
	}
	//拷贝一个字体文件到运行目录
	//fontBytes, err := ioutil.ReadFile("/Users/bytedance/Desktop/light.ttc")
	fontBytes := httpGet("http://voffline.byted.org/download/tos/schedule//inspirecloud-cn-bytedance-internal/baas/ttke36/0cee611bc2746abd_1612347517947.ttc","")
	//if err != nil {
	//	log.Println(err)
	//}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
	}

	f := freetype.NewContext()
	f.SetDPI(72)
	f.SetFont(font)
	f.SetFontSize(50)
	f.SetClip(jpgimg.Bounds())
	f.SetDst(img)
	f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 100}))

	pt := freetype.Pt(img.Bounds().Dx()-800, img.Bounds().Dy()-200)
	_, err = f.DrawString("2021.02.03 hackday", pt)

	//draw.Draw(img,jpgimg.Bounds(),jpgimg,image.ZP,draw.Over)

	//保存到新文件中
	newfile, _ := os.Create("/Users/bytedance/Desktop/aaa.png")
	defer newfile.Close()

	err = png.Encode(newfile, img)
	if err != nil {
		fmt.Println(err)
	}
}