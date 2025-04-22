package main 

import "fmt"

func letterCombinations(digits string) []string {
	var digitsMap = map[int]string{}
	digitsMap = make(map[int]string)
	digitsMap[2] = "abc"
	digitsMap[3] = "def"
	digitsMap[4] = "ghi"
	digitsMap[5] = "jkl"
	digitsMap[6] = "mno"
	digitsMap[7] = "pqrs"
	digitsMap[8] = "tuv"
	digitsMap[9] = "wxyz"
	result := []string{}
	if len(digits) <= 0 {
		return result
	}

	var chrDigit1 uint8
	var chrDigit2 uint8
	var chrDigit3 uint8
	var chrDigit4 uint8
	var str string
	var str1 string
	var str2 string
	var str3 string
	if len(digits) >= 1 { 
		chrDigit1 = digits[0]
		digit1 := int(chrDigit1 - '0')
		str = digitsMap[digit1]
    }

    if len(digits) >= 2 {
		chrDigit2 = digits[1]
		digit2 := int(chrDigit2 - '0')
		str1 = digitsMap[digit2]
    }

    if len(digits) >=3 {
		chrDigit3 = digits[2]
		digit3 := int(chrDigit3 - '0')
		str2 = digitsMap[digit3]
    }

    if len(digits) >=4 {
		chrDigit4 = digits[3]
		digit4 := int(chrDigit4 - '0')
		str3 = digitsMap[digit4]
    }

	fmt.Printf("Digits=%c\n", chrDigit1)
	if len(digits) == 1 {
		for _, chr := range str {
			fmt.Printf("level 1 chr=%c\n", chr)
			resultStr := string(chr)
			result = append(result, resultStr)
		}
	}
	if len(digits) == 2 {
		for _, chr := range str {
			fmt.Printf("level 1 chr=%c\n", chr)
			for i := 0; i < len(str1); i++ {
				fmt.Printf("--level 2 chr=%s\n", string(str1[i]))
				resultStr := string(chr) + string(str1[i])
				result = append(result, resultStr)
			}
		}
	}
	if len(digits) == 3 {
		for _, chr := range str {
			fmt.Printf("level 1 chr=%c\n", chr)
			for i := 0; i < len(str1); i++ {
				fmt.Printf("--level 2 chr=%s\n", string(str1[i]))
				for j := 0; j < len(str2); j++ {
					fmt.Printf("----level 3 chr=%s\n", string(str2[j]))
					resultStr := string(chr) + string(str1[i]) + string(str2[j])
					fmt.Printf("####str=%s\n", resultStr)
					result = append(result, resultStr)
				}
			}
		}
	}

	if len(digits) == 4 {
		for _, chr := range str {
			fmt.Printf("level 1 chr=%c\n", chr)
			for i := 0; i < len(str1); i++ {
				fmt.Printf("--level 2 chr=%s\n", string(str1[i]))
				for j := 0; j < len(str2); j++ {
                    fmt.Printf("--level 3 chr=%s\n", string(str2[j]))
                    for k := 0; k < len(str3); k++ {
					fmt.Printf("----level 4 chr=%s\n", string(str3[k]))
					resultStr := string(chr) + string(str1[i]) + string(str2[j]) + string(str3[k])
					fmt.Printf("####str=%s\n", resultStr)
					result = append(result, resultStr)
                    }
				}
			}
		}
    }
	return result
}

func main() {
	letterCombinations("234")
	letterCombinations("")
	letterCombinations("2")
	letterCombinations("2345")
	letterCombinations("24")
	letterCombinations("278")
	fmt.Println("hello world")
}
