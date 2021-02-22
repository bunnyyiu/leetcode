func moveZeroes(nums []int)  {
    writeIndex := 0
    cur := 0
    zeroCount := 0
    
    for cur < len(nums){
        if nums[cur] == 0 {
            zeroCount++
        } else {
            nums[writeIndex] = nums[cur]
            writeIndex++
        }
        cur++
    }
    
    for i := 0; i < zeroCount; i++ {
        nums[len(nums) - 1 - i] = 0
    }
}