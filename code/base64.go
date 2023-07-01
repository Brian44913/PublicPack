package PublicPackCode

import (
	"strings"
	"encoding/base64"
)

func Base64UrlEncode(str string) string {
	// 将字符串进行 Base64 编码
	base64String := base64.StdEncoding.EncodeToString([]byte(str))
	// 替换特殊字符
	paddedString := strings.NewReplacer("+", "-", "/", "_", "=", "").Replace(base64String)
	return paddedString
}

func Base64UrlDecode(str string) ([]byte, error) {
	// 添加填充字符
	padding := len(str) % 4
	if padding > 0 {
		str += strings.Repeat("=", 4-padding)
	}
	// 替换特殊字符
	urlSafeString := strings.NewReplacer("-", "+", "_", "/").Replace(str)
	// 解码 Base64 字符串
	decodedData, err := base64.StdEncoding.DecodeString(urlSafeString)
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}