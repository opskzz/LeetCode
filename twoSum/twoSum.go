package problems

// Time:  O(n)
// Space: O(n)

func twoSum(nums []int, target int) []int {
	lookUP := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := lookUP[target-num]; ok {
			return []int{j, i}
		}
		lookUP[nums[i]] = i
	}
	return nil
}
