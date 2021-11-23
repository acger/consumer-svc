package tool

func FirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}

	strArr := []rune(str)

	if strArr[0] >= 97 && strArr[0] <= 122 {
		strArr[0] -= 32
	}

	return string(strArr)
}