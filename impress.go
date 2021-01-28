package main
import (
	"bufio"
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
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Get() io.Reader  { //生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url
	url := "https://open.feishu.cn/open-apis/image/v4/get?image_key=img_379f66f7-8169-478b-b381-00bf6a45502g"
	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	request.Header.Add("Authorization", "Bearer t-86e9e0e349751b66fd8ce6a75f61eeac2bbcfebd")

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
 目前看做成具体 robot 主要涉及细节是： 如何读取文件，生成文件后如何给 robot
 */

/**
	getReadSizeFile func() (io.Reader,error)
	getDecodeFile func() (*os.File,error),
	to: 输出目录
	Quality: 图片质量
	base int,
	format: 后缀格式
 */

func imageCompress(
	getReadSizeFile func() (io.Reader,error),
	getDecodeFile func() (*os.File,error),
	to string,
	Quality,
	base int,
	format string,
	outputType string) bool{
	/** 读取文件 */
	//file_origin, err := getDecodeFile()
	file, err := getDecodeFile()
	fmt.Printf("====>file %v\n", file)
	file_origin := Get()
	//fmt.Printf("====>file_origin %v\n", file_origin)
	//defer file_origin.Close()
	if err != nil {
		fmt.Println("os.Open(file)错误");
		log.Fatal(err)
		return false
	}
	var origin image.Image
	var config image.Config
	var temp io.Reader
	/** 读取尺寸 */
	temp, err = getReadSizeFile()
	if err != nil {
		fmt.Println("os.Open(temp)");
		log.Fatal(err)
		return false
	}
	var typeImage int64
	format = strings.ToLower(format)
	/** jpg 格式 */
	if format=="jpg" || format =="jpeg" {
		typeImage = 1
		origin, err = jpeg.Decode(file_origin)
		if err != nil {
			fmt.Println("jpeg.Decode(file_origin)");
			log.Fatal(err)
			return false
		}
		temp, err = getReadSizeFile()
		if err != nil {
			fmt.Println("os.Open(temp)")
			log.Fatal(err)
			return false
		}
		config,err = jpeg.DecodeConfig(temp)
		if err != nil {
			fmt.Println("jpeg.DecodeConfig(temp)")
			return false
		}
	}else if format=="png" {
		typeImage = 0
		origin, err = png.Decode(file_origin)
		if err != nil {
			fmt.Println("png.Decode(file_origin)");
			log.Fatal(err)
			return false
		}
		temp, err = getReadSizeFile()
		if err != nil {
			fmt.Println("os.Open(temp)");
			log.Fatal(err)
			return false
		}
		config,err = png.DecodeConfig(temp);
		if err != nil {
			fmt.Println("png.DecodeConfig(temp)");
			return false
		}
	}
	/** 做等比缩放 */
	width  := uint(base) /** 基准 */
	height := uint(base*config.Height/config.Width)

	var canvas image.Image
	fmt.Printf("width: %v, height: %v", width, height)
	if outputType=="thumbnail" {
		canvas = resize.Thumbnail(width, height, origin, resize.Lanczos3)
	}else if outputType=="fixed" {
		canvas = resize.Resize(240, 240, origin, resize.Lanczos3)
	}
	fmt.Printf("====> %v", canvas)
	//canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	defer file_out.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	if typeImage==0 {
		// todo 如果不进行编码可以直接生成一个图片然后进行写入
		err = png.Encode(file_out, canvas)
		if err!=nil {
			fmt.Println("压缩图片失败");
			return false
		}
	}else{
		err = jpeg.Encode(file_out, canvas, &jpeg.Options{Quality})
		if err!=nil {
			fmt.Println("压缩图片失败");
			return false
		}
	}

	return true
}









func getFilelist(path string) {
	/** 创建输出目录 */
	errC := os.MkdirAll(inputArgs.OutputPath, 0777)
	if errC != nil {
		fmt.Printf("%s", errC)
		return
	}
	err := filepath.Walk(path, func(pathFound string, f os.FileInfo, err error) error {
		if ( f == nil ) {
			return err
		}
		if f.IsDir() { /** 是否是目录 */
			return nil
		}
		// println(pathFound)
		/** 找到一个文件 */
		/** 判断是不是图片 */
		localPath,format,_ := isPictureFormat(pathFound)
		/** 随机数 */
		t := time.Now()
		millis := t.Nanosecond() /** 纳秒 */
		outputPath := inputArgs.OutputPath+strconv.FormatInt(int64(millis),10)+"."+format
		if localPath!="" {
			if !imageCompress(
				func() (io.Reader,error){
					return os.Open(localPath)
				},
				func() (*os.File,error) {
					return os.Open(localPath)
				},
				outputPath,
				inputArgs.Quality,
				inputArgs.Width,
				format,
				"fixed") {
				fmt.Println("生成缩略图失败")
			}else{
				fmt.Println("生成缩略图成功 "+outputPath)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("输入的路径信息有误 %v\n", err)
	}
}

/**
格式化输出图片路径
*/
func isPictureFormat(path string) (string,string,string) {
	// go 字符串也可以分割
	temp := strings.Split(path,".")
	if len(temp) <=1 {
		return "","",""
	}
	mapRule := make(map[string]int64)
	mapRule["jpg"]  = 1
	mapRule["png"]  = 1
	mapRule["jpeg"] = 1
	// 如果满足格式
	if mapRule[temp[1]] == 1  {
		println(temp[1])
		// 返回 完整路径，后缀，文件名
		return path,temp[1],temp[0]
	}else{
		// 如果不满足格式或者说是批量进行修改
		return "","",""
	}
}

func execute()  {
	/** 获取输入 */
	//str := ""
	//fmt.Scanln (&str) /** 不要使用 scanf，它不会并以一个新行结束输入 */

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	/** 分割 */
	strPice := strings.Split(string(data)," ") /** 空格 */
	//if len(strPice) < 3 {
	//	fmt.Printf("输入有误，参数数量不足,请重新输入或退出程序：")
	//	execute()
	//	return
	//}

	// 路径
	inputArgs.LocalPath = strPice[0]
	// 质量
	inputArgs.Quality,_ = strconv.Atoi(strPice[1])
	// 尺寸
	inputArgs.Width,_   = strconv.Atoi(strPice[2])

	// 路径，后缀格式，文件名
	pathTemp,format,top := isPictureFormat(inputArgs.LocalPath)
	if pathTemp == "" {
		// 批量文件处理
		fmt.Println("开始批量压缩...")
		rs := []rune(inputArgs.LocalPath)
		end := len(rs)
		substr := string(rs[end-1:end])
		if substr=="/" {
			/** 有 / */
			rs := []rune(inputArgs.LocalPath)
			end := len(rs)
			substr := string(rs[0:end-1])
			endIndex := strings.LastIndex(substr,"/")
			inputArgs.OutputPath = string(rs[0:endIndex])+"/LghImageCompress/"
		}else {
			endIndex := strings.LastIndex(inputArgs.LocalPath,"/")
			inputArgs.OutputPath = string(rs[0:endIndex])+"/LghImageCompress/"
		}
		getFilelist(inputArgs.LocalPath)
		fmt.Println("图片保存在文件夹 "+inputArgs.OutputPath)
	}else{
		/** 单个 */
		/** 如果输入文件，那么是单个，允许自定义路径 */
		fmt.Println("开始单张压缩...") //C:\Users\lzq\Desktop\Apk.jpg 75 200
		// 输出文件地址为
		inputArgs.OutputPath = top+"_compress."+format
		if !imageCompress(
			func() (io.Reader,error){
				return os.Open(inputArgs.LocalPath)
			},
			func() (*os.File,error) {
				return os.Open(inputArgs.LocalPath)
			},
			inputArgs.OutputPath,
			inputArgs.Quality,
			inputArgs.Width,
			format,
			"thumbnail") {
			fmt.Println("生成缩略图失败")
		}else{
			fmt.Println("生成缩略图成功 "+inputArgs.OutputPath)
			finish()
		}
	}

	time.Sleep(5 * time.Minute) /** 如果不是自己点击退出，延时5分钟 */
}

func finish()  {
	fmt.Printf("继续输入进行压缩或者退出程序：")
	execute()
}

func showTips()  {
	tips := []string{
		"请输入文件夹或图片路径:",
		"如果输入文件夹,那么该目录的图片将会被批量压缩;",
		"如果是图片路径，那么将会被单独压缩处理。",
		"例如：",
		"C:/Users/lzq/Desktop/headImages/ 75 200",
		"指桌面 headImages 文件夹，里面的图片质量压缩到75%，宽分辨率为200，高是等比例计算",
		"C:/Users/lzq/Desktop/headImages/1.jpg 75 200",
		"指桌面的 headImages 文件夹里面的 1.jpg 图片,质量压缩到75%，宽分辨率为200，高是等比例计算 ",
		"请输入："}
	itemLen := len(tips)
	for i :=0;i<itemLen;i++ {
		if i == itemLen -1 {
			fmt.Printf(tips[i])
		}else{
			fmt.Println(tips[i])
		}
	}
}

type InputArgs struct {
	OutputPath string  /** 输出目录 */
	LocalPath  string  /** 输入的目录或文件路径 */
	Quality    int     /** 质量 */
	Width      int     /** 宽度尺寸，像素单位 */
	Format     string  /** 格式 */
}

// 声明变量 var name type
var inputArgs InputArgs

func main() {
	showTips()
	execute()
	time.Sleep(5 * time.Minute) /** 如果不是自己点击退出，延时5分钟 */
}