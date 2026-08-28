func getConcatenation(nums []int) []int {
    res := make([]int, 0, len(nums)*2)

    res = append(res, nums...)
    res = append(res[:len(nums)], nums...)

    return res
}