func reverseVowels(s string) string {
	var vowel string = "aiueoAIUEO"
	indexes := make([]bool, len(s))
	vowels := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if strings.IndexAny(string(s[i]), vowel) != -1 {
			vowels = append(vowels, s[i])
			indexes[i] = true
		} else {
			indexes[i] = false
		}
	}

	bytes := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if indexes[i] == true {
			bytes[i] = vowels[len(vowels)-1]
			vowels = vowels[:len(vowels)-1]
		} else {
			bytes[i] = s[i]
		}
	}

	return string(bytes)
}