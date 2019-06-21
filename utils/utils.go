package utils

import (
	"strconv"

	"github.com/atotto/clipboard"
)

//float32 转 String工具类，保留6位小数
func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 0, 64)
}
func CopyToClipBoard(content string) {
	// 复制内容到剪切板
	clipboard.WriteAll(content)
	// 读取剪切板中的内容到字符串
	_, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
}
