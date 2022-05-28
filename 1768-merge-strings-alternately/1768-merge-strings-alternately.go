func mergeAlternately(word1 string, word2 string) string {
	i := 0
	res := []byte{}

	for i < len(word1) && i < len(word2) {
		res = append(res, word1[i])
		res = append(res, word2[i])
		i++
	}

	for i < len(word1) {
		res = append(res, word1[i])
		i++
	}

	for i < len(word2) {
		res = append(res, word2[i])
		i++
	}

	return string(res)
}