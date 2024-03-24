func compress(chars []byte) int {
	var size int = len(chars)
	if size == 1 {
		return 1
	}

	var idx int = 1
	var count int = 1
	for i := 1; i < size; i++ {
		var same bool = false
		if chars[i] == chars[idx-1] {
			count++
			same = true

			if i < size-1 {
				continue
			}
		}

		if count >= 10 {
			var count_bytes []byte = []byte(strconv.Itoa(count))
			for k := 0; k < len(count_bytes); k++ {
				chars[idx], idx = count_bytes[k], idx+1
			}
		} else if count > 1 {
			chars[idx], idx = strconv.Itoa(count)[0], idx+1
		}

		if same == false {
			chars[idx], idx = chars[i], idx+1
			count = 1
		}
	}

	return idx
}