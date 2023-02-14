package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("请输入要转换的字符串：")
	scanner.Scan()
	input := scanner.Text()

	fmt.Print("请选择转换方向：1 - ASCII 码转字符串，2 - 字符串转 ASCII 码：")
	scanner.Scan()
	direction, _ := strconv.Atoi(scanner.Text())

	switch direction {
	case 1:
		// ASCII 码转字符串
		output := ""
		for _, c := range input {
			output += string(c)
		}
		fmt.Printf("ASCII 码 %v 转换结果：%v\n", input, output)
	case 2:
		// 字符串转 ASCII 码
		fmt.Println("字符串转 ASCII 码转换结果：")
		for _, c := range input {
			fmt.Printf("%d ", c)
		}
	}
}
