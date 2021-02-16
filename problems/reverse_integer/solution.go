func reverse(x int) int {
    result := 0
    current := x
    for current != 0 {
        remain := current % 10
        current /= 10
        result = result * 10 + remain
    }
    if result > int(math.Pow(2, 31)) -1 || result < int(math.Pow(-2, 31)) {
        return 0
    }
    return result
}