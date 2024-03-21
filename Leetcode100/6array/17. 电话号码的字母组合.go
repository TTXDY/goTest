package main

/**
 * @ Tool：IntelliJ IDEA
 * @ Author：Jay
 * @ Date：2024-02-29-09:15
 * @ Version：1.0
 * @ Description：17. 电话号码的字母组合
 */

func letterCombinations(digits string) []string {
	legth := len(digits)
	if legth == 0 {
		return make([]string, 0)
	}
	table := make(map[string][]string)
	table["2"] = []string{"a", "b", "c"}
	table["3"] = []string{"d", "e", "f"}
	table["4"] = []string{"g", "h", "i"}
	table["5"] = []string{"j", "k", "l"}
	table["6"] = []string{"m", "n", "o"}
	table["7"] = []string{"p", "q", "r", "s"}
	table["8"] = []string{"t", "u", "v"}
	table["9"] = []string{"w", "x", "y", "z"}
	allString := make([]string, 0)
	if legth == 1 {
		return table[digits]
	}
	if legth == 2 {

	}
	if legth == 3 {

	}
	if legth == 4 {

	}
	return allString
}
