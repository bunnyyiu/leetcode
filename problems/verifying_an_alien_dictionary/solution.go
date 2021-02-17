func compare(wordA, wordB string, hash map[rune]int) int{
    lenWordA := len(wordA)
    lenWordB := len(wordB)
    
    charA := []rune(wordA)
    charB := []rune(wordB)
    
    maxLen := lenWordA
    if lenWordB > maxLen {
        maxLen = lenWordB
    }
    
    for i:= 0; i < maxLen; i++ {
        if i >= lenWordA {
            return -1
        } else if i >= lenWordB {
            return 1
        } else {
            a := hash[charA[i]]
            b := hash[charB[i]]
            
            if a < b {
                return -1
            } else if a > b {
                return 1
            }
        }
    }
    return 0
} 

func isAlienSorted(words []string, order string) bool {
    hash := make(map[rune]int)
    for j, char := range(order) {
        hash[char] = j
    }
    for i := 1; i < len(words); i++ {
        wordA := words[i - 1]
        wordB := words[i]
        result := compare(wordA, wordB, hash)
        if result > 0 {
            return false
        }
    }
    return true
}