# go_prac


golang 练习使用

## 基本数据类型和引用类型

### 引用类型
注：引用类型零值即为 nil
1. slice
2. map

## 数组和切片的区别

数组类型的值（以下简称数组）的长度是固定的数组的长度在声明它的时候就必须给定，并且在之后不会再改变。
可以说，数组的长度是其类型的一部分（数组的容量永远等于其长度，都是不可变的）

切片类型的值是可变长的。而切片的类型字面量中只有其元素的类型，而没有其长度。
切片的长度可以自动地随着其中元素数量的增长而增长，但不会随着元素数量的减少而减少。
尤其需要注意的是左闭右开的

在每一个切片的底层数据结构中，会包含一个数组，可以被叫做底层数据，而切片就是对底层数组的引用，故而切片类型属于引用类型，
因此修改切片会修改原数组。

```
/*
// 切片的结构体
type slice struct {
    array unsafe.Pointer // 底层数据的指针
    len   int // 切的长度
    cap   int // 截取底层数据的容量
}

*/
```

### 切片掌握内容

1. 最好添加第三个参数即容量以避免导致出现访问不该暴露的元素
2. append
3. copy 遵循最小复制原则，被复制的元素个数总是等于长度较短的那个参数值的长度

## 格式化输出
``Println``可以打印出字符串，和变量


``Printf``只可以打印出格式化的字符串,可以输出字符串类型的变量


``％d``整型输出，``％ld``长整型输出，

``％o``以八进制数形式输出整数，

``％x``以小写十六进制数形式输出整数，

``％X``以大写十六进制数形式输出整数，

``％u``以十进制数输出unsigned型数据(无符号数)。

``％c``用来输出一个字符，

``％s``用来输出一个字符串，

``％f``用来输出实数，以小数形式输出，（备注：浮点数是不能定义如的精度的，所以“%6.2f”这种写法是“错误的”！！！）

``％e``以指数形式输出实数，

``％g``根据大小自动选f格式或e格式，且不输出无意义的零。

``%v``	使用默认格式输出值，或者如果方法存在，则使用类性值的String()方法输出自定义值


## TODO

1. go 的基本类型和引用类型区分
2. functions 学习，多个返回值学习
3. 声明变量的方式 尤其是冒号
4. 指针 
5. slice map struct 协程
6. map 在并发会 panic 原因？
7. switch 不需要进行 break
8. 在 for 循环中写函数也会面临闭包问题，当然也有相应的解决方案
9. 结构体，c++ 中的结构体