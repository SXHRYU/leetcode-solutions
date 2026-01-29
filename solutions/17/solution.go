package leetcode

var (
	letterMap = []string{
		" ",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	res   = []string{}
	final = 0
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
	return
}

func letterCombinations_(digits string) []string {
	if digits == "" {
		return []string{}
	}
	index := digits[0] - '0'
	letter := letterMap[index]
	tmp := []string{}
	for i := 0; i < len(letter); i++ {
		if len(res) == 0 {
			res = append(res, "")
		}
		for j := 0; j < len(res); j++ {
			tmp = append(tmp, res[j]+string(letter[i]))
		}
	}
	res = tmp
	final++
	letterCombinations(digits[1:])
	final--
	if final == 0 {
		tmp = res
		res = []string{}
	}
	return tmp
}

var result []string
var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinationsBT(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	letterFunc("", digits)
	return result
}

func letterFunc(res string, digits string) {
	if digits == "" {
		result = append(result, res)
		return
	}
	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		letterFunc(res, digits)
		res = res[0 : len(res)-1]
	}
}
