package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get()([]byte, error)  { //生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url
	url := "https://open.feishu.cn/open-apis/image/v4/get?image_key=img_e2186396-bc78-4d4f-ac24-371783fc14cg"
	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	request.Header.Add("Authorization", "Bearer t-c62dff385d9e12dcf7b0a9078cf0f1f760da656f")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()
	respByte, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("===> %v", respByte)
	return ioutil.ReadAll(response.Body)
}