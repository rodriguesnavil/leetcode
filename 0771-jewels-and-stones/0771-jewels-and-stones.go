func numJewelsInStones(jewels string, stones string) int {
    jewelRns := [] rune(jewels)
    stoneRns := [] rune(stones)
    
    counter := 0
    
    for _, jChar := range jewelRns {
        for _, sChar := range stoneRns {
            if(string(jChar) == string(sChar)) {
                counter += 1
            }
        }
    }
    
    return counter
}