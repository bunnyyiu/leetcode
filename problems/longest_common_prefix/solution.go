func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    result := ""
    j := 0
    for {
        if j >= len(strs[0]) {
            return result
        }
        
        char := strs[0][j]
        for i := 1; i < len(strs); i++ {
            if j >= len(strs[i])  {
                return result
            }
            if strs[i][j] != char {
                return result
            }
        }
        result = result + string(char)
        j++
    }
    return result
}