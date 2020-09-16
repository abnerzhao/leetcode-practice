package array

func twoNum(nums []int, target int) []int {
	mapDict := make(map[int]int)
	var data []int
	for i, k := range nums {
		_, status := mapDict[target-k]
		if !status {
			mapDict[k] = i
		} else {
			data = []int{i, mapDict[target-k]}
			break
		}
	}
	return data
}
