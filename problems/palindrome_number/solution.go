func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    current := x
    reversed := 0
    for current != 0 {
        remain := current % 10
        reversed = reversed * 10 + remain
        current /= 10
    }
    return x == reversed
}