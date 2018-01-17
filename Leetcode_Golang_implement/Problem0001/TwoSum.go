package Problem0001

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
