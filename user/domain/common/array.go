package common

func ArrayToString(arr []string, symbol string) string {
	arrLen := len(arr)
	if arrLen == 0 {
		return ""
	}

	str := ""

	for _, val := range(arr) {
		str += val + symbol
	}

	symbolLen := len(symbol)
	strLen := len(str)

	if symbolLen > 0 {
		str = string([]rune(str)[:strLen - symbolLen])
	}

	return str
}
