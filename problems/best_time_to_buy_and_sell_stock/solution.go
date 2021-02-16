func maxProfit(prices []int) int {
    maxEarn := math.SmallestNonzeroFloat64
    minPrice := math.MaxFloat64
    
    for _, price := range(prices) {
        minPrice = math.Min(minPrice, float64(price))
        maxEarn = math.Max(maxEarn, float64(float64(price) - minPrice))
    }
    
    return int(maxEarn)
}