package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("strings and strconv.")
	//strings.HasPrefix
	var str = "https://www.baidu.com"
	var prefix = "https"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?\n", str, prefix)
	fmt.Printf("%t\n", strings.HasPrefix(str, prefix))

	//strings.HasSuffix
	var strSuffix = "abce.jpg"
	var suffix = ".gif"
	fmt.Printf("the image \"%s\" whether %s image?\n", strSuffix, suffix)
	fmt.Printf("%t\n", strings.HasSuffix(strSuffix, suffix))

	//strings.Contains
	var strContains = "Machine Learning"
	var contains = "Learn"
	fmt.Printf("T/F?Does the string \"%s\" contains %s \n", strContains, contains)
	fmt.Printf("%t\n", strings.Contains(strContains, contains))

	//strings.Index
	var strIndex = "Hello world"
	var index = "o"
	fmt.Printf("The position of \"%s\" in the string \"%s\"  is %d \n", index, strIndex, strings.Index(strIndex, index))

	//strings.LastIndex
	var strLastIndex = "Hello world"
	var lastIndex = "o"
	fmt.Printf("The last position of \"%s\" in the string \"%s\"  is %d \n", lastIndex, strLastIndex, strings.LastIndex(strLastIndex, lastIndex))

	//strings.IndexRune
	var strIndexRune = "计算机程序的结构和解释"
	//var indexRune = ([]rune("结"))[0]
	var indexRune = rune('结')
	fmt.Printf("The position of \"%s\" in the string \"%s\"  is %d \n", string(indexRune), strIndexRune, strings.IndexRune(strIndexRune, indexRune))

	//strings.Replace
	var strReplace = "file edit view go debug"
	var replaceOld = "go"
	var replaceNew = "golang"
	fmt.Printf("After \"%s\" use \"%s\" replace \"%s\" is \"%s\" \n", strReplace, replaceNew, replaceOld, strings.Replace(strReplace, replaceOld, replaceNew, -1))

	//strings.Count
	var str1 = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"
	fmt.Printf("Number of H's in %s is: ", str1)
	fmt.Printf("%d\n", strings.Count(str1, "H"))
	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	//strings.Repeat
	var strRepeat = "Hi here!"
	fmt.Printf("The new repeated string is :" + strings.Repeat(strRepeat, 3))

	//strings.ToLower() and strings.ToUpper()
	var originalStr = "Hey ,How are you!"
	fmt.Printf("The original string is :%s\n", originalStr)
	fmt.Printf("The lowercase string is:%s\n", strings.ToLower(originalStr))
	fmt.Printf("The uppercase string is:%s\n", strings.ToUpper(originalStr))

	//strings.Trim(s,"cut") | strings.TrimSpace(s) |strings.TrimLeft() |strings.TrimRight()
	var trimStr = "=====I'm ybq====="
	var trim = "==="
	fmt.Printf("The original string is : %s \n", trimStr)
	fmt.Printf("The trimed string is : %s \n", strings.Trim(trimStr, trim))

	//strings.Fields() | strings.Split()
	var splitStr = "Hey ,How are you"
	fields := strings.Fields(splitStr)
	splits := strings.Split(splitStr, ",")
	fmt.Printf("\nThe origin string is : %s\n", splitStr)
	fmt.Printf("Split with space: \n ")
	for _, field := range fields {
		fmt.Printf("%s \n ", field)
	}
	fmt.Printf("Split with ,: \n ")
	for _, split := range splits {
		fmt.Printf("%s \n ", split)
	}

	//strings.Join()
	var joinArray = []string{"How", "are", "you"}
	fmt.Printf("Join string is : %s \n", strings.Join(joinArray, " "))

	//strings.NewReader(str)
	var readerStr = "该包包含了一些变量用于获取程序运行的操作系统平台下 int 类型所占的位数"
	var reader = strings.NewReader(readerStr)
	fmt.Printf("Use NewReader() read string:")
	for ch, _, _ := reader.ReadRune(); ch != 0; ch, _, _ = reader.ReadRune() {
		fmt.Printf("%v", string(ch))
	}

	//strconv
	var i = strconv.Itoa(100)
	fmt.Printf("\n%s\n", i)
	var f = strconv.FormatFloat(3.1415932, 'f', 3, 32)
	fmt.Printf("%s\n", f)

	var in, err = strconv.Atoi("111")
	fmt.Printf("%v\n", err)
	fmt.Printf("%d\n", in)

	var fl, err1 = strconv.ParseFloat(f, 32)
	fmt.Printf("%v\n",err1)
	fmt.Printf("%f\n",fl)
}
