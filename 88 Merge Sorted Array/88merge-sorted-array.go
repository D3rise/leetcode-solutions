import "math"

func merge(nums1 []int, m int, nums2 []int, n int)  {
    p1, p2 := m-1, n-1
    for i := len(nums1)-1; i >= 0; i-- {
        n1 := math.MinInt32
        n2 := math.MinInt32

        if p1 >= 0 {
            n1 = nums1[p1]
        }

        if p2 >= 0 {
            n2 = nums2[p2]
        }

        if n1 >= n2 {
            nums1[i] = n1
            p1--
        } else {
            nums1[i] = n2
            p2--
        }
    }
}