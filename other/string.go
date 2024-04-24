package other

import (
    "fmt"
	"math/big"
)
// a 大于 b 的情况下返回 true
func compareBigNumbers(a, b string) bool {
    // 如果 a 的长度大于 b 的长度，说明 a 更大，返回 true
	if len(a) > len(b) {
		return true
	}
    // 如果 a 的长度小于 b 的长度，说明 b 更大，返回 false
	if len(a) < len(b) {
		return false
	}
    // 如果 a 和 b 的长度相等，则需要逐个字符比较它们
	for i := range a {
        // 如果在某个位置上 a 的字符大于 b 的字符，则可以确定 a 大于 b，返回 true
		if a[i] > b[i] {
			return true
		}
        // 如果在某个位置上 a 的字符小于 b 的字符，则可以确定 a 小于 b，返回 false
		if a[i] < b[i] {
			return false
		}
	}
    // 如果所有字符都相等，说明两个数字相等，按照这个函数的设计，这种情况下返回 false
	return false
}

// SubtractFromString 函数从一个大的数字字符串中减去另一个数字字符串，
// 并返回结果的字符串表示。
func SubtractFromString(bigNumberStr, subtractStr string) (string, error) {
    // 将大数字字符串转换为 big.Int
    bigNumber, ok := new(big.Int).SetString(bigNumberStr, 10)
    if !ok {
        return "", fmt.Errorf("无效的大数字: %s", bigNumberStr)
    }

    // 将要减去的数字字符串转换为 big.Int
    subtractNumber, ok := new(big.Int).SetString(subtractStr, 10)
    if !ok {
        return "", fmt.Errorf("无效的减数: %s", subtractStr)
    }

    // 执行减法运算
    result := new(big.Int).Sub(bigNumber, subtractNumber)

    // 将结果转换回字符串
    return result.String(), nil
}