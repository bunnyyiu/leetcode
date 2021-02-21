func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxProfit(prices []int) int {
    maxEarn := math.MinInt32
    minPrice := math.MaxInt32
    
    for _, price := range(prices) {
        minPrice = min(minPrice, price)
        maxEarn = max(maxEarn, price - minPrice)
    }
    
    return maxEarn
}