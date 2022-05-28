func merge(nums1 []int, m int, nums2 []int, n int)  {
    curr := m+n-1
    i := m-1
    j := n-1
    
    if m == 0 {
        for j>=0 {
            nums1[curr] = nums2[j]
            curr--
            j--
        }
    } else {
    
        for i>=0 && j>=0 {
            if nums1[i] > nums2[j] {
                nums1[curr] = nums1[i]
                i--
            } else {
                nums1[curr] = nums2[j]
                j--
            }
            curr--
        }

        for i>=0 {
            nums1[curr] = nums1[i]
            curr--
            i--
        }

        for j>=0 {
            nums1[curr] = nums2[j]
            curr--
            j--
        }
    }
}