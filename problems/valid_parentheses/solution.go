func isValid(s string) bool {
    stack := []rune{}
    for _, c := range(s) {
        charToCheck := string(c)
        if charToCheck == "(" || charToCheck == "{" || charToCheck == "[" {
            stack = append(stack, c)
        } else if len(stack) == 0 {
            return false
        } else {
            last := string(stack[len(stack) - 1])
            stack = stack[0: len(stack) - 1] 
            
            if last == "(" && charToCheck != ")" {
                return false
            }
            if last == "[" && charToCheck != "]" {
                return false
            }
            if last == "{" && charToCheck != "}" {
                return false
            }
        }
    }
    if len(stack) > 0 {
        return false
    }
    return true
}