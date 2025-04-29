package main


func generateParenthesis(n int) []string {
	var results = []string{}
	str := make([]byte, n*2)
	for i := 0; i < len(str); i++ {
		if i < n {
			str[i] = '('
		} else {
			str[i] = ')'
		}
	}
	text := string(str)
	used := make([]bool, len(str))
	permutations(text, &results, used, "")
	return results
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

func permutations(str string, results *[]string, used []bool, path string) {
	if len(path) == len(str) {
		if checkItem(path) {
			*results = append(*results, path)
		}
		return
	}

	for i := 0; i < len(str); i++ {
		if used[i] {
			continue
		}

		if i > 0 && str[i] == str[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		path += string(str[i])
		permutations(str, results, used, path)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func main() {
	generateParenthesis(3)
}
