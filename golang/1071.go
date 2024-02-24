func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)
	g := gcd(len1, len2)
	substr := str1[:g]

	if strings.Repeat(substr, len1/g) == str1 && strings.Repeat(substr, len2/g) == str2 {
		return substr
	}
	return ""
}