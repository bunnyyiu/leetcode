func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minMoves(nums []int) int {
    if len(nums) == 0 || len(nums) == 1 {
        return 0
    }
    
    minInt := math.MaxInt32
    sum := 0
    for _, val := range nums {
        sum += val
        minInt = min(val, minInt)
    }
    
    return sum - minInt * len(nums)
}