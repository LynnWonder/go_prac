package main

import (
	"bytes"
	"flag"
	"fmt"
	"reflect"
	"unicode/utf8"
)

func append(slice, data []byte) []byte {
	// append 的实现原理
	// 直接使用比切片长度还要大的下标时，会报错内存溢出
	// 所以要先 make 一个新的切片
	lengthNewSlice := len(slice) + len(data)
	// make 新的切片时需注意总容量的大小，如果大于原切片，需扩充
	capNewSlice := cap(slice)
	if lengthNewSlice > cap(slice) {
		capNewSlice = lengthNewSlice
	}

	// 经测试，数据类型不能作为变量传递进来，所以应该用switch来实现，此处不再赘述
	newSlice := make([]byte, lengthNewSlice, capNewSlice)
	// 接下来赋值
	for sliceKey, sliceItem := range slice {
		newSlice[sliceKey] = sliceItem
	}
	for dataKey, item := range data {
		newSlice[dataKey+len(slice)] = item
	}
	// 赋值操作也可以用copy函数
	return newSlice
}

func main() {
	/**
	在 Go 语言中，字符串使用 UTF8 编码，
	UTF8 的好处在于，如果基本是英文，每个字符占 1 byte，和 ASCII 编码是一样的，非常节省空间，
	如果是中文，一般占3字节。包含中文的字符串的处理方式与纯 ASCII 码构成的字符串有点区别
	https://flaviocopes.com/javascript-unicode/#:~:text=While%20a%20JavaScript%20source%20file,single%20UTF%2D16%20code%20unit.
	在 JavaScript 语言中，JavaScript strings are all UTF-16 sequences, as the ECMAScript standard says:
	When a String contains actual textual data, each element is considered to be a single UTF-16 code unit.
	*/
	str1 := "Golang"
	str2 := "Go语言"
	// reflect.TypeOf().Kind() 可以知道某个变量的类型
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) // uint8
	fmt.Println(str1[2], string(str1[2]))       // 108 l
	// 注意字符串是以 byte 数组的形式存储的，所以有时候得到的并不是我们想要的数据
	fmt.Printf("%d %c\n", str2[2], str2[2]) // 232 è
	fmt.Println("len(str2)：", len(str2))    // len(str2)： 8
	// TIP 显示汉语
	//  正确的处理方式是将 string 转为 rune 数组，
	//  此时字符串中的每个字符，无论占多少个字节都用 int32 来表示，因而可以正确处理中文。
	//  注意此时会重新分配内存，并复制字节数组
	runeArr := []rune(str2)
	fmt.Println("=== TypeOf runeArr is ===", reflect.TypeOf(runeArr[2]).Kind()) // int32
	fmt.Println(runeArr[2], string(runeArr[2]))                                 // 35821 语
	fmt.Println("len(runeArr)：", len(runeArr))                                  // len(runeArr)： 4
	// 使用 len 拿到的有时候不是我们预期的字符串长度，此时可以使用 utf8.RuneCountInString，注意一些特殊字符还是可能会占有两个 rune
	fmt.Println("utf8.RuneCountInString===>", utf8.RuneCountInString(str2))
	mode := flag.String("mode", "http", "请填写服务的运行模式http|cron")
	// 可以得出结论是解指针后如果不输入默认就是 http
	fmt.Println("=====>", *mode)

	// range 迭代字符串中的值
	for _, v := range []rune(str1) {
		fmt.Println("======>str1 中的值", string(v))
	}

	// TIP 直接用字符串的索引来修改是不允许的，若要修改字符串必须将其转化为 []byte 或 []rune 后
	s1 := "hello"
	runeS1 := []rune(s1)
	runeS1[0] = 'g'
	fmt.Println("===转换后的数据是===", string(runeS1), s1)

	// TIP 使用 bytes 包可以辅助字符串操作，bytes.Buffer 提供 Read 和 Write 方法，因为读写长度未知的 bytes 最好使用 buffer
	//  通过 buffer.WriteString(s) 方法将字符串 s 追加到后面，最后再通过 buffer.String() 方法转换为 string
	//  这种方法比 += 更节省内存和 CPU
	sl := []byte{1, 2}
	data := []byte{3, 4}
	res := append(sl, data)
	strSl := []string{"hello", "world"}
	fmt.Printf("===>%v", res)

	var buffer bytes.Buffer
	for i := 0; i < len(strSl); i++ {
		buffer.WriteString(strSl[i])
	}
	fmt.Printf("buffer %s", buffer.String())

	buffer.WriteString("hello world")
	part1, part2 := buffer.Bytes()[:5], buffer.Bytes()[6:]
	fmt.Println(string(part1), string(part2))
}
