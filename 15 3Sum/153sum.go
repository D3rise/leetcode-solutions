import "sort"

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}

    for i, num1 := range nums {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        target := -num1

        p1, p2 := i+1, len(nums)-1
        for p1 < p2 {
            num2, num3 := nums[p1], nums[p2]
            sum := num2 + num3
            if sum == target {
                result = append(result, []int{num1, num2, num3})
                for p1 < p2 && nums[p1] == nums[p1+1] {
                    p1++
                }

                for p2 > p1 && nums[p2] == nums[p2-1] {
                    p2--
                }

                p1++
                p2--
            } else if sum < target {
                p1++
            } else if sum > target {
                p2--
            }
        }
    }

    return result
}