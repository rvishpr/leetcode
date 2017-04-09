func wordBreak(s string, words []string) bool {
    dict := make(map[string]bool)
    for _, word := range words {
        dict[word] = true
    }
    seen := make(map[string]bool)

    return wbreak(s, dict, seen)
}

func wbreak(s string, dict, seen map[string]bool) bool {
    l := len(s)
    if dict[s] {
        return true
    }
    
    for i := 1; i <= l; i++ {
        first := dict[string(s[0:i])] 
        rest := true
        if i < l {
            if seen[string(s[i:l])] {
                rest = dict[string(s[i:l])]
            } else {
                rest = wbreak(string(s[i:l]), dict, seen)
                seen[string(s[i:l])] = true
                dict[string(s[i:l])] = rest
            }
        }
        if first && rest {
            return true
        }
    }
    return false
}
