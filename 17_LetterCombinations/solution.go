package lettercombinations

/*
Link: https://leetcode.com/problems/letter-combinations-of-a-phone-number?envType=problem-list-v2&envId=backtracking
Type: Backtracking
Description:
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on telephone buttons) is given below:
2 -> abc
3 -> def
4 -> ghi
5 -> jkl
6 -> mno
7 -> pqrs
8 -> tuv
9 -> wxyz

Solution approach:
Use backtracking/DFS to generate all possible combinations. For each digit, try all possible letters mapped to that digit. Pass the current combination state through recursive DFS calls to build up the combinations.
*/

var DigitToAlphabets = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}

	var dfs func(digits string, path []rune)
	dfs = func(digits string, path []rune) {
		if len(digits) == 0 {
			result = append(result, string(path))
			return
		}

		prefix := rune(digits[0])
		for _, alphabet := range DigitToAlphabets[prefix] {
			temp := append(path, alphabet)
			dfs(digits[1:], temp)
		}
	}

	dfs(digits, make([]rune, 0))
	return result
}
