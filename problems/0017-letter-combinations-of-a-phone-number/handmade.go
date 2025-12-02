package lettercombinationsofaphonenumber

const debugLetterCombos = false

var digitMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// letterCombinations returns all possible letter strings for the digit input.
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var res []string
	path := make([]byte, len(digits))

	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == len(digits) {
			res = append(res, string(path))
			return
		}

		letters := digitMap[digits[pos]]
		for i := 0; i < len(letters); i++ {
			path[pos] = letters[i]
			dfs(pos + 1)
		}
	}

	dfs(0)
	return res
}
