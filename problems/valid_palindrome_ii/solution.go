func validPalindrome(s string) bool {
    runeList := []rune(s)
    i := 0
    j := len(runeList) - 1
    
    for i < j {
        if runeList[i] == runeList[j] {
            i++
            j--
        } else {
            curI := i + 1
            curJ := j
            
            // delete left
            for curI < curJ {
                if runeList[curI] == runeList[curJ] {    
                    curI++
                    curJ--
                } else {
                    break
                }
            }
            
            if curI >= curJ {
                return true
            }
            
            
            // delete right
            curI = i
            curJ = j - 1
            
            for curI < curJ {
                if runeList[curI] == runeList[curJ] {
                    curI++
                    curJ--
                } else {
                    return false
                }
            }
            return true
        }
    }
    return true
}