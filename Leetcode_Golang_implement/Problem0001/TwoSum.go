package Problem0001
/**
题目：
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
 */
/**
思路:
	将所有的值塞入map中，key存nums[index] value存index
 */
func twoSum(nums[] int,target int) []int{
	m := make(map[int]int,len(nums))

	for i,b:=range nums {
		if j,ok := m[target-b];ok {
			return []int{i,j}
		}
		m[nums[i]] = i
	}
	return nil
}
