func addStrings(num1 string, num2 string) string {
    overflow := 0
    curString := ""
    
    num1Char := []rune(num1)
    num2Char := []rune(num2)
    
    for i := 0; i < len(num1Char) || i < len(num2Char); i++ {
        num1 := 0
        num2 := 0
        
        if i < len(num1Char) {
            num1, _ = strconv.Atoi(string(num1Char[len(num1Char) - 1 - i]))
        }
        
        if i < len(num2Char) {
            num2, _ = strconv.Atoi(string(num2Char[len(num2Char) -1 - i]))
        }
        
        sum := num1 + num2 + overflow
        
        if sum >= 10 {
            sum -= 10
            overflow = 1
        } else {
            overflow = 0
        }
        curString = fmt.Sprintf("%d%s", sum, curString)
    }
    if overflow > 0 {
        curString = fmt.Sprintf("%d%s", 1, curString)
    }
    return curString
}