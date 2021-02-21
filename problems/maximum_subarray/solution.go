func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxSubArray(nums []int) int {
    currentMax := nums[0]
    globalMax := nums[0]
    
    for i := 1; i < len(nums); i++ {
        x := nums[i]
        currentMax = max(currentMax + x, x)
        globalMax = max(globalMax, currentMax)
    }
    
    return globalMax
}