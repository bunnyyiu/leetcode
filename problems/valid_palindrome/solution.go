func isDigit(d rune) bool {
    return d >= '0' && d <= '9'
}

func isAlpha(a rune) bool {
    return a >= 'a' && a <= 'z'
}

func toSkip(a rune) bool {
    return !isDigit(a) && !isAlpha(a)
}

func isPalindrome(s string) bool {
    sLower := strings.ToLower(s)
    i := 0
    j := len(sLower) - 1
    
    runes := []rune(sLower)
    for i < j {
        curI := runes[i]
        curJ := runes[j]
        
        if toSkip(curI) {
            i++
            continue
        }
        if toSkip(curJ) {
            j--
            continue
        }
        if curI == curJ {
            i++
            j--
        } else {
            return false
        }
        
    }
    
    return true
}