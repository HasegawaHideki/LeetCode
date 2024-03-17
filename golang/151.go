func reverseWords(s string) string {
	s = strings.Join(strings.Fields(s), " ")

	var list []string = strings.Split(s, " ")

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	return strings.Join(list, " ")
}