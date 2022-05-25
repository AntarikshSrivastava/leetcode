func makeGood(s string) string {    

    l := len(s)

    for i := 0; i < l - 1;  {
        if (s[i] == s[i+1] + 32 || s[i] == s[i+1] - 32) {
            s= s[:i] + s[i+2:]
            l -= 2
            if i > 0 {
                i--
            }
            continue
        }
        i++
    } 

    return s
}