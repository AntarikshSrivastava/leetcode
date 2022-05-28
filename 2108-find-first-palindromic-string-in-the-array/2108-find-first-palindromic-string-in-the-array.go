func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindromic(word) { return word }
    }
    return ""
}

func isPalindromic(s string) bool {
    for l, r := 0, len(s) - 1; l < r; l, r = l + 1, r - 1 {
        if s[l] != s[r] { return false }
    }
    return true
}