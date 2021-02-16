func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)
    for i,num := range nums {
        // find target - num
        if val, found := dict[target - num]; found {
            return []int{val, i}
        }
        dict[num] = i
    }
    return []int{}
}