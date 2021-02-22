func fizzBuzz(n int) []string {
    results := make([]string, n)
    
    for i := 1; i <= n; i++ {
        if i % 15 == 0 {
            results[i-1] = "FizzBuzz"
        } else if i % 5 == 0 {
            results[i-1] = "Buzz"
        } else if i % 3 == 0 {
            results[i-1] = "Fizz"
        } else {
            results[i-1] = fmt.Sprintf("%d", i)
        }
    }
    
    return results
}