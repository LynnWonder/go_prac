package main
import (
	"fmt"
	"reflect"
)
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
	fmt.Printf("%d %c\n", str2[2], str2[2])     // 232 è
	fmt.Println("len(str2)：", len(str2))       // len(str2)： 8
	// 正确的处理方式是将 string 转为 rune 数组，此时字符串中的每个字符，无论占多少个字节都用 int32 来表示，因而可以正确处理中文。
	runeArr := []rune(str2)
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind()) // int32
	fmt.Println(runeArr[2], string(runeArr[2]))    // 35821 语
	fmt.Println("len(runeArr)：", len(runeArr))    // len(runeArr)： 4
}