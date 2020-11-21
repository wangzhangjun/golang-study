package main // 声明 main 包，表明当前是一个可执行程序
import "fmt" // 导入内置 fmt 包

/*
因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。
字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
*/
func traversalString(_str string) {
	var count int = 0
	for i := 0; i < len(_str); i++ { //byte
		fmt.Printf("%v(%c) ", _str[i], _str[i]) //104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³)
	}
	fmt.Println()
	for _, r := range _str { //rune
		fmt.Printf("%v(%c) ", r, r) //104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)
		if len(string(r)) >= 3 {
			count++
		}
	}
	fmt.Println()
	fmt.Println("汉子的个数是：", count)

}

func main() { // main函数，是程序执行的入口
	fmt.Println("Hello World!") // 在终端打印 Hello World!
	var (
		_int    int32   = 1
		_float  float32 = 1.0
		_bool   bool    = true
		_string string  = "hello"
	)
	fmt.Printf("%d 的类型为:%T\n", _int, _int)
	fmt.Printf("%f 的类型为:%T\n", _float, _float)
	fmt.Printf("%t 的类型为:%T\n", _bool, _bool)
	fmt.Printf("%s 的类型为:%T\n", _string, _string)

	_str := "hello沙河你"
	traversalString(_str)
}
