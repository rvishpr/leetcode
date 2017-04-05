func longestPalindrome(s string) string {
    pal := make([][]bool, len(s))
    for i := range pal {
        pal[i] = make([]bool, len(s))
    }
    
    longest := 1
    start := 0
    end := 0
    for i := 0; i < len(s); i++ {
        pal[i][i] = true
    }
    
    for i := 0; i < len(s) - 1; i++ {
        if s[i] == s[i+1] {
            longest = 2
            pal[i][i+1] = true
            start = i
            end = i + 1
        }
    }
            
    for i := 3; i <= len(s); i++ {
        for j := 0; j <= len(s) - i; j++ {
            k := i + j - 1
            if s[j] == s[k] && pal[j + 1][k - 1]{
                pal[j][k] = true
                if longest < i {
                    longest = i
                    start = j
                    end = k
                }
            }
        }
    }
    
    return string(s[start:end+1])
}
