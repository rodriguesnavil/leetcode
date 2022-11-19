func isAnagram(s string, t string) bool {
    if(len(s) != len(t)) {
        return false
    }
    
    sRns := []rune(s)
    tRns := []rune(t)
    
    sort.Slice(sRns, func(i, j int) bool {
        return sRns[i] < sRns[j]
    })
    
    sort.Slice(tRns, func(i, j int) bool {
        return tRns[i] < tRns[j]
    })
    
    if (string(sRns) == string(tRns)) {
        return true
    } else {
        return false
    }
}