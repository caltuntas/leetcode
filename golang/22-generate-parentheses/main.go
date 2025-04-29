package main

import "fmt"

func generateParenthesis(n int) []string {
	var results = []string{}
	str := make([]byte, n*2)
	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			str[i] = '('
		} else {
			str[i] = ')'
		}
	}
	text := string(str)
	used := make([]bool, len(str))
	permutations(0, text, &results, used)
	results = removeDuplicateStr(results)
	results = checkResults(results)
	return results
}

func checkResults(results []string) []string {
	list := []string{}
	for _, item := range results {
		fmt.Println(item)
		res := checkItem(item)
		if res {
			list = append(list, item)
		}
		fmt.Println(res)
	}
	return list
}

func checkItem(item string) bool {
	total := 0
	for _, chr := range item {
		if chr == '(' {
			total++
		} else if chr == ')' {
			total--
		}

		if total < 0 {
			return false
		}
	}
	return total == 0
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func permutations(start int, str string, results *[]string, used []bool) {
	fmt.Println("if check")
	if start == len(str) {
		fmt.Println("if block")
		fmt.Println("str=" + str)
		*results = append(*results, str)
		return
	}

	for i := start; i < len(str); i++ {
		fmt.Printf("look for str=%s,i=%d\n", str, i)
		str = swap(str, i, start)
		permutations(start+1, str, results, used)
		str = swap(str, i, start)
		fmt.Println("backtrack item=" + str)
	}
}

func swap(str string, i int, j int) string {
	bytes := []byte(str)
	temp := bytes[i]
	bytes[i] = bytes[j]
	bytes[j] = temp
	return string(bytes)
}

func main() {
	generateParenthesis(3)
}
