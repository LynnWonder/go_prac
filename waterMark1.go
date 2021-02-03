package main

import (
"fmt"
"image"
"image/draw"
"image/jpeg"
"image/png"
"os"
"strings"
)
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func main() {
	//获取路径
	//fPath :=  os.Args[1];
	fPath :=  "/Users/bytedance/Desktop/test.png";

	//获取文件扩展名
	nPath := strings.Replace(fPath, ".", "_new.", -1 )
	r,_:= PathExists(fPath)
	if(!r){
		fmt.Println(fPath,"文件不存在\n");
		return
	}
	fmt.Println("输入文件路径:\t",fPath,"\nlogo路径:\tlogo.png\n输出文件路径:\t",nPath);

	//原始图片是sam.jpg
	imgb, _ := os.Open(fPath)
	img, _ := png.Decode(imgb)
	defer imgb.Close()

	wmb, _ := os.Open("/Users/bytedance/Desktop/draft.png")

	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-5, img.Bounds().Dy()-watermark.Bounds().Dy()-5)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create(nPath)
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()

	fmt.Println("水印添加结束,请查看new.jpg图片...")
	fmt.Println(nPath)
}
