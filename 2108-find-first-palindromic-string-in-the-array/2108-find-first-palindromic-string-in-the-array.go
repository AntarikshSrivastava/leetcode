func firstPalindrome(words []string) string {
    ans := ""
    for _, word := range words {
        if isPalindromic(word) { ans = word; break }
    }
    return ans
}

func isPalindromic(s string) bool {
    for l, r := 0, len(s) - 1; l < r; l, r = l + 1, r - 1 {
        if s[l] != s[r] { return false; break }
    }
    return true
}