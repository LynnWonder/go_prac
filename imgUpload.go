package main

import (
"bytes"
"fmt"
"io"
"io/ioutil"
"mime/multipart"
"net/http"
"os"
"unsafe"
)

func main()  {
	data := make(map[string]string)
	newFileUploadRequest(data,"image","/Users/bytedance/Desktop/test.png")

}

func Get() io.Reader  { //生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url
	url := "https://open.feishu.cn/open-apis/image/v4/get?image_key=img_379f66f7-8169-478b-b381-00bf6a45502g"
	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	request.Header.Add("Authorization", "Bearer t-a215f7bfea3b8af18df9d9e33fdb1711b7bbee74")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	img,err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return bytes.NewReader(img)
}

func newFileUploadRequest(params map[string]string, paramName, path string)  error {
	uri :="https://open.feishu.cn/open-apis/image/v4/put"
	file, err := os.Open(path)
	if err != nil {
		return  err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return  err
	}
	//_, err = io.Copy(part, file)
	_, err = io.Copy(part, Get())

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
	request.Header.Set("Authorization", "Bearer t-6fd30ee3784982904166bab48b55e416c68a3d78")


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