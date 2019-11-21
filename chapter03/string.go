package main

import (
	"fmt"
)

func main() {
	//string 的值不可以变，是一个定长的字节数组

	//使用反引号括起来，支持换行
	var s = `This is a raw string \n`
	var normalS = "This is a normal string \n"
	fmt.Println(s)
	fmt.Println(normalS)

	//获取字符串的字节长度
	var lenS = "我是一个字串"
	fmt.Printf("字符串(%s)字节长度:%d\n", lenS, len(lenS))
	//range方式以UTF-8单个字作为步长
	for ix, ss := range lenS {
		fmt.Printf("utf-8 format: index:%d,content:%v \n", ix, ss)
		fmt.Println(string(ss))
	}

	//以字节索引的方式循环字符串
	for i := 0; i < len(lenS); i++ {
		fmt.Printf("byte format:index:%d,content:%v \n", i, lenS[i])
	}

	//字符串拼接
	//方式1
	var joinStr = "hello " + "world"
	//方式2
	joinStr += "!"
	//方式3 跨行
	joinStr += "This is a long string" +
		" ......"
	fmt.Printf("%s\n", joinStr)

	//统计字符串的字节数和字符串
	countStr1 := "asSASA ddd dsjkdsjs dk"
	countStr2 := "asSASA ddddsjkdsjsこん dk"
	bytes1, chars1 := CountStr(countStr1)
	bytes2, chars2 := CountStr(countStr2)
	fmt.Printf("bytes:%d and charactors:%d of (%s)\n", bytes1, chars1, countStr1)
	fmt.Printf("bytes:%d and charactors:%d of (%s)\n", bytes2, chars2, countStr2)
}

//CountStr 统计字符串字节数和字符数
func CountStr(str string) (bytes int, chars int) {
	bytes = len([]byte(str))
	chars = len([]rune(str))
	return bytes, chars
}
