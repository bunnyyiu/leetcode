// Three Pointers (with extra num1 copy)
/*
func merge(nums1 []int, m int, nums2 []int, n int)  {
    tmpNum1Copy := make([]int, len(nums1))
    
    for i := 0; i < len(nums1); i++ {
        tmpNum1Copy[i] = nums1[i]
    }
    
    i := 0
    j := 0
    writeIndex := 0
    for i < m && j < n {
        if tmpNum1Copy[i] < nums2[j] {
            nums1[writeIndex] = tmpNum1Copy[i]
            i += 1
            writeIndex += 1
        } else {
            nums1[writeIndex] = nums2[j]
            j += 1
            writeIndex += 1
        }
    }
    
    for i < m {
        nums1[writeIndex] = tmpNum1Copy[i]
        i += 1
        writeIndex += 1
    }
    
    for j < n {
        nums1[writeIndex] = nums2[j]
        j += 1
        writeIndex += 1
    }
}
*/

// Three pointers without extra storage
func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := m - 1
    j := n - 1
    
    k := m + n - 1
    
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
            k--
        } else {
            nums1[k] = nums2[j]
            j--
            k--
        }
    }
    
    for i >= 0 {
        nums1[k] = nums1[i]
        i--
        k--
    }
    
    for j >= 0 {
        nums1[k] = nums2[j]
        j--
        k--
    }
}