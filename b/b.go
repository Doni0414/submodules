package b

func Sum(a ...int) int {
	var res = 0
	for _, el := range a {
		res += el
	}
	return res
}
