func firstUniqChar(s string) int {
	ch := make([]int, 26)

	for _, c := range s {
		ch[c-'a']++
	}

	for i := 0; i < len(s); i++ {
		if ch[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}