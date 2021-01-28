package main

import (
"bytes"
"fmt"
"io"
"io/ioutil"
"mime/multipart"
"net/http"
"unsafe"
)

func main()  {
	data := make(map[string]string)
	newFileUploadRequest(data,"image","/Users/bytedance/Desktop/111.png")

}

func Get() io.Reader  { //生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url
	url := "https://open.feishu.cn/open-apis/image/v4/get?image_key=img_d1f4de3a-4e0e-413a-9337-d4c33c206d1g"
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

func newFileUploadRequest(params map[string]string, paramName, path string)  error {
	uri :="https://open.feishu.cn/open-apis/image/v4/put"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return  err
	}
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