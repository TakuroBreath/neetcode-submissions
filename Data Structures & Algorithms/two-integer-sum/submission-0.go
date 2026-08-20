func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))

    for i, v := range nums {
        t := target - v
        if a, exists := m[v]; exists {
            return []int{a, i}
        }
        m[t] = i
    }

    return []int{}
}
