package myself

var Test int

// Tip init 函数会先于 main 函数执行，实现包级别的一些初始化操作
func init() {
	println("===now run myself pkg===")
	Test = 12345
}