package main 
var digitsMap = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) <= 0 {
		return result
	}
	combination(digits, 0, "", &result)
	return result
}

func combination(digits string, level int, path string, result *[]string) {
	if level == len(digits) {
		*result = append(*result, path)
		return
	}
	digit := int(digits[level] - '0')
	letters := digitsMap[digit]
	for i := 0; i < len(letters); i++ {
		path += string(letters[i])
		combination(digits, level+1, path, result)
		path = path[:len(path)-1]
	}
}

func main() {
	letterCombinations("234")
	letterCombinations("")
	letterCombinations("2")
	letterCombinations("2345")
	letterCombinations("24")
	letterCombinations("278")
}
