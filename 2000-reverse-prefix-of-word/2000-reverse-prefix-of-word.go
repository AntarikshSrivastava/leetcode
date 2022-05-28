func reversePrefix(word string, ch byte) string {
    
    idx := -1
    for i:=0; i<len(word); i++ {
        if word[i] == ch {
            idx = i
            break
        }   
    }
    
    if idx == -1 {
        return word
    }
    
    return reverse(word[0:idx+1]) + word[idx+1:]
}

func reverse(s string) string {
    r := []rune(s)
    for i,j:=0,len(s)-1; i<j; i,j=i+1,j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}