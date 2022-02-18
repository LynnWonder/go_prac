# go_prac


golang 练习使用

> 2022.02.17 开始系统的学习 golang，💪💪💪
## 值类型和引用类型
### 值类型
值类型：变量直接存储值，内存通常在栈中分配（属于值类型的数据类型有：int、float、bool、string、数组以及struct）


### 引用类型
引用类型：变量存储的是一个地址，这个地址存储最终的值，内存通常在堆中分配，通过GC回收（属于引用类型的的数据类型有：指针、slice、map、chan等）
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
  切片的结构体,
  结构体可以容纳许多不同的数据值。
  在过去，面向对象编程的应用尚未普及之前，我们通常将一些从逻辑上连接在一起的数据组合到一个单元中。
  一旦结构体类型被声明并且其数据成员被标识，即可创建该类型的多个变量，就像可以为同一个类创建多个对象一样。
*/
type slice struct {
    array unsafe.Pointer // 底层数据的指针
    len   int // 切的长度
    cap   int // 截取底层数据的容量
}
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

## go 的特殊表示
interface{}作为Go的重要特性之一，它代表的是一个类似``*void``的指针，可以指向不同类型的数据。所以我们可以使用它来指向任何数据。
## TODO

1. go 的基本类型和引用类型区分
2. functions 学习，多个返回值学习
3. 声明变量的方式 尤其是冒号
4. 指针 
5. ~~slice map struct~~ 协程
6. map 在并发会 panic 原因？
7. switch 不需要进行 break
8. 在 for 循环中写函数也会面临闭包问题，当然也有相应的解决方案
9. ~~结构体，c++ 中的结构体~~


## 一些 Q and A

1. ~~当我们执行 `go run main.go` 的时候发生了什么~~

- go build main.go 编译成二进制可执行程序
- ./main 执行该程序

2. ~~golang 中单引号和双引号什么区别~~

- 单引号是 字符
- 双引号是 字符串
- 反引号是 表示原生，其中的反斜杠比如常见的换行符 \n 不会被转义

3. ~~golang 中可以使用分号吗~~
golang 中一般只有在分隔 for 循环的初始化语句时经常用到，但在代码段末尾的分号一般是忽略的

4. ~~golang 怎么获取数据类型~~

```go
// vars 为变量
reflect.TypeOf(vars).Kind()
```
5. ~~uint8, int32 是什么类型，能不能举个例子？~~

uint8: [0,255] 无符号整型

int32: [-2147483648 : 2147483647] 整型

6. ~~rune 数组是什么？~~

golang 中 rune 其实 int32 的别名，用于区分字符值和整数值。

当需要处理中文、日文或者其他复合字符时，需要用到 rune 类型（本身是 golang 的一种特殊类型），rune 实际上就是 int32

7. go 如何做类型转换，比如 string()，是不是强类型语言不需要进行类型转换？
    - 空接口可以存储任意类型的数据，因此存储在接口中的数据可以进行类型转换。（不过可以看出来就是显而易见的相同类型转换，而非从类型 1 转换为类型 2）==todo==

8. ~~如何遍历一个数组~~

9. ~~函数如何写 return，尤其在指针那一节想直接 return 改变后的值发现不允许~~

10. const 在 golang 中指的是什么

11. ~~执行 fallthrough 之后 不再进行匹配直接一条路走到 default？~~

12. 如何执行全等操作， golang 数据类型总结

https://juejin.cn/post/6844903923166232589

13. ~~什么是协程， defer 函数是立即执行函数吗~~

Go 协程（Goroutine）是与其他函数同时运行的函数。 可以认为 Go 协程是轻量级的线程，由 Go 运行时来管理。 
在函数调用前加上 go 关键字，这次调用就会在一个新的 goroutine 中并发执行。
当被调用的函数返回时，这个 goroutine 也自动结束。

defer 确实是一个闭包，它在声明时不会立刻去执行，而是在函数 return 之后去执行的

14. golang 中是否有同步函数和异步函数的概念呢

15. ~~golang 如何拼接字符串~~

- 运算符 + 
- fmt.Sprintf()  关于它：内部使用 []byte 实现，不像直接运算符 + 这种会产生很多临时的字符串，但是内部的逻辑比较复杂，有很多额外的判断，还用到了 interface，所以性能也不是很好
- strings.Join 关于它：在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小
- buffer.WriteString()

```go
var buffer bytes.Buffer
buffer.WriteString("hello")
buffer.WriteString(",world")
s := buffer.String()
```
16. golang 如何将整型转化为字符串类型

使用 `strconv`

17. golang 项目中一般使用什么 lint 工具，golang 有开发环境包、生产环境包之分吗

18. golang 中的 go xx 是什么意思


## shell

1. `mv`

修改文件名： hello.go --> main.go

```shell
mv hello.go main.go
```